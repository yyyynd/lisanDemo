package dataAccess

import "time"

type StudentsInfo struct {
	SId        string     `gorm:"column:primaryKey;s_id" json:"SId"`    //type:string       comment:学生id;students_info主键列                     version:2023-05-27 20:40
	SName      *int       `gorm:"column:s_name" json:"SName"`           //type:*int         comment:学生姓名                                       version:2023-05-27 20:40
	Class      string     `gorm:"column:class" json:"Class"`            //type:string       comment:所属班级                                       version:2023-05-27 20:40
	UpdateTime *time.Time `gorm:"column:update_time" json:"UpdateTime"` //type:*time.Time   comment:最后更新时间;YYYY-MM-DD                        version:2023-05-27 20:40
	DeleteMark *int       `gorm:"column:delete_mark" json:"DeleteMark"` //type:*int         comment:逻辑删除标记;0代表未删除，1代表删除，默认值0   version:2023-05-27 20:40
}

// TableName 表名:students_info，学生信息表。
// 说明:
func (StudentsInfo) TableName() string {
	return "students_info"
}

func GetStuName(stuId string) (string, error) {
	db, err := InitConnection(USER, PASSWD, "", "lisandb")
	if err != nil {
		return "", err
	}
	//check if student exist
	deleteMark := 1
	db.Model(StudentsInfo{}).
		Select("delete_mark").
		Where("s_id = (?)", stuId).
		Find(&deleteMark)
	if db.Error != nil {
		return "", db.Error
	} else if deleteMark != 0 {
		return "", ErrStuNotExist
	}

	//query
	var name string
	db.Model(StudentsInfo{}).
		Select("s_name").
		Where("s_id = (?)", stuId).
		Find(&name)

	if db.Error != nil {
		return "", db.Error
	}

	return name, nil
}
