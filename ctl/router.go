package ctl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log2 "github.com/olongfen/contrib/log"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/ctl/api_v1"
	"github.com/olongfen/user_base/middleware/cors"
	"github.com/olongfen/user_base/middleware/mdw_sessions"
	"github.com/olongfen/user_base/middleware/rbac"
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/pkg/setting"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
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
	engine.Use(cors.CORS())
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
		Output: log2.NewLogFile("./log/ctl", !setting.ProjectSetting.IsProduct).Out,
	}))
	// 没有路由请求
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": fmt.Sprintf("%v ", http.StatusNotFound) + http.StatusText(http.StatusNotFound),
		})
	})
	// 获取验证码
	engine.GET("/captcha", api_v1.Captcha)
	engine.GET("/captcha/verify/", api_v1.VerifyCaptcha)
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
			v1_user.POST("/login", api_v1.UserLogin)
			v1_user.Use(mdw_sessions.CheckUserAuth(false))

			v1_user.POST("/logout", api_v1.UserLogout)
			v1_user.POST("/verified", api_v1.Verified)
			v1_user.POST("/setPayPwd", api_v1.SetPayPwd)
			v1_user.PUT("/modifyProfile", api_v1.ModifyProfile)
			v1_user.PUT("/modifyLoginPwd", api_v1.ModifyLoginPwd)
			v1_user.PUT("/modifyPayPwd", api_v1.ModifyPayPwd)
			v1_user.PUT("/modifyHeadIcon", api_v1.ModifyHeadIcon)
			v1_user.GET("/getHeadIcon", api_v1.GetHeadIcon)
			v1_user.GET("/profile", api_v1.GetUserProfile)

		}

		// Adminmodify
		{
			v1_admin := v1.Group("/admin")
			v1_admin.POST("/login", api_v1.AdminLogin)
			v1_admin.Use(mdw_sessions.CheckUserAuth(true))
			auth, err := rbac.NewCasbinMiddleware(setting.ProjectSetting.RBACModelDir, setting.ProjectSetting.RABCPolicyDir, func(c *gin.Context) string {
				_d, _ := c.Get("sessionTag")
				s := _d.(*session.Session)
				u := new(models.UserBase)
				if err := u.GetUserByUId(s.UID); err != nil {
					return ""
				}
				return u.Username
			})
			if err != nil {
				log.Fatal(err)
			}
			v1_admin.POST("/logout", auth.RequiresPermissions([]string{"logout:admin"}), api_v1.AdminLogout)
			v1_admin.POST("/editUser", auth.RequiresPermissions([]string{"editUser:rootAdmin"}))
			v1_admin.GET("/userList", auth.RequiresPermissions([]string{"userList:admin"}), api_v1.UserList)
			v1_admin.POST("/editUser/:uid", auth.RequiresPermissions([]string{"editUser:admin"}), api_v1.EditUser)
			// v1_admin.GET("/profile", auth.RequiresPermissions([]string{"profile:admin"}), api_v1.GetUserProfile)

		}

	}
	return engine
}
