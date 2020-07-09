package models

import "github.com/olongfen/user_base/utils"

type CasbinAPI struct {
	ID     uint   `json:"id" gorm:"column:id"`
	PType  string `json:"pType" gorm:"column:p_type"`
	Uid    string `json:"uid" gorm:"column:v0"`
	Path   string `json:"path" gorm:"column:v1"`
	Method string `json:"method" gorm:"column:v2"`
}

func CasbinAPITableName() string {
	return "casbin_apis"
}

func (c *CasbinAPI) InsertCasbinAPI() (err error) {
	if err = db.Table(CasbinAPITableName()).Create(c).Error; err != nil {
		logModel.Errorln("[InsertCasbinAPI] err: ", err)
		err = utils.ErrGetDataFailed
		return err
	}
	return
}

func (c *CasbinAPI) UpdateCasbinAPI(id int64, m interface{}) (err error) {
	if err = db.Table(CasbinAPITableName()).Updates(m).Where("id = ?", id).Error; err != nil {
		logModel.Errorln("[UpdateCasbinAPI] err: ", err)
		err = utils.ErrUpdateDataFailed
		return err
	}
	return
}

func (c *CasbinAPI) DeleteCasbinAPI(id int64) (err error) {
	if err = db.Table(CasbinAPITableName()).Delete(c, "id = ?", id).Error; err != nil {
		logModel.Errorln("[DeleteCasbinAPI] err: ", err)
		err = utils.ErrDeleteDataFailed
		return err
	}
	return
}

func GetCasbinAPIListByUID(uid string) (ret []*CasbinAPI, err error) {
	if err = db.Table(CasbinAPITableName()).Where("uid = ?", uid).Find(&ret).Order("id asc").Error; err != nil {
		logModel.Errorln("[GetCasbinAPIListByUID] err: ", err)
		err = utils.ErrGetDataFailed
		return nil, err
	}
	return
}
