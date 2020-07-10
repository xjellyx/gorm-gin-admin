package models

import (
	"fmt"

	"github.com/olongfen/contrib/log"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/pkg/adapter"
	"github.com/olongfen/user_base/pkg/setting"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	AdminKey *session.Key
	UserKey  *session.Key
	Adapter  *adapter.Adapter
	DB       *gorm.DB
	logModel *logrus.Logger
	//Captcha
)

// InitModel 初始化模型
func InitModel() {
	var (
		err error
	)
	logModel = log.NewLogFile(setting.ProjectSetting.LogDir+"/"+"model", setting.ProjectSetting.IsProduct)
	dataSourceName := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", setting.ProjectSetting.Db.Driver, setting.ProjectSetting.Db.Username,
		setting.ProjectSetting.Db.Password, setting.ProjectSetting.Db.Host, setting.ProjectSetting.Db.Port, setting.ProjectSetting.Db.DatabaseName)
	if DB, err = gorm.Open(postgres.Open(dataSourceName), nil); err != nil {
		logrus.Fatal(err)
	}
	DB = DB.Debug()
	// 初始化密钥对
	if err = UserKey.SetRSA(setting.ProjectSetting.AdminKeyDir, setting.ProjectSetting.AdminPubDir); err != nil {
		logrus.Fatal(err)
	}
	if err = AdminKey.SetRSA(setting.ProjectSetting.UserKeyDir, setting.ProjectSetting.UserPubDir); err != nil {
		logrus.Fatal(err)
	}
	err = DB.AutoMigrate(&UserBase{}, &UserCard{}, &APIGroup{})
	if err != nil {
		panic(err)
	}
	if Adapter, err = adapter.NewAdapterByDB(DB); err != nil {
		panic(err)
	}

}

func init() {
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
