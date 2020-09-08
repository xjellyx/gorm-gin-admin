package models

import (
	"github.com/olongfen/gorm-gin-admin/src/utils"
	"gorm.io/gorm"
)

type RoleAPI struct {
	ID     uint   `json:"id" gorm:"column:id"`
	PType  string `json:"pType" gorm:"column:p_type"`
	Role   string `json:"role" gorm:"column:v0"`
	Path   string `json:"path" gorm:"column:v1"`
	Method string `json:"method" gorm:"column:v2"`
}

func RuleAPITableName() string {
	return "casbin_rule"
}

func (c *RoleAPI) Insert(options ...*gorm.DB) (err error) {
	if err = getDB(options...).Table(RuleAPITableName()).Create(c).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return err
	}
	return
}

func (c *RoleAPI) Update(path string, method string, m map[string]interface{}, options ...*gorm.DB) (err error) {
	if err = getDB(options...).Table(RuleAPITableName()).Updates(m).Where("v1 = ? and v2 = ?", path, method).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrUpdateDataFailed
		return err
	}
	return
}

func (c *RoleAPI) Delete(id int64, options ...*gorm.DB) (err error) {
	if err = getDB(options...).Table(RuleAPITableName()).Delete(c, "id = ?", id).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrDeleteDataFailed
		return err
	}
	return
}

func GetRoleAPIListByRole(role string) (ret []*RoleAPI, err error) {
	if err = DB.Table(RuleAPITableName()).Where("v0 = ?", role).Find(&ret).Order("id asc").Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return nil, err
	}
	return
}

//
func (c *RoleAPI) DeleteByPathAndMethod(path string, method string, options ...*gorm.DB) (err error) {
	if err = getDB(options...).Table(RuleAPITableName()).Delete(c, "v1 = ? and v2 = ?", path, method).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrDeleteDataFailed
		return err
	}
	return
}


func (c *RoleAPI) GetByPathAndMethodAndRole(path string, method string, role string,options ...*gorm.DB) (err error) {
	if err = getDB(options...).Table(RuleAPITableName()).First(c, "v0 = ? and v1 = ? and v2 = ?", role,path, method).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrDeleteDataFailed
		return err
	}
	return
}