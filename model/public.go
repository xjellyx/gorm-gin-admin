package model

import (
	"github.com/srlemon/contrib"
	"github.com/srlemon/contrib/log"
	"github.com/srlemon/contrib/session"
	"time"
)

// TimeData 时间信息
type TimeData struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// SessionCheck
func SessionCheck(s *session.Session) (err error) {
	var (
		data *UserDetail
	)
	//
	if s == nil {
		err = contrib.ErrSessionUidUndefined
		return
	}

	//
	if data, err = PubUserGet(s.UID); err != nil {
		log.Errorln("[SessionCheck] err: ", err)
		err = contrib.ErrSessionUidUndefined
		return
	}

	if s.Password != data.LoginPassword {
		s.Password = data.LoginPassword
	}

	if s.ExpireTime <= time.Now().Unix() {
		err = contrib.ErrSessionExpired
		return
	}

	return
}
