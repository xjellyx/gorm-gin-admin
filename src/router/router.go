package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olongfen/user_base/src/middleware"
	"github.com/olongfen/user_base/src/pkg/setting"
	v1 "github.com/olongfen/user_base/src/router/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// InitRouter 初始化路由模块
func InitRouter() (ret *gin.Engine) {
	// 初始化路由
	var engine = gin.Default()

	if setting.Setting.IsProduct {

		gin.SetMode(gin.ReleaseMode)
		engine.Use(gin.Logger())
	}

	// 添加中间件
	engine.Use(gin.Recovery())
	engine.Use(middleware.CORS())
	engine.Use(middleware.GinAPILog())
	// 没有路由请求
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": fmt.Sprintf("%v ", http.StatusNotFound) + http.StatusText(http.StatusNotFound),
		})
	})
	engine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// TODO 路由
	{
		api := engine.Group("api/v1")
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
	return engine
}

func userRouterAPI(r *gin.RouterGroup) {
	apiUser := r.Group("user")
	apiUser.POST("login", v1.UserLogin)
	apiUser.Use(middleware.CheckUserAuth(false))
	apiUser.POST("logout", v1.UserLogout)
	apiUser.POST("verified", v1.Verified)
	apiUser.POST("setPayPwd", v1.SetPayPwd)
	apiUser.PUT("modifyProfile", v1.ModifyProfile)
	apiUser.PUT("modifyLoginPwd", v1.ModifyLoginPwd)
	apiUser.PUT("modifyPayPwd", v1.ModifyPayPwd)
	apiUser.PUT("modifyHeadIcon", v1.ModifyHeadIcon)
	apiUser.GET("getHeadIcon", v1.GetHeadIcon)
	apiUser.GET("profile", v1.GetUserProfile)
}

func adminRouterAPI(r *gin.RouterGroup) {
	apiAdmin := r.Group("admin")
	apiAdmin.POST("login", v1.AdminLogin)
	apiAdmin.Use(middleware.CheckUserAuth(true))
	apiAdmin.POST("logout", v1.AdminLogout)
	apiAdmin.Use(middleware.CasbinHandler())
	// apiAdmin.POST("/editUser", v1.EditUser)
	apiAdmin.GET("listUser", v1.ListUser)
	apiAdmin.PUT("editUser", v1.EditUser)
	apiAdmin.POST("addRoleApiPerm", v1.AddRoleAPIPerm)
	apiAdmin.DELETE("removeRoleApiPerm", v1.RemoveRolePermAPI)
	apiAdmin.GET("getAllApiGroup", v1.GetAllAPIGroup)
	apiAdmin.POST("addApiGroup", v1.AddApiGroup)
	apiAdmin.DELETE("removeApiGroup", v1.RemoveApiGroup)
	apiAdmin.PUT("editApiGroup", v1.EditApiGroup)
	apiAdmin.GET("getRoleApiList", v1.GetRoleApiList)
	apiAdmin.POST("addMenu", v1.AddMenu)
	apiAdmin.GET("getMenu", v1.GetMenu)
	apiAdmin.GET("getMenuList", v1.GetMenuList)
	apiAdmin.DELETE("delMenu",v1.DelMenu)
	apiAdmin.PUT("editMenu",v1.EditMenu)
	apiAdmin.GET("userTotal",v1.UserTotal)
	// apiAdmin.GET("/profile", auth.RequiresPermissions([]string{"profile:admin"}), api.GetUserProfile)

}
