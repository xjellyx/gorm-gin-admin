package models

import (
	"github.com/olongfen/user_base/pkg/query"
	"github.com/olongfen/user_base/utils"
	"gorm.io/gorm"
)

// UserCard
type UserCard struct {
	gorm.Model
	Uid string `json:"uid" gorm:"type:varchar(36); unique_index"`
	//
	Name        string `json:"name" `
	CardId      string `json:"cardId" gorm:"unique_index;type:varchar(18)" ` // 身份证号
	IssueOrg    string `json:"issueOrg" `                                    // 身份证发证机关
	Birthday    string `json:"birthday" gorm:"type:varchar(12)"`             // 出生日期
	ValidPeriod string `json:"validPeriod"  gorm:"type:varchar(24)"`         // 有效时期
	CardIdAddr  string `json:"cardIdAddr"  gorm:"type:varchar(64)"`          // 身份证地址
	Sex         int    `json:"sex" `
	Nation      string `json:"nation" `
}

// Insert 插入一条数据
func (u *UserCard) Insert(options ...*gorm.DB) (err error) {
	if err = getDB(options...).Create(u).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrInsertDataFailed
		return
	}
	return
}

// Update 更新数据
func (u *UserCard) Update(uid string, options ...*gorm.DB) (err error) {
	if err = getDB(options...).Model(u).Where("uid = ?", uid).Updates(u).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrUpdateDataFailed
		return
	}
	return
}

// UpdateOne 更新一个字段
func (u *UserCard) UpdateOne(uid string, column string, value interface{}, options ...*gorm.DB) (err error) {
	if err = getDB(options...).Model(u).Where("uid = ?", uid).Update(column, value).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrUpdateDataFailed
		return err
	}
	return
}

// GetById
func (u *UserCard) GetById(id int64) (err error) {
	if err = DB.Model(u).Where("id = ?", id).Find(&u).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return err
	}
	return
}

// GetByUid
func (u *UserCard) GetByUid(uid string) (err error) {

	if err = DB.Model(u).Where("uid = ?", uid).Find(u).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}

// GetByIDCard
func (u *UserCard) GetByIDCard(cardId string) (err error) {
	if err = DB.Model(u).Where("card_id = ?", cardId).Find(u).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}

// GetIDCardList
func GetIDCardList(q *query.Query) (ret []*UserCard, err error) {
	if err = DB.Model(&UserCard{}).Where(q.Cond, q.Values...).Offset(q.PageNum).Limit(q.PageSize).Find(&ret).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}

// GetUserIDCardTotal 获取总数
func GetUserIDCardTotal(cond interface{}) (ret int64, err error) {
	var count int64
	if err = DB.Model(&UserCard{}).Where(cond).Count(&count).Error; err != nil {
		logModel.Errorln(err)
		err = utils.ErrGetDataFailed
		return 0, err
	}

	return count, nil
}
