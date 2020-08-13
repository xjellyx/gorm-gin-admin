package models

import (
	"github.com/olongfen/user_base/utils"
	"gorm.io/gorm"
)

type Menu struct {
	Model
	Name     string `json:"name" gorm:"type:varchar(36)"`
	ParentId uint   `json:"parentId"`
	Router   string `json:"router" gorm:"type:varchar(24)"`
	Icon     string `json:"icon" gorm:"type:varchar(36)"`
	Sorts    int64  `json:"sorts"`
	Children []Menu `json:"children" gorm:"-"`
}

func (m *Menu) Insert(options ...*gorm.DB) (err error) {
	if err = getDB(options...).Create(m).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrInsertDataFailed
		return
	}
	return
}

func (m *Menu) Update(id int, options ...*gorm.DB) (err error) {
	if err = getDB(options...).Model(m).Where("id = ?", id).Updates(m).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrUpdateDataFailed
		return
	}
	return
}

func (m *Menu) Get(id int) (err error) {
	if err = DB.Model(m).Where("id = ?", id).First(m).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	if err = DB.Model(m).Where("parent_id = ?", m.ID).Find(&m.Children).Error; err != nil && err != gorm.ErrRecordNotFound {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return err
	}
	return
}

func (m *Menu) Delete(id int, options ...*gorm.DB) (err error) {
	if err = getDB(options...).Model(m).Where("id = ?", id).Delete(m).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}
