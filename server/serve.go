package main

import (
	"github.com/jinzhu/gorm"
	"github.com/srlemon/contrib/log"
	"github.com/srlemon/userDetail/conf"
	"github.com/srlemon/userDetail/model"
	ctrl "github.com/srlemon/userDetail/server/ctrl_v1"
	"io/ioutil"
	"sync"
)

func main() {
	var (
		db                                   *gorm.DB
		err                                  error
		userKey, userPub, adminKey, adminPub []byte
		wg                                   = &sync.WaitGroup{}
	)

	// 先初始化配置文件
	if err = conf.InitConfig(); err != nil {
		panic(err)
	}

	// 初始化日志
	if err = initLog(); err != nil {
		panic(err)
	}

	// 获取密钥配置
	if userKey, err = ioutil.ReadFile(conf.ProjectSetting.UserKeyDir); err != nil {
		panic(err)
	}
	if userPub, err = ioutil.ReadFile(conf.ProjectSetting.UserPubDir); err != nil {
		panic(err)
	}
	if adminKey, err = ioutil.ReadFile(conf.ProjectSetting.AdminKeyDir); err != nil {
		panic(err)
	}
	if adminPub, err = ioutil.ReadFile(conf.ProjectSetting.AdminPubDir); err != nil {
		panic(err)
	}
	// 初始化模型
	if db, err = model.InitModel(model.InitModelParam{
		Db:       *conf.ProjectSetting.Db,
		Sync:     conf.ProjectSetting.Sync,
		Mode:     conf.ProjectSetting.Mode,
		UserPub:  userPub,
		UserKey:  userKey,
		AdminPub: adminPub,
		AdminKey: adminKey,
	}); err != nil {
		panic(err)
	}
	defer db.Close()

	// 开启接口服务
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = ctrl.InitCtrl(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

// initLog
func initLog() (err error) {
	if conf.ProjectSetting.Mode == "production" {
		model.LogModel = log.NewLogFile("./log/log_model")
	} else {
		if model.LogModel, err = log.NewLog(nil); err != nil {
			return
		}
	}
	return
}
