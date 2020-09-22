package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/gorm-gin-admin/src/models"
	"github.com/olongfen/gorm-gin-admin/src/pkg/app"
	"github.com/olongfen/gorm-gin-admin/src/setting"
	"github.com/olongfen/gorm-gin-admin/src/utils"
	"strings"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		e, err := casbin.NewEnforcer(setting.Settings.FilePath.RBACModelDir, models.Adapter)
		if err != nil {
			app.NewGinResponse(c).Fail(500, err.Error()).SetStatus(500).Response()
			return
		}
		e.LoadPolicy()
		obj := strings.Split(c.Request.URL.RequestURI(), "?")[0]
		act := strings.ToLower(c.Request.Method)
		_d, _ := c.Get("sessionTag")
		s := _d.(*session.Session)

		d := &models.UserBase{}
		if err = d.GetByUId(s.UID); err != nil {
			err = utils.ErrUserNotExist
			return
		}
		sub := d.Role.Role
		//// 验证是否已经是登录状态
		//if !CheckUserLogin(d){
		//	app.NewGinResponse(c).Fail(403, "does't login").SetStatus(403).Response()
		//	c.Abort()
		//	return
		//}
		if ok, err := e.Enforce(sub, obj, act); err != nil {
			app.NewGinResponse(c).Fail(401, "casbnin check failed").SetStatus(403).Response()
			c.Abort()
			return
		} else if !ok && !setting.DevEnv && d.Role.GetLevelMust() < setting.Settings.Project.MaxRoleLevel {
			app.NewGinResponse(c).Fail(401, "illegal permission").SetStatus(403).Response()
			c.Abort()
			return
		}
		c.Next()
	}
}
