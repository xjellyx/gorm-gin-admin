package main

/*import (
	"gorm.io/gorm"
	"github.com/olongfen/contrib"
	"github.com/olongfen/userDetail/pkg/setting"
	ctrl "github.com/olongfen/userDetail/server/ctrl_v1"
	"google.golang.org/grpc"
	"os"

	"github.com/olongfen/contrib/log"
	pb "github.com/olongfen/models.grpc"
	"github.com/olongfen/userDetail/models"
	userRpc "github.com/olongfen/userDetail/rpc"
	"io/ioutil"
	"net"
	"sync"
)

func main() {
	var (
		db                                   *gorm.DB
		err                                  error
		userKey, userPub, adminKey, adminPub []byte
		wg                                   = &sync.WaitGroup{}
	)
	dir, _ := os.Getwd()
	dir = contrib.PubGetEnvString("CONF_DIR", dir)
	// 先初始化配置文件
	if err = setting.InitConfig(dir); err != nil {
		panic(err)
	}

	// 初始化日志
	if err = initLog(); err != nil {
		panic(err)
	}

	// 获取密钥配置
	if userKey, err = ioutil.ReadFile(setting.ProjectSetting.UserKeyDir); err != nil {
		panic(err)
	}
	if userPub, err = ioutil.ReadFile(setting.ProjectSetting.UserPubDir); err != nil {
		panic(err)
	}
	if adminKey, err = ioutil.ReadFile(setting.ProjectSetting.AdminKeyDir); err != nil {
		panic(err)
	}
	if adminPub, err = ioutil.ReadFile(setting.ProjectSetting.AdminPubDir); err != nil {
		panic(err)
	}
	// 初始化模型
	if db, err = models.InitModel(models.InitModelParam{
		Db:       *setting.ProjectSetting.Db,
		Sync:     setting.ProjectSetting.Sync,
		Mode:     setting.ProjectSetting.IsProduct,
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
	wg.Add(1)
	go func() {
		var (
			lis net.Listener
		)
		defer wg.Done()
		if lis, err = net.Listen("tcp", setting.ProjectSetting.RpcHost+":"+setting.ProjectSetting.RpcPort); err != nil {
			panic(err)
		}
		sv, _err := userRpc.NewRpc(setting.ProjectSetting.RpcHost + ":" + setting.ProjectSetting.RpcPort)
		if _err != nil {
			panic(_err)
		}
		defer sv.Close()
		s := grpc.NewServer()
		pb.RegisterUserBaseServer(s, sv)
		models.logModel.Infof(`hall rpc serve will be run in %s`, setting.ProjectSetting.RpcHost+":"+setting.ProjectSetting.RpcPort)
		if err = s.Serve(lis); err != nil {
			panic(err)
		}

	}()

	wg.Wait()
}

// initLog
func initLog() (err error) {
	if setting.ProjectSetting.IsProduct {
		models.logModel = log.NewLogFile("./log/log_model")
	} else {
		if models.logModel, err = log.NewLog(nil); err != nil {
			return
		}
	}
	return
}
*/