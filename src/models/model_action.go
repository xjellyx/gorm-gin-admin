package models

import (
	"github.com/olongfen/gorm-gin-admin/src/pkg/query"
	"github.com/olongfen/gorm-gin-admin/src/utils"
	"gorm.io/gorm"
)

type ActionRecord struct {
	Model
	Uid    string `json:"uid" gorm:"type:varchar(36); index"`
	Username string `json:"username" gorm:"type:varchar(16);index"`
	Action string `json:"action" gorm:"type: varchar(500)"`
	From   string `json:"from" gorm:"type:varchar(36); index"`
	Method string `json:"method" gorm:"type:varchar(12);index"`
	Path  string `json:"path" gorm:"type:varchar(120);index"`
	IP string `json:"ip" gorm:"type:varchar(20);index"`
}

func (e *ActionRecord) Insert(dbs ...*gorm.DB) (err error) {
	if err = getDB(dbs...).Model(&ActionRecord{}).Create(e).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrInsertDataFailed
		return
	}
	return
}

func (e *ActionRecord) Get(dbs ...*gorm.DB) (err error) {
	if err = getDB(dbs...).Model(&ActionRecord{}).First(e.ID).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}

func GetEditRecordList(q *query.Query) (ret []*ActionRecord, err error) {
	if err = DB.Model(&ActionRecord{}).Where(q.Cond, q.Values...).Offset(q.PageNum).Limit(q.PageSize).Order("id asc").Find(&ret).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}
