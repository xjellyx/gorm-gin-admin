package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olongfen/userDetail/ctl/routers/api_v1"
	"github.com/olongfen/userDetail/middleware/cors"
	"github.com/olongfen/userDetail/pkg/setting"
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
	// TODO 路由
	{
		v1 := engine.Group("/v1")
		v1.Use(cors.Common())

		// 测试连接
		v1.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"ping": "pong >>>>>>> update"})
		})

		// User
		{
			//v1_user := v1.Group("/user")
			//v1_user.POST("/register", Ctrl.Register)
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

		}

	}
	return engine
}
