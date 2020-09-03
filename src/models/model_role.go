package models

import (
	"github.com/olongfen/gorm-gin-admin/src/utils"
	"gorm.io/gorm"
)

type Role struct {
	Model
	Role     string `json:"role" gorm:"varchar(36);uniqueIndex"`
	Level    int `json:"level" gorm:"default:0"`
}

// Insert
func (r *Role)Insert(dbs ...*gorm.DB)(err error)  {
	if err = getDB(dbs...).Model(r).Create(r).Error;err!=nil{
		logModel.Errorln(err)
		err = utils.ErrInsertDataFailed
		return
	}
	return
}

// Delete
func (r *Role)Delete(dbs ...*gorm.DB)(err error)  {
	if err = getDB(dbs...).Model(r).Delete(r,"id = ?",r.ID).Error;err!=nil{
		logModel.Errorln(err)
		err = utils.ErrDeleteDataFailed
		return
	}
	return
}

// Get
func (r *Role)Get(dbs ...*gorm.DB)(err error)  {
	if err = getDB(dbs...).Model(r).First(r,"id = ?",r.ID).Error;err!=nil{
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}

// Get
func (r *Role)GetByRole(role string,dbs ...*gorm.DB)(err error)  {
	if err = getDB(dbs...).Model(r).First(r,"role = ?",role).Error;err!=nil{
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}

// Get
func (r *Role)Update(data interface{},dbs ...*gorm.DB)(err error)  {
	if err = getDB(dbs...).Model(r).Where("id = ?",r.ID).Updates(data).Error;err!=nil{
		logModel.Errorln(err)
		err = utils.ErrUpdateDataFailed
		return
	}
	return
}

// GetRoleList
func GetRoleList()(ret []*Role,err error)  {
	if err =DB.Model(&Role{}).Find(&ret).Error;err!=nil{
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}