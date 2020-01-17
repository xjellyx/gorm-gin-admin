package model

import (
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	userBase "github.com/srlemon/userDetail"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type UserStatus int

const (
	UserStatusNormal UserStatus = iota
	UserStatusNotActive
	UserStatusFreeze
)

func (u *UserDetail) TableName() string {
	return "user_detail"
}

// UserDetail
type UserDetail struct {
	Uid              string         `json:"uid" gorm:"primary_key;size:36;index"`
	Nickname         string         `json:"nickname" gorm:"size:24;index"`
	Username         string         `json:"username" gorm:"size:24;unique_index"`
	Status           UserStatus     `json:"status" gorm:"index;default(0)"` // 用户状态,0初始化状态,1账号未激活,2账号冻结
	IsChangeUsername bool           `json:"isChangeUsername"gorm:"default:false"`
	RealName         string         `json:"realName" gorm:"size:36;index"`
	LocNum           string         `json:"locNum" gorm:"unique_index:loc_num_phone;size:6"` // 	所在地区(国家),国际电话区号,不带"+"号
	Phone            string         `json:"phone"gorm:"size:11;unique_index:loc_num_phone"`
	HeadIcon         string         `json:"headIcon"gorm:"size:256"`
	Email            string         `json:"email"gorm:"size:64;index"`
	LoginPassword    string         `json:"loginPassword" gorm:"size:128;column:loginPassword"`
	PayPassword      string         `json:"payPassword"gorm:"size:128"`
	Sex              int            `json:"sex" gorm:"index;size:2"`    // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Birth            string         `json:"birth" gorm:"index;size:32"` // 生日
	Sign             string         `json:"sign" gorm:"size:64"`
	Role             pq.StringArray `json:"role" gorm:"not null;type:varchar(36)[]"` // 角色, 用户可以拥有多个角色
	Level            int32          `json:"level" gorm:"default:1"`                  // 等级
	Secret           string         `json:"secret" gorm:"not null;type:varchar(32)"` // 用户自己的密钥
	Ethnic           string         `json:"ethnic" gorm:"index;size:24"`             // 民族
	Religion         string         `json:"religion" gorm:"index;size:24"`           // 宗教
	TimeData
	// 外键关联
	BankCards []BankCard      `json:"bankCards" gorm:"foreignkey:Uid"`
	Addr      []AddressDetail `json:"addr" gorm:"foreignkey:Uid"`
}

// NewUser
func NewUser(u *UserDetail) (ret *UserDetail) {
	if u != nil {
		ret = u
	} else {
		ret = new(UserDetail)
	}
	//
	if len(ret.Uid) != 36 {
		ret.Uid = uuid.NewV4().String()
	}
	//
	if len(ret.Secret) == 0 {
		ret.Secret = ""
	}
	ret.Role = []string{
		"a",
	}
	if len(ret.LocNum) == 0 {
		ret.LocNum = "86"
	}
	return
}

// PubUserGet
func PubUserGet(uid string) (ret *UserDetail, err error) {
	ret = new(UserDetail)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err = Database.Model(ret).Where("uid = ?", uid).Find(ret).Related(new(BankCard)).Error; err != nil {
			return
		}
	}()
	go func() {
		defer wg.Done()
		if err = Database.Model(ret).Where("uid = ?", uid).Find(ret).Related(new(AddressDetail)).Error; err != nil {
			return
		}
	}()
	wg.Wait()
	return
}

// PubUserAdd 添加用户
func PubUserAdd(f *userBase.FormRegister) (ret *UserDetail, err error) {
	if err = f.Valid(); err != nil {
		return
	}
	var (
		data = NewUser(nil)
		pwd  []byte
	)
	data.Phone = f.Phone
	if pwd, err = bcrypt.GenerateFromPassword([]byte(f.Password), bcrypt.DefaultCost); err != nil {
		return
	}
	data.LoginPassword = string(pwd)
	// TODO 随机生成数据

	if err = Database.Model(new(UserDetail)).Create(data).Error; err != nil {
		return
	}

	ret = data
	return
}

// PubUserDel 删除用户,软删除
func PubUserDel(uid string) (err error) {
	var (
		data = new(UserDetail)
	)
	data.Uid = uid
	if err = Database.Model(data).Association("Addr").Clear().Error; err != nil {
		return
	}
	if err = Database.Model(data).Association("BankCards").Clear().Error; err != nil {
		return
	}

	if err = Database.Delete(data).Error; err != nil {
		return
	}

	return
}

// PubUserUpdate
func PubUserUpdate(uid string) {

}
