package srv_user

import (
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// UserLogin 用户登入
func UserLogin(f *utils.LoginForm, isAdmin bool) (token string, err error) {
	if err = f.Valid(); err != nil {
		return
	}
	var (
		data = &models.UserBase{}
		s    = new(session.Session)
	)
	if err = data.GetUserByUsername(f.Username); err != nil {
		return
	}
	// 验证密码是否正确
	if err = bcrypt.CompareHashAndPassword([]byte(data.LoginPasswd), []byte(f.Password)); err != nil {
		return
	}

	s.Password = data.LoginPasswd
	if f.DeviceId != nil {
		s.ID = *f.DeviceId
	}
	s.UID = data.Uid
	s.IP = f.IP
	n := time.Now()
	s.CreateTime = n.Unix()
	s.ExpireTime = n.Add(session.SessionExpMaxSecure).Unix()
	s.Level = session.SessionLevelSecure
	if !isAdmin {
		if token, err = models.UserKey.SessionEncode(s); err != nil {
			return
		}
	} else {
		if token, err = models.AdminKey.SessionEncode(s); err != nil {
			return
		}
	}
	if err = data.UpdateUserOneColumn(s.UID, "status", 1); err != nil {
		return
	}
	//var (
	//	dataOnline *UserOnline
	//)
	//if dataOnline, err = PubUserOnlineGet(data.Uid); err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		dataOnline = new(UserOnline)
	//		dataOnline.Uid = s.UID
	//		if f.DeviceId != nil {
	//			dataOnline.Device = *f.DeviceId
	//		}
	//		dataOnline.IsOnline = true
	//		dataOnline.LoginIp = f.IP
	//		dataOnline.LoginTime = time.Now()
	//		if err = Database.Table(dataOnline.TableName()).Create(dataOnline).Error; err != nil {
	//			return
	//		}
	//	}
	//	return
	//}
	//
	//isOnline := true
	//if dataOnline, err = PubUserOnlineUpdate(s.UID, &utils.FormUserOnline{
	//	LoginTime: dataOnline.LoginTime.Unix(),
	//	IsOnline:  &isOnline,
	//}); err != nil {
	//	return
	//}
	//// 缓存token
	//if red, _err := GetRedisClient(); _err != nil {
	//	err = _err
	//	return
	//} else {
	//	if err = red.Set("cache_token"+data.Uid, token, session.SessionExpMaxSecure).Err(); err != nil {
	//		return
	//	}
	//	red.Close()
	//}

	return
}
