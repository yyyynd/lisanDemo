package dataAccess

import (
	demoServer "demoProject/biz/model/demoServer"
	"time"
)

type KnowledgePointsRel struct {
	KprId      string     `gorm:"column:primaryKey;kpr_id" json:"KprId"` //type:string       comment:知识点关系记录id;knowledge_ points_rel主键列，使用自增列，无需手动指定   version:2023-05-27 17:55
	KpId       string     `gorm:"column:kp_id" json:"KpId"`
	ParentKpId string     `gorm:"column:parent_kp_id" json:"ParentKpId"` //type:string       comment:父级知识点的id;对应k_points_info中的kp_id列                              version:2023-05-27 17:55
	ChildKpId  string     `gorm:"column:child_kp_id" json:"ChildKpId"`   //type:string       comment:子级知识点的id;对应k_points_info中的kp_id列                              version:2023-05-27 17:55
	UpdateTime *time.Time `gorm:"column:update_time" json:"UpdateTime"`  //type:*time.Time   comment:更新时间;YYYY-MM-DD                                                      version:2023-05-27 17:55
	DeleteMark *int       `gorm:"column:delete_mark" json:"DeleteMark"`  //type:*int         comment:逻辑删除标记;0代表未删除，1代表删除，默认值0                             version:2023-05-27 17:55
}

func (KnowledgePointsRel) TableName() string {
	return "knowledge_points_rel"
}

func GetTreeStructure() ([]*demoServer.TreeStructureRespData, error) {
	//connect
	db, err := InitConnection(USER, PASSWD, "", "lisandb")
	if err != nil {
		return nil, err
	}
	infos := make([]KnowledgePointsRel, 0)
	//query
	db.Model(&KnowledgePointsRel{}).Select([]string{"kp_id", "child_kp_id"}).Find(&infos)
	if db.Error != nil {
		return nil, db.Error
	}

	//organize data structure
	data := make([]*demoServer.TreeStructureRespData, 0)
	for i := 0; i < len(infos); i++ {
		newData := demoServer.NewTreeStructureRespData()
		newData.ID = infos[i].KpId
		//query knowledge content
		db.Model(KnowledgePointsInfo{}).Select("kp_content").Where("kp_id = ? ", newData.ID).Find(&newData.Name)
		if db.Error != nil {
			return nil, db.Error
		}

		if infos[i].ChildKpId != "0" {
			for ; i < len(infos); i++ {
				newData.Children = append(newData.Children, infos[i].ChildKpId)
				if i == len(infos)-1 || infos[i].KpId != infos[i+1].KpId {
					break
				}
			}
		}
		data = append(data, newData)
	}

	return data, nil
}
