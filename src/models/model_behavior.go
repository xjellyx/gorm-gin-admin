package models

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/gorm-gin-admin/src/pkg/query"
	"github.com/olongfen/gorm-gin-admin/src/utils"
	"gorm.io/gorm"
	"strings"
)

type BehaviorRecord struct {
	Model
	Uid      string `json:"uid" gorm:"type:varchar(36); index"`
	Username string `json:"username" gorm:"type:varchar(16);index"`
	Behavior string `json:"behavior" gorm:"type: varchar(500)"`
	Method   string `json:"method" gorm:"type:varchar(12);index"`
	Path     string `json:"path" gorm:"type:varchar(120);index"`
	IP       string `json:"ip" gorm:"type:varchar(20);index"`
}

func NewActionRecord(s *session.Session, c *gin.Context, action string) *BehaviorRecord {
	d := &BehaviorRecord{
		Uid:      s.UID,
		Username: s.Content["username"].(string),
		Behavior: action,
		Method:   strings.ToLower(c.Request.Method),
		Path:     c.Request.URL.Path,
		IP:       c.Request.RemoteAddr,
	}
	return d
}

func (e *BehaviorRecord) Insert(dbs ...*gorm.DB) (err error) {
	if err = getDB(dbs...).Model(&BehaviorRecord{}).Create(e).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrInsertDataFailed
		return
	}
	return
}

func (e *BehaviorRecord) Get(dbs ...*gorm.DB) (err error) {
	if err = getDB(dbs...).Model(&BehaviorRecord{}).First(e.ID).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}

func DeleteBehaviorList(id []int64, dbs ...*gorm.DB) (err error) {
	return getDB(dbs...).Delete(&BehaviorRecord{}, "id in ?", id).Error
}

func GetBehaviorListByIDs(ids []int64, dbs ...*gorm.DB) (ret []*BehaviorRecord, err error) {
	if err = getDB(dbs...).Model(&BehaviorRecord{}).Where("id in ?", ids).Find(&ret).Error; err != nil {
		return
	}
	return
}

func GetBehaviorRecordList(q *query.Query) (ret []*BehaviorRecord, err error) {
	if err = DB.Model(&BehaviorRecord{}).Where(q.Cond, q.Values...).Offset(q.PageNum).Limit(q.PageSize).Order("id asc").Find(&ret).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}

func BehaviorCount() (ret int64, err error) {
	if err = DB.Model(&BehaviorRecord{}).Where("id > 0").Count(&ret).Error; err != nil {
		return
	}
	return
}
