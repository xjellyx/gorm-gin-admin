package models

import (
	"fmt"
	"time"

	"github.com/olongfen/contrib/log"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/src/pkg/adapter"
	"github.com/olongfen/user_base/src/pkg/setting"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"gorm:"index"`
}

var (
	AdminKey *session.Key
	UserKey  *session.Key
	Adapter  *adapter.Adapter
	DB       *gorm.DB
	logModel *log.Logger
	//Captcha
)

// InitModel 初始化模型
func InitModel() {
	var (
		err error
	)
	logModel = log.NewLogFile(log.ParamLog{Path: setting.Setting.LogDir + "/" + "model", Stdout: setting.Setting.IsProduct, P: setting.Setting.LogPatent})
	dataSourceName := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", setting.Setting.Db.Driver, setting.Setting.Db.Username,
		setting.Setting.Db.Password, setting.Setting.Db.Host, setting.Setting.Db.Port, setting.Setting.Db.DatabaseName)
	if DB, err = gorm.Open(postgres.Open(dataSourceName), nil); err != nil {
		logrus.Fatal(err)
	}
	DB = DB.Debug()
	// 初始化密钥对
	if err = UserKey.SetRSA(setting.Setting.AdminKeyDir, setting.Setting.AdminPubDir); err != nil {
		logrus.Fatal(err)
	}
	if err = AdminKey.SetRSA(setting.Setting.UserKeyDir, setting.Setting.UserPubDir); err != nil {
		logrus.Fatal(err)
	}
	err = DB.AutoMigrate(&UserBase{}, &UserCard{}, &APIGroup{}, &Menu{})
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
