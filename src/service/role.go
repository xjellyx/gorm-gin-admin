package service

import (
	"github.com/olongfen/gorm-gin-admin/src/models"
	"github.com/olongfen/gorm-gin-admin/src/pkg/setting"
	"github.com/olongfen/gorm-gin-admin/src/utils"
)

func AddRole(f *utils.FormRole)(ret *models.Role,err error)  {
	var(
		data = new(models.Role)
	)
	data.Role = f.Role
	data.Level = f.Level
	if err = data.Insert();err!=nil{
		return
	}
	ret =data
	return
}

func DelRole(uid string,id int)(err error)  {
	var(
		data = new(models.Role)
		user = new(models.UserBase)
	)
	if err =user.GetByUId(uid);err!=nil{
		return
	}
	data.ID = uint(id)
	if err = data.Get();err!=nil{
		return
	}
	if data.Level>=user.Role.Level && user.Role.Level<setting.Setting.MaxRoleLevel  {
		err = utils.ErrActionNotAllow
		return
	}

	return data.Delete()
}

func UpdateRole(uid string,f *utils.FormUpdateRole)(err  error) {
	var(
		data = new(models.Role)
		user = new(models.UserBase)

	)
	if err =user.GetByUId(uid);err!=nil{
		return
	}
	if data.Level>=user.Role.Level && user.Role.Level<setting.Setting.MaxRoleLevel  {
		err = utils.ErrActionNotAllow
		return
	}

	data.Level = f.Level
	return data.Update(map[string]interface{}{"role":f.Role})
}

func GetRoleList()(ret []*models.Role,err error)  {
	return models.GetRoleList()
}