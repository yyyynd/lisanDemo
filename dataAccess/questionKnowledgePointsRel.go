package dataAccess

import "time"

type QuestionKnowledgePointsRel struct {
	QkpId      string     `gorm:"column:primaryKey;qkp_id" json:"QkpId"` //type:string       comment:题目和知识点关系记录id;question_knowledge_points_rel主键列，使用自增列，无需手动指定   version:2023-06-02 15:11
	QId        string     `gorm:"column:q_id" json:"QId"`                //type:string       comment:题目id;对应question_info主键列 q_id                                                    version:2023-06-02 15:11
	KpId       string     `gorm:"column:kp_id" json:"KpId"`              //type:string       comment:知识点id;对应knowledge_points_info主键列 kp_id                                         version:2023-06-02 15:11
	UpdateTime *time.Time `gorm:"column:update_time" json:"UpdateTime"`  //type:*time.Time   comment:更新时间;YYYY-MM-DD                                                                    version:2023-06-02 15:11
	DeleteMark *int       `gorm:"column:delete_mark" json:"DeleteMark"`  //type:*int         comment:逻辑删除标记;0代表未删除，1代表删除，默认值0                                           version:2023-06-02 15:11
}

// TableName 表名:question_knowledge_points_rel，题目知识点关系表。
// 说明:
func (QuestionKnowledgePointsRel) TableName() string {
	return "question_knowledge_points_rel"
}
