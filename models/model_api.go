package models

import (
	"github.com/olongfen/user_base/pkg/query"
	"github.com/olongfen/user_base/utils"
	"gorm.io/gorm"
)

type APIGroup struct {
	gorm.Model
	Path        string `json:"path" gorm:"type:varchar(100);comment:'api路径'"`
	Description string `json:"description" gorm:"type:varchar(64);comment:'api中文描述'"`
	ApiGroup    string `json:"apiGroup" gorm:"type:varchar(36);comment:'api组'"`
	Method      string `json:"method" gorm:"type:varchar(10);unique_index:idx_api_groups_path_method;comment:'方法'"`
	// CreateUID   string `json:"create_uid" gorm:"type:varchar(36)"`
}

func APIGroupTableName() string {
	return "api_groups"
}

// InsertAPIGroup 插入数据
func (a *APIGroup) InsertAPIGroup() (err error) {
	if err = db.Model(&APIGroup{}).Create(a).Error; err != nil {
		logModel.Errorln("[InsertAPIGroup] err: ", err)
		err = utils.ErrInsertDataFailed
		return
	}
	return
}

// UpdateAPIGroup 更新数据
func (a *APIGroup) UpdateAPIGroup(id int64, m interface{}) (err error) {
	if err = db.Table(APIGroupTableName()).Where("id = ?", id).Updates(m).Error; err != nil {
		logModel.Errorln("[UpdateAPIGroup] err: ", err)
		err = utils.ErrUpdateDataFailed
		return err
	}
	return
}

// DeleteAPIGroup 删除api数据
func (a *APIGroup) DeleteAPIGroup(id int64) (err error) {
	if err = db.Table(APIGroupTableName()).Where("id =  ?", id).Delete(a).Error; err != nil {
		logModel.Errorln("[DeleteAPIGroup] err: ", err)
		err = utils.ErrDeleteDataFailed
		return err
	}
	return
}

// GetAPIGroup 获取api数据
func (a *APIGroup) GetAPIGroup(id int64) (err error) {
	if err = db.Table(APIGroupTableName()).Where("id = ?", id).Error; err != nil {
		logModel.Errorln("[GetAPIGroup] err: ", err)
		err = utils.ErrGetDataFailed
		return err
	}
	return
}

// GetAPIGroupList 获取api列表
func GetAPIGroupList(q *query.Query) (ret []*APIGroup, err error) {
	if err = db.Table(APIGroupTableName()).Where(q.Cond, q.Values).Find(&ret).Error; err != nil {
		logModel.Errorln("[GetAPIGroupList] err: ", err)
		err = utils.ErrGetDataFailed
		return nil, err
	}
	return
}
