package models

import (
	"github.com/go-redis/redis"
	"github.com/olongfen/gorm-gin-admin/src/setting"
	"github.com/sirupsen/logrus"
)

var RDB *redis.Client

// init Initialize the Redis instance
func init() {
	var (
		err error
	)
	if RDB, err = GetRDB(); err != nil {
		logrus.Fatal(err)
	}
	logrus.Infoln("redis init success !")
}

// GetRDB 获取redis连接
func GetRDB() (ret *redis.Client, err error) {

	if RDB != nil && RDB.Ping().Err() == nil {
		return RDB, nil
	}
	// 重新连接
	RDB = redis.NewClient(&redis.Options{
		Addr:     setting.Settings.RDB.Host + ":" + setting.Settings.RDB.Port,
		Password: setting.Settings.RDB.Password,
	})
	if err = RDB.Ping().Err(); err != nil {
		return
	}

	ret = RDB
	return
}
