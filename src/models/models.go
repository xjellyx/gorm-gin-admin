package models

import (
	"fmt"
	"time"

	"github.com/olongfen/contrib/log"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/gorm-gin-admin/src/pkg/adapter"
	"github.com/olongfen/gorm-gin-admin/src/setting"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-"gorm:"index"`
}

var (
	AdminKey *session.Key
	UserKey  *session.Key
	Adapter  *adapter.Adapter
	DB       *gorm.DB
	logModel *log.Logger
	//Captcha
)

// init 初始化模型
func init() {
	var (
		err error
	)
	initKey()
	logModel = log.NewLogFile(log.ParamLog{Path: setting.Settings.FilePath.LogDir + "/" + "model", Stdout: setting.DevEnv, P: setting.Settings.FilePath.LogPatent})
	dataSourceName := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", setting.Settings.DB.Driver, setting.Settings.DB.Username,
		setting.Settings.DB.Password, setting.Settings.DB.Host, setting.Settings.DB.Port, setting.Settings.DB.DatabaseName)
	if DB, err = gorm.Open(postgres.Open(dataSourceName), nil); err != nil {
		logrus.Fatal(err)
	}
	if setting.DevEnv {
		DB = DB.Debug()
	}
	err = DB.AutoMigrate(&UserBase{}, &UserCard{}, &APIGroup{}, &Menu{}, &Role{}, &BehaviorRecord{})
	if err != nil {
		panic(err)
	}
	if Adapter, err = adapter.NewAdapterByDB(DB); err != nil {
		panic(err)
	}
	// 初始化密钥对
	if err = UserKey.SetRSA(setting.Settings.FilePath.AdminKeyDir, setting.Settings.FilePath.AdminPubDir); err != nil {
		logrus.Fatal(err)
	}
	if err = AdminKey.SetRSA(setting.Settings.FilePath.UserKeyDir, setting.Settings.FilePath.UserPubDir); err != nil {
		logrus.Fatal(err)
	}

	logrus.Infoln("database init success !")
}

func initKey() {
	UserKey = session.NewKey("RS256")
	AdminKey = session.NewKey("RS256")
	UserKey.SetHookSessionCheck(func(sess *session.Session) error {
		return nil
	})
	UserKey.SetHookTokenCheck(func(token interface{}) error {
		return nil
	})
	AdminKey.SetHookSessionCheck(func(sess *session.Session) error {
		return nil
	})
	AdminKey.SetHookTokenCheck(func(token interface{}) error {
		return nil
	})
}
