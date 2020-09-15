package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "github.com/olongfen/gorm-gin-admin/src/controller/api/v1"
	"github.com/olongfen/gorm-gin-admin/src/middleware"
	"github.com/olongfen/gorm-gin-admin/src/setting"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// 初始化路由
var Engine = gin.Default()

// init 初始化路由模块
func init() {
	if !setting.DevEnv {
		gin.SetMode(gin.ReleaseMode)
		Engine.Use(gin.Logger())
	}

	// 添加中间件
	Engine.Use(gin.Recovery())
	Engine.Use(middleware.CORS())
	Engine.Use(middleware.GinAPILog())
	// 没有路由请求
	Engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": fmt.Sprintf("%v ", http.StatusNotFound) + http.StatusText(http.StatusNotFound),
		})
	})
	// TODO 路由
	{
		api := Engine.Group("api/v1")
		api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.Use(middleware.Common())

		// 测试连接
		api.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"ping": "pong >>>>>>> update"})
		})
		// 获取验证码
		api.GET("captcha", v1.Captcha)
		// api.GET("captcha/verify/", v1.VerifyCaptcha)
		api.POST("register", v1.UserRegister)
		userRouterAPI(api)
		adminRouterAPI(api)

	}
}

func userRouterAPI(r *gin.RouterGroup) {
	apiUser := r.Group("user")
	apiUser.POST("login", v1.UserLogin)
	apiUser.Use(middleware.CheckUserAuth(false))
	userBaseRouterAPI(apiUser)
}

func adminRouterAPI(r *gin.RouterGroup) {
	apiAdmin := r.Group("admin")
	apiAdmin.POST("login", v1.AdminLogin)
	apiAdmin.Use(middleware.CheckUserAuth(true))
	apiAdmin.POST("logout", v1.AdminLogout)
	apiAdmin.Use(middleware.CasbinHandler())
	adminActionUserRouterAPI(apiAdmin)
	apiRolePermRouterAPI(apiAdmin)
	apiRouterAPI(apiAdmin)
	menuRouterAPI(apiAdmin)
	roleRouterAPI(apiAdmin)
	// apiAdmin.GET("/profile", auth.RequiresPermissions([]string{"profile:admin"}), api.GetUserProfile)

}
