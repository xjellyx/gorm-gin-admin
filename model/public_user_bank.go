package model

import (
	base "github.com/olefen/userDetail"
)

// BankCard 银行卡信息
type BankCard struct {
	Number    string `json:"number" gorm:"not null;primary_key;size:19;index"`
	Name      string `json:"name" gorm:"index;size:64"`       // 用户姓名
	Bank      string `json:"bank" gorm:"index;size:64"`       // 银行卡名称
	BankStart string `json:"bankStart" gorm:"index;size:128"` // 开户银行名称
	Uid       string `json:"uid"gorm:"index;size:36"`
	TimeData
}

// TableName
func (b *BankCard) TableName() string {
	return "bank_card"
}

func NewBankCard(in *BankCard) (ret *BankCard) {
	if in != nil {
		ret = in
	} else {
		ret = new(BankCard)
	}
	return
}

// PubBankCardAdd 添加银行卡信息
func PubBankCardAdd(uid string, f *base.FormBankCard) (ret *BankCard, err error) {
	if err = f.Valid(); err != nil {
		return
	}
	var (
		data     *BankCard
		dataUser *UserDetail
	)

	if dataUser, err = PubUserGet(uid); err != nil {
		return
	}

	// 还没有进行实名认证
	if !dataUser.Verified {
		err = base.ErrNotRealNameVerified
		return
	}
	// 银行家阿绑定数据上限
	if count := Database.Model(dataUser).Association("BankCards").Count(); count == 6 {
		err = base.ErrCapOfBankCard
		return
	}
	data = NewBankCard(&BankCard{
		Number:    f.Number,
		Name:      f.Name,
		Bank:      f.Bank,
		BankStart: f.BankStart,
		Uid:       uid,
	})

	if err = Database.Model(&BankCard{}).Create(data).Error; err != nil {
		return
	}

	ret = data
	return
}

// PubBankCardGet 获取银行卡信息
func PubBankCardGet(number string) (ret *BankCard, err error) {
	ret = new(BankCard)
	ret.Number = number
	if err = Database.Model(ret).First(ret).Error; err != nil {
		return
	}
	return
}

// PubBankCardDel 删除银行卡信息,硬删除
func PubBankCardDel(uid string, number string) (err error) {
	data := new(BankCard)
	if data, err = PubBankCardGet(number); err != nil {
		return
	}
	// 不是自己的信息
	if data.Uid != uid {
		err = base.ErrActionNotAllow
		return
	}
	if err = Database.Unscoped().Delete(data).Error; err != nil {
		return
	}

	return
}

// PubBankCardGetList 获取自己的银行卡列表
func PubBankCardGetList(uid string) (ret []BankCard, err error) {
	data := &UserDetail{Uid: uid}
	if err = Database.Model(data).Association("BankCards").Find(&ret).Error; err != nil {
		return
	}
	return
}
