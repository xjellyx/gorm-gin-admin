package models

import (
	"fmt"
	"github.com/olongfen/contrib/log"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/pkg/setting"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	adminKey *session.Key
	userKey  *session.Key
	db       *gorm.DB
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
	if db, err = gorm.Open(postgres.Open(dataSourceName), nil); err != nil {
		logrus.Fatal(err)
	}
	// 初始化密钥对
	if err = userKey.SetRSA(setting.ProjectSetting.AdminKeyDir, setting.ProjectSetting.AdminPubDir); err != nil {
		logrus.Fatal(err)
	}
	if err = adminKey.SetRSA(setting.ProjectSetting.UserKeyDir, setting.ProjectSetting.UserPubDir); err != nil {
		logrus.Fatal(err)
	}
	userKey.SetHookSessionCheck(SessionCheck)
	adminKey.SetHookSessionCheck(SessionCheck)
	err = db.AutoMigrate(&UserBase{})
	if err != nil {
		panic(err)
	}
	// debug
	db.Create(&UserBase{
		Uid:      uuid.NewV4().String(),
		Username: "dasdasd",
		Phone:    "4543534534",
	})
	d := new(UserBase)
	d.ID = 5
	fmt.Println(d.UpdateUser())
}

func init() {
	userKey = session.NewKey("RS256")
	adminKey = session.NewKey("RS256")
}
