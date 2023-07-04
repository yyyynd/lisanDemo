package dataAccess

import "time"

type KnowledgePointsInfo struct {
	KpId       string     `gorm:"column:primaryKey;kp_id" json:"KpId"`  //type:string       comment:知识点id;knowledge_points_info主键列           version:2023-05-27 15:50
	KpContent  string     `gorm:"column:kp_content" json:"KpContent"`   //type:string       comment:知识点内容                                     version:2023-05-27 15:50
	Section    string     `gorm:"column:section" json:"Section"`        //type:string       comment:所属章节                                       version:2023-05-27 15:50
	UpdateTime *time.Time `gorm:"column:update_time" json:"UpdateTime"` //type:*time.Time   comment:最后更新时间;YYYY-MM-DD                        version:2023-05-27 15:50
	DeleteMark *int       `gorm:"column:delete_mark" json:"DeleteMark"` //type:*int         comment:逻辑删除标记;0代表未删除，1代表删除，默认值0   version:2023-05-27 15:50
}

// TableName 表名:knowledge_points_info，知识点信息表。
func (KnowledgePointsInfo) TableName() string {
	return "knowledge_points_info"
}
