package models

import (
	"github.com/olongfen/gorm-gin-admin/src/utils"
	"gorm.io/gorm"
	"strconv"
)

type Role struct {
	Model
	Role     string `json:"role" gorm:"type:varchar(36);uniqueIndex"`
	Level    string `json:"level" gorm:"type:varchar(1)"`
}

func (r *Role)GetLevelMust() int {
	level,_:=strconv.Atoi(r.Level)
	return level
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
	if err = getDB(dbs...).Model(r).Updates(data).Error;err!=nil{
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