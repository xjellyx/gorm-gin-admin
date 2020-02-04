package model

import (
	"github.com/go-redis/redis"
	"github.com/olefen/contrib"
	"github.com/olefen/contrib/log"
	"github.com/olefen/contrib/session"
	pb "github.com/olefen/model.grpc"
	base "github.com/olefen/userDetail"
	"golang.org/x/crypto/bcrypt"
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

// TokenCheck
func TokenCheck(token string) (err error) {
	var (
		s          *session.Session
		rdb        *redis.Client
		cacheToken string
	)
	if s, err = UserKey.SessionDecodeAuto(token); err != nil {
		return
	}
	if rdb, err = GetRedisClient(); err != nil {
		return
	}
	if cacheToken, err = rdb.Get(s.UID).Result(); err != nil {
		return
	}

	// 不是缓存的token,是非法token
	if cacheToken != token {
		err = contrib.ErrTokenInvalid
		return
	}

	return
}

// Login 用户登入
func Login(f *base.LoginForm) (ret *UserDetail, token string, err error) {
	if err = f.Valid(); err != nil {
		return
	}
	var (
		data *UserDetail
		s    = new(session.Session)
	)
	if data, err = PubUserGetByUsername(f.Username); err != nil {
		return
	}
	// 验证密码是否正确
	if err = bcrypt.CompareHashAndPassword([]byte(data.LoginPassword), []byte(f.Password)); err != nil {
		return
	}

	s.Password = data.LoginPassword
	if f.DeviceId != nil {
		s.ID = *f.DeviceId
	}
	s.UID = data.Uid
	s.IP = f.IP
	n := time.Now()
	s.CreateTime = n.Unix()
	s.ExpireTime = n.Add(session.SessionExpMaxSecure).Unix()
	s.Level = session.SessionLevelSecure
	if token, err = UserKey.SessionEncodeAuto(s); err != nil {
		return
	}
	if err = Database.Model(&UserDetail{}).Where("uid = ?", data.Uid).Update("status", UserStatusNormal).Error; err != nil {
		return
	}

	// 缓存token
	if red, _err := GetRedisClient(); _err != nil {
		err = _err
		return
	} else {
		if err = red.Set(data.Uid, token, session.SessionExpMaxSecure).Err(); err != nil {
			return
		}
	}

	ret = data
	return
}

// GetUserToken 获取用户token
func GetUserToken(f *pb.ArgLogin) (ret string, err error) {
	if f == nil {
		err = base.ErrFormParamInvalid
		return
	}
	var (
		data  *UserDetail
		token string
		s     = &session.Session{}
	)
	if len(f.Uid) > 0 {
		if data, err = PubUserGet(f.Uid); err != nil {
			return
		}
	} else if len(f.Username) > 0 {
		if data, err = PubUserGetByUsername(f.Username); err != nil {
			return
		}
	} else if len(f.Phone) > 0 {
		if data, err = PubUserGetByPhone(f.LocNum, f.Phone); err != nil {
			return
		}
	} else if len(f.Email) > 0 {
		if data, err = PubUserGetByEmail(f.Email); err != nil {
			return
		}
	}

	// 验证密码是否正确
	if err = bcrypt.CompareHashAndPassword([]byte(data.LoginPassword), []byte(f.Password)); err != nil {
		return
	}

	// 验证账号状态,被冻结返回
	if data.Status == UserStatusFreeze {
		err = base.ErrUserAccountFroze
		return
	}

	s.Password = data.LoginPassword
	s.ID = f.Device
	s.UID = data.Uid
	s.IP = f.Ip
	n := time.Now()
	s.CreateTime = n.Unix()
	s.ExpireTime = n.Add(session.SessionExpMaxSecure).Unix()
	s.Level = session.SessionLevelSecure

	//
	if data.IsAdmin {
		if token, err = AdminKey.SessionEncodeAuto(s); err != nil {
			return
		}
	} else {
		if token, err = UserKey.SessionEncodeAuto(s); err != nil {
			return
		}
	}
	if data.Status != UserStatusNormal && data.Status != UserStatusFreeze {
		if err = Database.Model(&UserDetail{}).Where("uid = ?", data.Uid).Update("status", UserStatusNormal).Error; err != nil {
			return
		}
	}

	// 缓存token
	if red, _err := GetRedisClient(); _err != nil {
		err = _err
		return
	} else {
		if err = red.Set(data.Uid, token, session.SessionExpMaxSecure).Err(); err != nil {
			return
		}
	}

	//
	ret = token
	return
}

// TokenDecodeSession
func TokenDecodeSession(token interface{}, isAdmin bool) (ret *session.Session, err error) {
	if isAdmin {
		return AdminKey.SessionDecodeAuto(token)
	} else {
		return UserKey.SessionDecodeAuto(token)
	}

}
