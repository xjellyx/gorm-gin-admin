package models

import (
	"github.com/olongfen/gorm-gin-admin/src/pkg/query"
	"github.com/olongfen/gorm-gin-admin/src/utils"
	"gorm.io/gorm"
)

type EditRecord struct {
	Model
	Uid    string `json:"uid" gorm:"type:varchar(36); index"`
	Action string `json:"action" gorm:"type: varchar(500)"`
	From   string `json:"from" gorm:"type:varchar(36); index"`
}

func (e *EditRecord) Insert(dbs ...*gorm.DB) (err error) {
	if err = getDB(dbs...).Model(&EditRecord{}).Create(e).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrInsertDataFailed
		return
	}
	return
}

func (e *EditRecord) Get(dbs ...*gorm.DB) (err error) {
	if err = getDB(dbs...).Model(&EditRecord{}).First(e.ID).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}

func GetEditRecordList(q *query.Query) (ret []*EditRecord, err error) {
	if err = DB.Model(&EditRecord{}).Where(q.Cond, q.Values...).Offset(q.PageNum).Limit(q.PageSize).Order("id asc").Find(&ret).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}
