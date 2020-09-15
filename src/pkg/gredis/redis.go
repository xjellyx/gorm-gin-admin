package gredis

import (
	"github.com/go-redis/redis"
	"github.com/olongfen/gorm-gin-admin/src/setting"
	"github.com/sirupsen/logrus"
)

var redisClient *redis.Client

// init Initialize the Redis instance
func init() {
	var (
		err error
	)
	if redisClient, err = GetRedisClient(); err != nil {
		logrus.Fatal(err)
	}

}

// GetRedisClient 获取redis连接
func GetRedisClient() (ret *redis.Client, err error) {

	if redisClient != nil && redisClient.Ping().Err() == nil {
		return redisClient, nil
	}

	// 重新连接
	redisClient = redis.NewClient(&redis.Options{
		Addr:     setting.Settings.RDB.Host + ":" + setting.Settings.RDB.Port,
		Password: setting.Settings.RDB.Password,
	})
	// 报错直接恐慌
	if err = redisClient.Ping().Err(); err != nil {
		return
	}

	ret = redisClient
	return
}
