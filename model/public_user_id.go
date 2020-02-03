package model

import base "github.com/olefen/userDetail"

// IDCard
type IDCard struct {
	IDCard      string `json:"idCard" gorm:"primary_key;size:18;index" ` // 身份证号
	IssueOrg    string `json:"issueOrg" gorm:"not null"`                 // 身份证发证机关
	Uid         string `json:"uid"gorm:"index;size:36;unique_index"`
	Birthday    string `json:"birthday" gorm:"not null;size:32"` // 出生日期
	ValidPeriod string `json:"validPeriod" gorm:"not null"`      // 有效时期
	IDCardAddr  string `json:"idCardAddr" gorm:"not null"`       // 身份证地址
	Name        string `json:"name" gorm:"not null;index"`       // 姓名
	Nation      string `json:"nation" gorm:"not null; index"`    // 民族
	Sex         int    `json:"sex" gorm:"index;size:2"`          // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
}

func (i *IDCard) TableName() string {
	return "id_card"
}

// PubGetIDCard 获取身份证信息
func PubGetIDCard(uid string) (ret *IDCard, err error) {
	ret = new(IDCard)
	if err = Database.First(ret, "uid = ?", uid).Error; err != nil {
		return
	}
	return
}

// NewIDCard
func NewIDCard(in *IDCard) (ret *IDCard) {
	if in != nil {
		ret = in
	} else {
		ret = new(IDCard)
	}
	return
}

// PubIDCardAdd 添加身份证,实名验证
func PubIDCardAdd(uid string, f *base.FormIDCard) (ret *IDCard, err error) {
	if err = f.Valid(); err != nil {
		return
	}
	var (
		dataUser *UserDetail
		data     *IDCard
	)
	if dataUser, err = PubUserGet(uid); err != nil {
		return
	}
	// 已经进行实名验证
	if dataUser.Verified {
		err = base.ErrUserVerified
		return
	}
	data = NewIDCard(&IDCard{
		IDCard:      f.IDCard,
		IssueOrg:    f.IssueOrg,
		Uid:         dataUser.Uid,
		Birthday:    f.Birthday,
		ValidPeriod: f.ValidPeriod,
		IDCardAddr:  f.IDCardAddr,
		Name:        f.Name,
		Nation:      f.Nation,
		Sex:         f.Sex,
	})
	tx := Database.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		} else {
			tx.Commit()
			return
		}
	}()
	if err = tx.Create(data).Error; err != nil {
		return
	}
	// 实名登记
	if err = tx.Model(&UserDetail{}).Where("uid = ?", dataUser.Uid).Updates(map[string]interface{}{
		"verified":  true,
		"real_name": f.Name,
	}).Error; err != nil {
		return
	}

	//
	ret = data
	return
}

// TODO 更新删除接口
