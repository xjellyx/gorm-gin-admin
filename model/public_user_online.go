package model

import (
	base "github.com/olongfen/userDetail"
	"time"
)

// UserOnline 用户在线表
type UserOnline struct {
	Uid         string    `json:"uid" gorm:"primary_key;size:36;index"` // uid
	IsOnline    bool      `json:"isOnline" gorm:"index"`                // true: 在线
	LoginTime   time.Time `json:"loginTime" gorm:"index"`
	OfflineTime time.Time `json:"offlineTime" gorm:"index"`
	LoginIp     string    `json:"loginIp" gorm:"index"`
	Device      string    `json:"device" gorm:"index"`
	TimeData
}

func (u *UserOnline) TableName() string {
	return "user_online"
}

// TODO 获取所有在线用户
func NewUserOnline(in *UserOnline) (ret *UserOnline) {
	if in != nil {
		ret = in
	} else {
		ret = new(UserOnline)
	}
	return
}

// PubUserOnlineGet 获取在线用户
func PubUserOnlineGet(uid string) (ret *UserOnline, err error) {
	ret = new(UserOnline)
	if err = Database.Model(&UserOnline{}).First(ret, "uid=?", uid).Error; err != nil {
		return
	}
	return
}

func PubUserOnlineUpdate(uid string, f *base.FormUserOnline) (ret *UserOnline, err error) {
	if err = f.Valid(); err != nil {
		return
	}
	var (
		data *UserOnline
	)
	if data, err = PubUserOnlineGet(uid); err != nil {
		return
	}
	if f.IsOnline != nil {
		data.IsOnline = *f.IsOnline
	}
	if f.LoginTime > 0 {
		data.LoginTime = time.Unix(f.LoginTime, 0)
	}
	if len(f.LoginIp) > 0 {
		data.LoginIp = f.LoginIp
	}
	if f.OfflineTime > 0 {
		data.OfflineTime = time.Unix(f.OfflineTime, 0)
	}
	if len(f.Device) > 0 {
		data.Device = f.Device
	}
	if err = Database.Model(data).Updates(data).Error; err != nil {
		return
	}

	ret = data
	return
}
