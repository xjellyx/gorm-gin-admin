package models

import (
	"bytes"
	"fmt"
	"github.com/olongfen/user_base/src/utils"
	"gorm.io/gorm"
)

type APIGroup struct {
	Model
	Path        string `json:"path" gorm:"type:varchar(100);comment:'api路径'"`
	Description string `json:"description" gorm:"type:varchar(64);comment:'api中文描述'"`
	ApiGroup    string `json:"apiGroup" gorm:"type:varchar(36);comment:'api组'"`
	Method      string `json:"method" gorm:"type:varchar(10);unique_index:idx_api_groups_path_method;comment:'方法'"`
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

// GetAPIGroupList 获取api列表
func GetAPIGroupList() (ret []*APIGroup, err error) {
	if err = DB.Table(APIGroupTableName()).Where("id > ?", 0).Find(&ret).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return nil, err
	}
	return
}

// BatchInsertAPIGroup
func BatchInsertAPIGroup(datas []*APIGroup) (err error) {
	var buffer bytes.Buffer
	sql := "insert into `api_groups` (`path`,`description`,`api_group`,`method`) values"
	if _, err := buffer.WriteString(sql); err != nil {
		return err
	}
	for i, v := range datas {
		if i == len(datas)-1 {
			buffer.WriteString(fmt.Sprintf("('%s','%s',%s,%s);", v.Path, v.Description, v.ApiGroup, v.Method))
		} else {
			buffer.WriteString(fmt.Sprintf("('%s','%s',%s,%s),", v.Path, v.Description, v.ApiGroup, v.Method))
		}
	}
	if err = DB.Exec(buffer.String()).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrInsertDataFailed
		return
	}
	return
}
