package dataAccess

import (
	demoServer "demoProject/biz/model/demoServer"
	"time"
)

const CORRECT_FLAG = 1

type AnswerRecords struct {
	ArId       string     `gorm:"column:primaryKey;ar_id" json:"ArId"`  //type:string       comment:答题记录id;answer_record主键列，使用自增列，无需手动指定   version:2023-06-02 14:51
	SId        string     `gorm:"column:s_id" json:"SId"`               //type:string       comment:关联学生                                                   version:2023-06-02 14:51
	QId        string     `gorm:"column:q_id" json:"QId"`               //type:string       comment:关联题目                                                   version:2023-06-02 14:51
	IsCorrect  *int       `gorm:"column:is_correct" json:"IsCorrect"`   //type:*int         comment:对错状态;0代表错误，1代表正确                              version:2023-06-02 14:51
	UpdateTime *time.Time `gorm:"column:update_time" json:"UpdateTime"` //type:*time.Time   comment:最后更新时间;YYYY-MM-DD                                    version:2023-06-02 14:51
	DeleteMark *int       `gorm:"column:delete_mark" json:"DeleteMark"` //type:*int         comment:逻辑删除标记;0代表未删除，1代表删除，默认值0               version:2023-06-02 14:51
}

// TableName 表名:answer_records，答题情况记录表。
// 说明:
func (AnswerRecords) TableName() string {
	return "answer_records"
}

type accuracyInfoData struct {
	KpId     string
	Accuracy float64
}

func GetAllStudentAccuracyInfo() ([]*demoServer.StuInfoRespData, error) {
	ids, err := GetAllStuId()
	if err != nil {
		return nil, err
	}

	data := make([]*demoServer.StuInfoRespData, 0)
	for i := 0; i < len(ids); i++ {
		accuracyData, err := getStudentAccuracyInfo("", ids[i])
		if err != nil {
			return nil, err
		}
		stuData := new(demoServer.StuInfoRespData)
		stuData.Accuracy = accuracyData
		stuData.ID = ids[i]
		name, _ := GetStuName(ids[i])
		if err != nil {
			return nil, err
		}
		stuData.Name = name
		data = append(data, stuData)
	}

	return data, nil
}

func GetStudentAccuracyInfo(examID string, stuId string) ([]*demoServer.KnowledgePointAccuracy, error) {
	if accuracyData, err := getStudentAccuracyInfo(examID, stuId); err != nil {
		return nil, err
	} else {
		return accuracyData, nil
	}
}

func getStudentAccuracyInfo(examID string, stuId string) ([]*demoServer.KnowledgePointAccuracy, error) {
	db, err := InitConnection(USER, PASSWD, "", "lisandb")
	if err != nil {
		return nil, err
	}
	//check if student exist
	deleteMark := 1
	db.Model(StudentsInfo{}).
		Select("delete_mark").
		Where("s_id = (?)", stuId).
		Find(&deleteMark)
	if db.Error != nil {
		return nil, db.Error
	} else if deleteMark != 0 {
		return nil, ErrStuNotExist
	}

	//get accuracy data
	accuracy := make([]accuracyInfoData, 0)
	//collect each knowledge
	subQueryA := db.Model(QuestionKnowledgePointsRel{}).
		Select("kp_id", "COUNT(*) AS total_num").
		Group("kp_id")
	subQueryB := db.Model(AnswerRecords{}).
		Select("question_knowledge_points_rel.kp_id, COUNT(*) AS correct_num").
		Joins("INNER JOIN question_knowledge_points_rel ON answer_records.s_id = (?) AND "+
			"answer_records.q_id = question_knowledge_points_rel.q_id", stuId).
		Where("answer_records.is_correct = (?)", CORRECT_FLAG).
		Group("question_knowledge_points_rel.kp_id")
	//combine two subQuery by inner join
	db.Table("(?) AS a", subQueryA).
		Select("a.kp_id, b.correct_num / a.total_num AS accuracy").
		Joins("INNER JOIN (?) AS b ON a.kp_id = b.kp_id", subQueryB).
		Group("a.kp_id").
		Find(&accuracy)

	/** SELECT g.kp_id, f.correct_num / g.total_num AS accuracy
	FROM
		(SELECT b.kp_id, COUNT(*) AS total_num  FROM
		question_knowledge_points_rel AS b
		GROUP BY b.kp_id) AS g
	INNER JOIN
		(SELECT d.kp_id, COUNT(*) AS correct_num
		FROM answer_records AS c INNER JOIN question_knowledge_points_rel AS d
		ON c.s_id = "2019213860" AND c.q_id = d.q_id
		WHERE c.is_correct = 1
		GROUP BY d.kp_id) AS f
	ON g.kp_id = f.kp_id
	GROUP BY g.kp_id;
	*/

	if db.Error != nil {
		return nil, db.Error
	}

	//create resp body
	accuracyData := make([]*demoServer.KnowledgePointAccuracy, len(accuracy))
	for i := 0; i < len(accuracy); i++ {
		tmp := demoServer.NewKnowledgePointAccuracy()
		tmp.Kid = accuracy[i].KpId
		tmp.Accuracy = accuracy[i].Accuracy
		accuracyData[i] = tmp
	}

	return accuracyData, nil
}
