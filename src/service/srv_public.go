package service

import (
	"github.com/dchest/captcha"
	"github.com/go-redis/redis"
	"github.com/olongfen/contrib/log"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/gorm-gin-admin/src/models"
	"github.com/olongfen/gorm-gin-admin/src/setting"
	"github.com/olongfen/gorm-gin-admin/src/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var (
	logServe = log.NewLogFile(log.ParamLog{
		Path:   setting.Settings.FilePath.LogDir + "/service",
		Stdout: setting.DevEnv,
		P:      setting.Settings.FilePath.LogPatent,
	})
)

// UserLogin 用户登入
func UserLogin(f *utils.LoginForm, isAdmin bool) (token string, err error) {
	if err = f.Valid(); err != nil {
		return
	}
	var (
		data = &models.UserBase{}
		s    = new(session.Session)
		rds  *redis.Client
	)
	if !setting.DevEnv {
		verify := captcha.VerifyString(f.CaptchaId, f.Digits)
		if !verify {
			err = utils.ErrCaptchaVerifyFail
			return
		}
	}
	if err = data.GetByUsername(f.Username); err != nil {
		return
	}
	if data.Status == models.UserStatusLock {
		err = utils.ErrUserAccountFroze
		return
	}
	// 验证密码是否正确
	if err = bcrypt.CompareHashAndPassword([]byte(data.LoginPwd), []byte(f.Password)); err != nil {
		return
	}
	s.Content = map[string]interface{}{}
	s.Content["password"] = data.LoginPwd
	if f.DeviceId != nil {
		s.Content["deviceId"] = *f.DeviceId
	}
	s.UID = data.Uid
	s.Content["ip"] = f.IP

	n := time.Now()
	s.Content["createTime"] = n.Unix()
	s.ExpireTime = n.Add(session.ExpMaxSecure).Unix()
	s.Content["level"] = session.SessionLevelSecure
	s.Content["id"] = int64(data.ID)
	s.Content["username"] = data.Username
	if !isAdmin {
		if token, err = models.UserKey.SessionEncode(s); err != nil {
			return
		}
	} else {
		if token, err = models.AdminKey.SessionEncode(s); err != nil {
			return
		}
	}
	if rds, err = models.GetRDB(); err != nil {
		return
	}
	rds.Set(data.Uid, token, session.ExpMaxNormal)
	if err = data.UpdateOne(s.UID, "status", models.UserStatusLogin); err != nil {
		return
	}
	//var (
	//	dataOnline *UserOnline
	//)
	//if dataOnline, err = PubUserOnlineGet(data.RoleRefer); err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		dataOnline = new(UserOnline)
	//		dataOnline.RoleRefer = s.UID
	//		if f.DeviceId != nil {
	//			dataOnline.Device = *f.DeviceId
	//		}
	//		dataOnline.IsOnline = true
	//		dataOnline.LoginIp = f.IP
	//		dataOnline.LoginTime = time.Now()
	//		if err = Database.Table(dataOnline.APIGroupTableName()).Create(dataOnline).Error; err != nil {
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
	//	if err = red.Set("cache_token"+data.RoleRefer, token, session.SessionExpMaxSecure).Err(); err != nil {
	//		return
	//	}
	//	red.Close()
	//}

	return
}

// UserLogout 用户登出
func UserLogout(uid string) (err error) {
	var (
		data = new(models.UserBase)
	)
	if err = data.GetByUId(uid); err != nil {
		return err
	}
	//if data.Status != models.UserStatusLogin {
	//	err = utils.ErrActionNotAllow
	//	return
	//}
	return data.UpdateOne(uid, "status", models.UserStatusLogout)
}
