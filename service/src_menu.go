package service

import (
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/utils"
)

// AddMenu
func AddMenu(forms []*utils.FormAddMenu) (ret []*models.Menu, err error) {
	db := models.DB.Begin()
	defer func() {
		if err != nil {
			db.Rollback()
			return
		}
		db.Commit()
	}()
	for _, v := range forms {
		var m = new(models.Menu)
		if v.ParentId != 0 {
			m1 := &models.Menu{}
			if err = m1.Get(int(v.ParentId)); err != nil {
				return
			}
		}
		m.ParentId = v.ParentId
		m.Icon = v.Icon
		m.Name = v.Name
		m.Router = v.Router
		if err = m.Insert(db); err != nil {
			return
		}
		ret = append(ret, m)
	}
	return
}

// GetMenu
func GetMenu(id int) (ret *models.Menu, err error) {
	var data = new(models.Menu)
	if err = data.Get(id); err != nil {
		return
	}
	ret = data
	return
}
