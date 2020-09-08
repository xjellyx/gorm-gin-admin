package models

import (
	"github.com/olongfen/contrib/session"
	"gorm.io/gorm"
)

type KV struct {
	Label string `json:"label"`
	Value int `json:"value"`
}

func getDB(options ...*gorm.DB) *gorm.DB {
	var (
		db *gorm.DB
	)
	if len(options) > 0 {
		db = options[0]
	} else {
		db = DB
	}
	return db
}

// SessionCheck
func SessionCheck(s *session.Session) (err error) {
	//var (
	//	data *UserBase
	//)
	////
	//if s == nil {
	//	err = contrib.ErrSessionUidUndefined
	//	return
	//}
	//
	////
	//if data, err = PubUserGet(s.UID); err != nil {
	//	log.Errorln("[SessionCheck] err: ", err)
	//	err = contrib.ErrSessionUidUndefined
	//	return
	//}
	//
	//if s.Password != data.LoginPassword {
	//	s.Password = data.LoginPassword
	//}
	//
	//if s.ExpireTime <= time.Now().Unix() {
	//	err = contrib.ErrSessionExpired
	//	return
	//}

	return
}

//// TokenCheck
//func TokenCheck(token string, isAdmin bool) (err error) {
//	var (
//		s          *session.Session
//		rdb        *redis.Client
//		cacheToken string
//		dataOnline *UserOnline
//	)
//	if isAdmin {
//		if s, err = AdminKey.SessionDecodeAuto(token); err != nil {
//			return
//		}
//	} else {
//		if s, err = UserKey.SessionDecodeAuto(token); err != nil {
//			return
//		}
//	}
//	if rdb, err = GetRedisClient(); err != nil {
//		return
//	}
//	defer rdb.Close()
//
//	if cacheToken, err = rdb.Get("cache_token" + s.UID).Result(); err != nil {
//		return
//	}
//
//	if dataOnline, err = PubUserOnlineGet(s.UID); err != nil {
//		return
//	}
//	// 不是在线状态，退出，不给用户进行下一步的操作
//	if !dataOnline.IsOnline {
//		rdb.Del("cache_token" + s.UID)
//		err = utils.ErrUserNotOnline
//		return
//	}
//	// TODO 用户ip地址改变警告
//	// 不是缓存的token,是非法token
//	if cacheToken != token {
//		err = contrib.ErrTokenInvalid
//		return
//	}
//
//	return
//}
//
//// Login 用户登入
//func Login(f *utils.LoginForm) (token string, err error) {
//	if err = f.Valid(); err != nil {
//		return
//	}
//	var (
//		data *UserDetail
//		s    = new(session.Session)
//	)
//	if data, err = PubUserGetByUsername(f.Username); err != nil {
//		return
//	}
//	// 验证密码是否正确
//	if err = bcrypt.CompareHashAndPassword([]byte(data.LoginPassword), []byte(f.Password)); err != nil {
//		return
//	}
//
//	s.Password = data.LoginPassword
//	if f.DeviceId != nil {
//		s.Value = *f.DeviceId
//	}
//	s.UID = data.RoleRefer
//	s.IP = f.IP
//	n := time.Now()
//	s.CreateTime = n.Unix()
//	s.ExpireTime = n.Add(session.SessionExpMaxSecure).Unix()
//	s.Level = session.SessionLevelSecure
//	if token, err = UserKey.SessionEncodeAuto(s); err != nil {
//		return
//	}
//	if err = Database.Model(&UserDetail{}).Where("uid = ?", data.RoleRefer).Update("status", UserStatusNormal).Error; err != nil {
//		return
//	}
//
//	var (
//		dataOnline *UserOnline
//	)
//	if dataOnline, err = PubUserOnlineGet(data.RoleRefer); err != nil {
//		if err == gorm.ErrRecordNotFound {
//			dataOnline = new(UserOnline)
//			dataOnline.RoleRefer = s.UID
//			if f.DeviceId != nil {
//				dataOnline.Device = *f.DeviceId
//			}
//			dataOnline.IsOnline = true
//			dataOnline.LoginIp = f.IP
//			dataOnline.LoginTime = time.Now()
//			if err = Database.Table(dataOnline.APIGroupTableName()).Create(dataOnline).Error; err != nil {
//				return
//			}
//		}
//		return
//	}
//
//	isOnline := true
//	if dataOnline, err = PubUserOnlineUpdate(s.UID, &utils.FormUserOnline{
//		LoginTime: dataOnline.LoginTime.Unix(),
//		IsOnline:  &isOnline,
//	}); err != nil {
//		return
//	}
//	// 缓存token
//	if red, _err := GetRedisClient(); _err != nil {
//		err = _err
//		return
//	} else {
//		if err = red.Set("cache_token"+data.RoleRefer, token, session.SessionExpMaxSecure).Err(); err != nil {
//			return
//		}
//		red.Close()
//	}
//
//	return
//}
//
//func Logout(uid string) (err error) {
//	if err = UserOfflineDo(uid); err != nil {
//		return
//	}
//	return
//}
//
//// GetUserToken 获取用户token
//func GetUserToken(f *pb.ArgLogin) (ret string, uid string, err error) {
//	if f == nil {
//		err = utils.ErrFormParamInvalid
//		return
//	}
//	var (
//		data  *UserDetail
//		token string
//		s     = &session.Session{}
//	)
//	if len(f.RoleRefer) > 0 {
//		if data, err = PubUserGet(f.RoleRefer); err != nil {
//			return
//		}
//	} else if len(f.Username) > 0 {
//		if data, err = PubUserGetByUsername(f.Username); err != nil {
//			return
//		}
//	} else if len(f.Phone) > 0 {
//		if data, err = PubUserGetByPhone(f.LocNum, f.Phone); err != nil {
//			return
//		}
//	} else if len(f.Email) > 0 {
//		if data, err = PubUserGetByEmail(f.Email); err != nil {
//			return
//		}
//	}
//
//	// 验证密码是否正确
//	if err = bcrypt.CompareHashAndPassword([]byte(data.LoginPassword), []byte(f.Password)); err != nil {
//		return
//	}
//
//	// 验证账号状态,被冻结返回
//	if data.Status == UserStatusFreeze {
//		err = utils.ErrUserAccountFroze
//		return
//	}
//
//	s.Password = data.LoginPassword
//	s.Value = f.Device
//	s.UID = data.RoleRefer
//	s.IP = f.Ip
//	n := time.Now()
//	s.CreateTime = n.Unix()
//	s.ExpireTime = n.Add(session.SessionExpMaxSecure).Unix()
//	s.Level = session.SessionLevelSecure
//
//	//
//	if data.RoleRefer {
//		if token, err = AdminKey.SessionEncodeAuto(s); err != nil {
//			return
//		}
//	} else {
//		if token, err = UserKey.SessionEncodeAuto(s); err != nil {
//			return
//		}
//	}
//	if data.Status != UserStatusNormal && data.Status != UserStatusFreeze {
//		if err = Database.Model(&UserDetail{}).Where("uid = ?", data.RoleRefer).Update("status", UserStatusNormal).Error; err != nil {
//			return
//		}
//	}
//
//	// 缓存token
//	if red, _err := GetRedisClient(); _err != nil {
//		err = _err
//		return
//	} else {
//		if err = red.Set("cache_token"+data.RoleRefer, token, session.SessionExpMaxSecure).Err(); err != nil {
//			return
//		}
//		red.Close()
//	}
//
//	// 创建或者更新用户在线表
//	var (
//		dataOnline = new(UserOnline)
//	)
//	dataOnline.IsOnline = true
//	dataOnline.LoginIp = f.Ip
//	dataOnline.Device = f.Device
//	dataOnline.LoginTime = time.Now()
//	if err = Database.Model(&UserOnline{}).First(dataOnline, "uid = ?", data.RoleRefer).Error; err != nil {
//		if err == gorm.ErrRecordNotFound {
//			dataOnline.RoleRefer = data.RoleRefer
//			if err = Database.Model(&UserOnline{}).Create(dataOnline).Error; err != nil {
//				return
//			}
//
//		}
//	} else {
//		if err = Database.Model(&UserOnline{}).Where("uid=?", uid).Updates(dataOnline).Error; err != nil {
//			return
//		}
//	}
//
//	//
//	ret = token
//	uid = data.RoleRefer
//	return
//}
//
//// TokenDecodeSession
//func TokenDecodeSession(token interface{}, isAdmin bool) (ret *session.Session, err error) {
//	if isAdmin {
//		return AdminKey.SessionDecodeAuto(token)
//	} else {
//		return UserKey.SessionDecodeAuto(token)
//	}
//
//}
//
//func UserOfflineDo(uid string) (err error) {
//	var (
//		data *UserOnline
//		b    = false
//	)
//	if data, err = PubUserOnlineGet(uid); err != nil {
//		return
//	}
//	if data, err = PubUserOnlineUpdate(data.RoleRefer, &utils.FormUserOnline{
//		IsOnline:    &b,
//		OfflineTime: time.Now().Unix(),
//	}); err != nil {
//		return
//	}
//	// 删除缓存token
//	if red, _err := GetRedisClient(); _err != nil {
//		err = _err
//		return
//	} else {
//		if err = red.Del("cache_token" + data.RoleRefer).Err(); err != nil {
//			return
//		}
//		red.Close()
//
//	}
//	return
//}
//
//func GetOnlineUser(uid string) (ret int64, err error) {
//	var (
//		data   []*UserOnline
//		rdb    *redis.Client
//		uidArr []string
//		count  int64
//	)
//	if err = Database.Table("user_online").Where("is_online = ?", true).Find(&data).Error; err != nil {
//		return
//	}
//	rdb, err = GetRedisClient()
//	if err != nil {
//		return
//	}
//	defer rdb.Close()
//	for _, v := range data {
//		uidArr = append(uidArr, "cache_token"+v.RoleRefer)
//	}
//	count, _ = rdb.Exists(uidArr...).Result()
//
//	ret = count
//	return
//}
