package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/pkg/app"
	"github.com/olongfen/user_base/pkg/setting"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		e, err := casbin.NewEnforcer(setting.ProjectSetting.RBACModelDir, models.Adapter)
		if err != nil {
			app.NewGinResponse(c).Fail(500, err.Error()).Response()
			return
		}
		e.LoadPolicy()
		obj := c.Request.URL.RequestURI()
		act := c.Request.Method
		_d, _ := c.Get("sessionTag")
		s := _d.(*session.Session)
		sub := s.UID
		if ok, err := e.Enforce(sub, obj, act); err != nil {
			app.NewGinResponse(c).Fail(403, "casbnin check failed").Response()
			c.Abort()
			return
		} else if !ok && setting.ProjectSetting.IsProduct {
			app.NewGinResponse(c).Fail(403, "illegal permission").Response()
			c.Abort()
			return
		}
		c.Next()
	}
}
