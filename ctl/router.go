package ctl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olongfen/user_base/ctl/api_v1"
	"github.com/olongfen/user_base/middleware/cors"
	"github.com/olongfen/user_base/middleware/mdw_sessions"
	"github.com/olongfen/user_base/pkg/setting"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// InitRouter 初始化路由模块
func InitRouter() (ret *gin.Engine) {
	// 初始化路由
	var engine = gin.Default()

	if setting.ProjectSetting.IsProduct {

		gin.SetMode(gin.ReleaseMode)
		engine.Use(gin.Logger())
	}

	// 添加中间件
	engine.Use(gin.Recovery())
	engine.Use(cors.CORS())

	// 没有路由请求
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": fmt.Sprintf("%v ", http.StatusNotFound) + http.StatusText(http.StatusNotFound),
		})
	})
	// 获取验证码
	engine.GET("/srv_captcha", api_v1.Captcha)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// TODO 路由
	{
		v1 := engine.Group("/v1")
		v1.Use(cors.Common())

		// 测试连接
		v1.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"ping": "pong >>>>>>> update"})
		})
		// User
		v1.POST("/register", api_v1.UserRegister)
		{
			v1_user := v1.Group("/user")
			v1_user.POST("/login", api_v1.Login)
			v1_user.Use(mdw_sessions.CheckUserAuth(false))
			v1_user.POST("/updateUser", api_v1.UserUpdate)
			v1_user.GET("/getUserProfile", api_v1.GetUserProfile)
			//v1_user.POST("/login", Ctrl.Login)
			//v1_user.GET("/logout", Ctrl.Logout)
			//v1_user.GET("/getUserDetailSelf", Ctrl.GetUserDetailSelf)
			//v1_user.GET("/getUserBankCardList", Ctrl.GetUserBankCardList)
			//v1_user.GET("/getUserIDCard", Ctrl.GetUserIDCard)
			//v1_user.GET("/getUserAddressList", Ctrl.GetUserAddressList)
			//
			//v1_user.POST("/updateUserDetail", Ctrl.UpdateUserDetail)
			//v1_user.POST("/addUserIDCard", Ctrl.AddUserIDCard)
			//v1_user.POST("/addUserBankCard", Ctrl.AddUserBankCard)
			//v1_user.POST("/addUserAddress", Ctrl.AddUserAddress)
			//v1_user.POST("/updateUserAddress", Ctrl.UpdateUserAddress)
			//
			//v1_user.DELETE("/deleteUserBankCard/:number", Ctrl.DeleteUserBankCard)
			//v1_user.DELETE("/deleteUserAddress/:id", Ctrl.DeleteUserAddress)
		}

		// Admin
		{
			v1_admin := v1.Group("/admin")
			v1_admin.Use(mdw_sessions.CheckUserAuth(true))
		}

	}
	return engine
}
