package models

import (
	"github.com/olongfen/user_base/utils"
	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(36)"`
	ParentId uint `json:"parentId"`
	Router  string `json:"router" gorm:"type:varchar(24)"`
	Icon  string `json:"icon" gorm:"type:varchar(36)"`
	Sorts int64 `json:"sorts"`
	Children []Menu `json:"children"`
}

func (m *Menu)Insert()(err error)  {
	if err = DB.Create(m).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrInsertDataFailed
		return
	}
		return
}

func (m *Menu)Update(id int)(err error)  {
	if err = DB.Model(m).Where("id = ?",id).Updates(m).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrUpdateDataFailed
		return
	}
	return
}

func (m *Menu)Get(id int)(err error)  {
	if err = DB.Model(m).Where("id = ?",id).First(m).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	if err = DB.Model(m).Where("parent_id = ?",m.ID).Find(&m.Children).Error;err!=nil && err!=gorm.ErrRecordNotFound{
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return err
	}
	return
}

func (m *Menu)Delete(id int)(err error)  {
	if err = DB.Model(m).Where("id = ?",id).Delete(m).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}