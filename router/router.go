package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log2 "github.com/olongfen/contrib/log"
	"github.com/olongfen/user_base/middleware"
	"github.com/olongfen/user_base/pkg/setting"
	v1 "github.com/olongfen/user_base/router/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"time"
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
	engine.Use(middleware.CORS())
	engine.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(param gin.LogFormatterParams) string {
			// 你的自定义格式
			return fmt.Sprintf(`address: %s, time: [%s], method: %s,  message: %s `+"\n"+`path: %s, proto: %s, code: %d, latency: %s, agent: %s`+"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.ErrorMessage,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
			)
		},
		Output: log2.NewLogFile("./log/router", !setting.ProjectSetting.IsProduct).Out,
	}))
	// 没有路由请求
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": fmt.Sprintf("%v ", http.StatusNotFound) + http.StatusText(http.StatusNotFound),
		})
	})
	// 获取验证码
	engine.GET("captcha", v1.Captcha)
	engine.GET("captcha/verify/", v1.VerifyCaptcha)
	engine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// TODO 路由
	{
		api := engine.Group("api")
		api.Use(middleware.Common())

		// 测试连接
		api.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"ping": "pong >>>>>>> update"})
		})
		// User
		api.POST("register", v1.UserRegister)
		{
			api_user := api.Group("user")
			api_user.POST("login", v1.UserLogin)
			api_user.Use(middleware.CheckUserAuth(false))

			api_user.POST("logout", v1.UserLogout)
			api_user.POST("verified", v1.Verified)
			api_user.POST("setPayPwd", v1.SetPayPwd)
			api_user.PUT("modifyProfile", v1.ModifyProfile)
			api_user.PUT("modifyLoginPwd", v1.ModifyLoginPwd)
			api_user.PUT("modifyPayPwd", v1.ModifyPayPwd)
			api_user.PUT("modifyHeadIcon", v1.ModifyHeadIcon)
			api_user.GET("getHeadIcon", v1.GetHeadIcon)
			api_user.GET("profile", v1.GetUserProfile)

		}

		// Adminmodify
		{
			api_admin := api.Group("admin")
			api_admin.POST("login", v1.AdminLogin)
			api_admin.Use(middleware.CheckUserAuth(true)).Use(middleware.CasbinHandler())
			api_admin.POST("logout", v1.AdminLogout)
			// api_admin.POST("/editUser", v1.EditUser)
			api_admin.GET("userList/:uid", v1.UserList)
			api_admin.POST("editUser/:uid", v1.EditUser)
			// api_admin.GET("/profile", auth.RequiresPermissions([]string{"profile:admin"}), api.GetUserProfile)

		}

	}
	return engine
}
