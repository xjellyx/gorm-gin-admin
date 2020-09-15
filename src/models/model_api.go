package models

import (
	"fmt"
	"github.com/olongfen/gorm-gin-admin/src/pkg/query"
	"github.com/olongfen/gorm-gin-admin/src/utils"
	"gorm.io/gorm"
)

type APIGroup struct {
	Model
	Path        string `json:"path" gorm:"type:varchar(100);uniqueIndex:idx_api_groups_path_method;comment:'api路径'"`
	Description string `json:"description" gorm:"type:varchar(64);comment:'api中文描述'"`
	ApiGroup    string `json:"apiGroup" gorm:"type:varchar(36);comment:'api组'"`
	Method      string `json:"method" gorm:"type:varchar(10);uniqueIndex:idx_api_groups_path_method;comment:'方法'"`
	// CreateUID   string `json:"create_uid" gorm:"type:varchar(36)"`
}

func APIGroupTableName() string {
	return "api_groups"
}

func NewAPIGroup() *APIGroup {
	return new(APIGroup)
}

// Insert 插入数据
func (a *APIGroup) Insert(options ...*gorm.DB) (err error) {
	if err = getDB(options...).Model(&APIGroup{}).Create(a).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrInsertDataFailed
		return
	}
	return
}

// Update 更新数据
func (a *APIGroup) Update(id int64, m interface{}, options ...*gorm.DB) (err error) {
	if err = getDB(options...).Table(APIGroupTableName()).Where("id = ?", id).Updates(m).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrUpdateDataFailed
		return err
	}
	return
}

// Delete 删除api数据
func (a *APIGroup) Delete(id int64, options ...*gorm.DB) (err error) {
	if err = getDB(options...).Table(APIGroupTableName()).Where("id =  ?", id).Delete(a).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrDeleteDataFailed
		return err
	}
	return
}

// Get 获取api数据
func (a *APIGroup) Get(id int64) (err error) {
	if err = DB.First(a, "id = ?", id).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}

func (a *APIGroup) GetBPathAndMethod(path, method string) (err error) {
	if err = DB.First(a, "path = ? and method= ?", path, method).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}

// GetAPIGroupList 获取api列表
func GetAPIGroupList(q *query.Query) (ret []*APIGroup, err error) {
	if q != nil {
		fmt.Println(q)
		if err = DB.Table(APIGroupTableName()).Where(q.Cond, q.Values...).Offset(q.PageNum).Limit(q.PageSize).Order("id asc").Find(&ret).Error; err != nil {
			logModel.Errorln(err)
			err = utils.ErrGetDataFailed
			return nil, err
		}
	} else {
		if err = DB.Table(APIGroupTableName()).Order("id asc").Find(&ret).Error; err != nil {
			logModel.Errorln(err)
			err = utils.ErrGetDataFailed
			return nil, err
		}
	}
	return
}

// BatchInsertAPIGroup
func BatchInsertAPIGroup(datas []*APIGroup) (err error) {
	if err = DB.Model(&APIGroup{}).Create(datas).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrInsertDataFailed
		return
	}
	return
}
