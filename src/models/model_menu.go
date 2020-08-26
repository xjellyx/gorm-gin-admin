package models

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/olongfen/user_base/src/utils"
	"gorm.io/gorm"
)

type Menu struct {
	Model
	Name      string   `json:"name" gorm:"type:varchar(36)"`
	ParentId  uint     `json:"parentId"`
	Path      string   `json:"path" gorm:"type:varchar(24)"`
	Component string   `json:"component" gorm:"type:varchar(36)"`
	Sorts     int64    `json:"sorts"`
	Meta      MenuMate `json:"meta" gorm:"type:json"`
	Children  []*Menu  `json:"children" gorm:"-"`
}

type MenuMate struct {
	Icon  string `json:"icon"`
	Title string `json:"title"`
}

func (m MenuMate) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *MenuMate) Scan(d interface{}) (err error) {
	return json.Unmarshal(d.([]byte), m)
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

func (m *Menu) Get(id int, options ...*gorm.DB) (err error) {
	if err = getDB(options...).Model(&Menu{}).First(m, "id = ?", id).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	if err = DB.Model(m).Where("parent_id = ?", m.ID).Find(&m.Children).Error; err != nil && err != gorm.ErrRecordNotFound {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return err
	} else {
		err = nil
	}
	GetChildren(m.Children)

	return
}

func GetChildren(Children []*Menu) (err error) {
	for _, v := range Children {
		if err = DB.Model(&Menu{}).Where("parent_id = ?", v.ID).Find(&v.Children).Error; err != nil && err != gorm.ErrRecordNotFound {
			logModel.Errorln(err)
			continue
		}
		if len(v.Children) > 0 {
			GetChildren(v.Children)
		}
	}
	err = nil
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

func GetMenuList(options ...*gorm.DB) (ret []*Menu, err error) {
	if err = getDB(options...).Model(&Menu{}).Where("id > 0").Find(&ret).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}
