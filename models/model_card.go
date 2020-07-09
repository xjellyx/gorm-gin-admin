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

// InsertIDCard 插入一条数据
func (u *UserCard) InsertIDCard() (err error) {
	if err = db.Create(u).Error; err != nil {
		logModel.Errorln("[InsertIDCard] err: ", err)
		err = utils.ErrInsertDataFailed
		return
	}
	return
}

// UpdateIDCard 更新数据
func (u *UserCard) UpdateIDCard(uid string) (err error) {
	if err = db.Model(u).Where("uid = ?", uid).Updates(u).Error; err != nil {
		logModel.Errorln("[UpdateIDCard] err: ", err)
		err = utils.ErrUpdateDataFailed
		return
	}
	return
}

// UpdateIDCardOne 更新一个字段
func (u *UserCard) UpdateIDCardOne(uid string, column string, value interface{}) (err error) {
	if err = db.Model(u).Where("uid = ?", uid).Update(column, value).Error; err != nil {
		logModel.Errorln("[UpdateIDCardOne] err: ", err)
		err = utils.ErrUpdateDataFailed
		return err
	}
	return
}

// GetIDCardById
func (u *UserCard) GetIDCardById(id int64) (err error) {
	if err = db.Model(u).Where("id = ?", id).Find(&u).Error; err != nil {
		logModel.Errorln("[GetIDCardById] err: ", err)
		err = utils.ErrGetDataFailed
		return err
	}
	return
}

// GetIDCardByUid
func (u *UserCard) GetIDCardByUid(uid string) (err error) {

	if err = db.Model(u).Where("uid = ?", uid).Find(u).Error; err != nil {
		logModel.Errorln("[GetIDCardByUid] err: ", err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}

// GetIDCardByCardID
func (u *UserCard) GetIDCardByCardID(cardId string) (err error) {
	if err = db.Model(u).Where("card_id = ?", cardId).Find(u).Error; err != nil {
		logModel.Errorln("[GetIDCarsByCardID] err: ", err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}

// GetIDCardList
func GetIDCardList(q *query.Query) (ret []*UserCard, err error) {
	if err = db.Model(&UserCard{}).Where(q.Cond, q.Values...).Offset(q.PageNum).Limit(q.PageSize).Find(&ret).Error; err != nil {
		logModel.Errorln("[GetIDCardList] err: ", err)
		err = utils.ErrGetDataFailed
		return
	}
	return
}

// GetUserIDCardTotal 获取总数
func GetUserIDCardTotal(cond interface{}) (ret int64, err error) {
	var count int64
	if err = db.Model(&UserCard{}).Where(cond).Count(&count).Error; err != nil {
		logModel.Errorln("[GetUserIDCardTotal] err: ", err)
		err = utils.ErrGetDataFailed
		return 0, err
	}

	return count, nil
}
