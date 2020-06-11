package models

import (
	"fmt"
	"github.com/go-redis/redis"

	log "github.com/olongfen/contrib/log"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/userDetail/pkg/setting"
)

var (
	adminKey *session.Key
	userKey  *session.Key
	rdb      *redis.Client
	logModel = log.NewLogFile(setting.ProjectSetting.LogDir+"/"+"model",setting.ProjectSetting.IsProduct)
  	//Captcha
)

// InitModel 初始化模型
func InitModel()  {
	var (
		err error
	)
	dataSourceName := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", setting.ProjectSetting.Db.Driver, setting.ProjectSetting.Db.Username,
		setting.ProjectSetting.Db.Password, setting.ProjectSetting.Db.Host, setting.ProjectSetting.Db.Port, setting.ProjectSetting.Db.DatabaseName)

   _ =dataSourceName
	// 初始化密钥对
	if err = userKey.SetRSA(setting.ProjectSetting.AdminKeyDir, setting.ProjectSetting.AdminPubDir); err != nil {
		return
	}
	if err = adminKey.SetRSA(setting.ProjectSetting.UserKeyDir, setting.ProjectSetting.UserPubDir); err != nil {
		return
	}
	//userKey.SetHookSessionCheck( SessionCheck)
	//adminKey.SetHookSessionCheck( SessionCheck)
	// 初始化redis连接
	if rdb, err = GetRedisClient(); err != nil {
		return
	}
	return
}

func init() {
	userKey = session.NewKey("RS256")
	adminKey = session.NewKey("RS256")
}

// GetRedisClient 获取redis连接
func GetRedisClient() (ret *redis.Client, err error) {

	if rdb != nil && rdb.Ping().Err() == nil {
		return rdb, nil
	}

	// 重新连接
	rdb = redis.NewClient(&redis.Options{
		Addr:     setting.ProjectSetting.RDB.Host + ":" + setting.ProjectSetting.RDB.Port,
		Password: setting.ProjectSetting.RDB.Password,
	})
	// 报错直接恐慌
	if err = rdb.Ping().Err(); err != nil {
		return
	}

	ret = rdb
	return
}
