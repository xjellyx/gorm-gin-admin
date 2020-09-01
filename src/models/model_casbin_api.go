package models

import (
	"github.com/olongfen/gorm-gin-admin/src/utils"
	"gorm.io/gorm"
)

type RuleAPI struct {
	ID     uint   `json:"id" gorm:"column:id"`
	PType  string `json:"pType" gorm:"column:p_type"`
	Uid    string `json:"uid" gorm:"column:v0"`
	Path   string `json:"path" gorm:"column:v1"`
	Method string `json:"method" gorm:"column:v2"`
}

func RuleAPITableName() string {
	return "casbin_rule"
}

func (c *RuleAPI) Insert(options ...*gorm.DB) (err error) {
	if err = getDB(options...).Table(RuleAPITableName()).Create(c).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return err
	}
	return
}

func (c *RuleAPI) Update(path string, method string, m map[string]interface{}, options ...*gorm.DB) (err error) {
	if err = getDB(options...).Table(RuleAPITableName()).Updates(m).Where("v1 = ? and v2 = ?", path, method).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrUpdateDataFailed
		return err
	}
	return
}

func (c *RuleAPI) Delete(id int64, options ...*gorm.DB) (err error) {
	if err = getDB(options...).Table(RuleAPITableName()).Delete(c, "id = ?", id).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrDeleteDataFailed
		return err
	}
	return
}

func GetRuleAPIListByUID(uid string) (ret []*RuleAPI, err error) {
	if err = DB.Table(RuleAPITableName()).Where("uid = ?", uid).Find(&ret).Order("id asc").Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return nil, err
	}
	return
}

//
func (c *RuleAPI) DeleteByPathAndMethod(path string, method string, options ...*gorm.DB) (err error) {
	if err = getDB(options...).Table(RuleAPITableName()).Delete(c, "v1 = ? and v2 = ?", path, method).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrDeleteDataFailed
		return err
	}
	return
}
