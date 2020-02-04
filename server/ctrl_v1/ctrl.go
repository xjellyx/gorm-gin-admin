package ctrl

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olefen/contrib/log"
	"github.com/olefen/userDetail/conf"
)

var (
	Ctrl *ControlServe
)

// ControlServe 服务控制器
type ControlServe struct {
	Log    *log.Logger
	Engine *gin.Engine
}

// InitCtrl 初始化控制器，开启服务
func InitCtrl() (err error) {
	Ctrl = &ControlServe{}
	// 初始化路由
	Ctrl.Engine = gin.Default()
	if conf.ProjectSetting.Mode == conf.ModeProduction {
		Ctrl.Log = log.NewLogFile("./log/log_ctrl")
		gin.SetMode(gin.ReleaseMode)
		Ctrl.Engine.Use(gin.Logger())
	} else {
		Ctrl.Log, _ = log.NewLog(nil)
	}

	// 添加中间件
	Ctrl.Engine.Use(gin.Recovery())
	Ctrl.Engine.Use(Ctrl.CORS())

	// 没有路由请求
	Ctrl.Engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": fmt.Sprintf("%v ", http.StatusNotFound) + http.StatusText(http.StatusNotFound),
		})
	})

	// TODO 路由
	{
		v1 := Ctrl.Engine.Group("/v1")
		v1.Use(Ctrl.Common)

		// 测试连接
		v1.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"ping": "pong"})
		})
		v1.POST("/register", Ctrl.Register)
		v1.POST("/login", Ctrl.Login)

		// User
		{
			v1_user := v1.Group("/user")
			v1_user.GET("/getUserDetailSelf", Ctrl.GetUserDetailSelf)
			v1_user.GET("/getUserBankCardList", Ctrl.GetUserBankCardList)
			v1_user.GET("/getUserIDCard", Ctrl.GetUserIDCard)
			v1_user.GET("/getUserAddressList", Ctrl.GetUserAddressList)

			v1_user.POST("/updateUserDetail", Ctrl.UpdateUserDetail)
			v1_user.POST("/addUserIDCard", Ctrl.AddUserIDCard)
			v1_user.POST("/addUserBankCard", Ctrl.AddUserBankCard)
			v1_user.POST("/addUserAddress", Ctrl.AddUserAddress)
			v1_user.POST("/updateUserAddress", Ctrl.UpdateUserAddress)

			v1_user.DELETE("/deleteUserBankCard/:number", Ctrl.DeleteUserBankCard)
			v1_user.DELETE("/deleteUserAddress/:id", Ctrl.DeleteUserAddress)
		}

		// Admin
		{

		}

	}

	// 开启服务
	s := &http.Server{
		Addr:           conf.ProjectSetting.ServerAddr + ":" + conf.ProjectSetting.ServerPort,
		Handler:        Ctrl.Engine,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20, // 10M
	}
	Ctrl.Log.Infof("server listen on: %s\n ", s.Addr)

	if conf.ProjectSetting.IsTLS { // 开启tls
		TLSConfig := &tls.Config{
			MinVersion:               tls.VersionTLS11,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		}

		TLSProto := make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)

		s.TLSConfig = TLSConfig
		s.TLSNextProto = TLSProto

		if err = s.ListenAndServeTLS(conf.ProjectSetting.TLS.Cert, conf.ProjectSetting.TLS.Key); err != nil {
			Ctrl.Log.Errorln(err)
			return
		}
	} else {
		if err = s.ListenAndServe(); err != nil {
			Ctrl.Log.Errorln(err)
			return
		}
	}

	return
}
