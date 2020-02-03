package main

import (
	"github.com/jinzhu/gorm"
	"github.com/olefen/contrib/log"
	"github.com/olefen/userDetail/conf"
	"github.com/olefen/userDetail/model"
	ctrl "github.com/olefen/userDetail/server/ctrl_v1"
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
	//var (
	//	kaep = keepalive.EnforcementPolicy{
	//		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	//		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	//	}
	//	kasp = keepalive.ServerParameters{
	//		MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	//		MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	//		MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	//		Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	//		Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
	//	}
	//)
	// wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	if lis, err = net.Listen("tcp", conf.RpcHost+":"+conf.RpcPort); err != nil {
	//		panic(err)
	//	}
	//	sv, _err := userRpc.NewRpc("")
	//	if _err != nil {
	//		panic(_err)
	//	}
	//	defer sv.Close()
	//	s := grpc.NewServer()
	//	pb.RegisterUserBaseServer(s, sv)
	//	model.LogModel.Infof(`hall rpc serve will be run in %s`, conf.RpcHost+":"+conf.RpcPort)
	//	if err = s.Serve(lis); err != nil {
	//		panic(err)
	//	}
	//
	//}()

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
