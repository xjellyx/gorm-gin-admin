package model

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/olongfen/contrib/log"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/userDetail/conf"
)

var (
	AdminKey *session.Key
	UserKey  *session.Key
	RDB      *redis.Client
	Database *gorm.DB
	LogModel *log.Logger
  	//Captcha
)

// InitModelParam 初始化模型参数
type InitModelParam struct {
	Db       conf.Database
	Sync     bool
	Mode     bool
	UserKey  []byte
	UserPub  []byte
	AdminKey []byte
	AdminPub []byte
}

// InitModel 初始化模型
func InitModel(d InitModelParam) (ret *gorm.DB, err error) {
	var (
		db *gorm.DB
	)
	dataSourceName := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", d.Db.Driver, d.Db.Username,
		d.Db.Password, d.Db.Host, d.Db.Port, d.Db.DatabaseName)

	if db, err = gorm.Open(d.Db.Driver, dataSourceName); err != nil {
		return
	}
	db.LogMode(d.Mode)
	db.DB().SetMaxIdleConns(d.Db.MaxIdleConn) // 设置最大闲置个数
	db.DB().SetMaxOpenConns(d.Db.MaxOpenConn) // 最大打开的连接数
	db.SingularTable(true)                    // 表生成结尾不带s
	// 生成数据
	if d.Sync {
		LogModel.Infoln("[正在同步数据...]")
		db.AutoMigrate(
			new(UserDetail),
			new(BankCard),
			new(AddressDetail),
			new(IDCard),
			new(UserOnline),
		)
		LogModel.Infoln("[同步数据完成]")
	}
	Database = db

	// 初始化密钥对
	if err = UserKey.SetRSA(d.UserKey, d.UserPub); err != nil {
		return
	}
	if err = AdminKey.SetRSA(d.UserKey, d.UserPub); err != nil {
		return
	}
	UserKey.HookSessionCheck = SessionCheck
	AdminKey.HookSessionCheck = SessionCheck
	// 初始化redis连接
	if RDB, err = GetRedisClient(); err != nil {
		return
	}
	go initCorn()
	//
	ret = db
	return
}

func init() {
	UserKey, _ = session.NewKey(nil)
	AdminKey, _ = session.NewKey(nil)
}

// GetRedisClient 获取redis连接
func GetRedisClient() (ret *redis.Client, err error) {

	if RDB != nil && RDB.Ping().Err() == nil {
		return RDB, nil
	}

	// 重新连接
	RDB = redis.NewClient(&redis.Options{
		Addr:     conf.ProjectSetting.RDB.Host + ":" + conf.ProjectSetting.RDB.Port,
		Password: conf.ProjectSetting.RDB.Password,
	})
	// 报错直接恐慌
	if err = RDB.Ping().Err(); err != nil {
		return
	}

	ret = RDB
	return
}
