package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/src/models"
	"github.com/olongfen/user_base/src/pkg/app"
	"github.com/olongfen/user_base/src/pkg/setting"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		e, err := casbin.NewEnforcer(setting.Setting.RBACModelDir, models.Adapter)
		if err != nil {
			app.NewGinResponse(c).Fail(500, err.Error()).SetStatus(500).Response()
			return
		}
		e.LoadPolicy()
		obj := c.Request.URL.RequestURI()
		act := c.Request.Method
		_d, _ := c.Get("sessionTag")
		s := _d.(*session.Session)
		sub := s.UID
		if ok, err := e.Enforce(sub, obj, act); err != nil {
			app.NewGinResponse(c).Fail(403, "casbnin check failed").SetStatus(403).Response()
			c.Abort()
			return
		} else if !ok && setting.Setting.IsProduct {
			app.NewGinResponse(c).Fail(403, "illegal permission").SetStatus(403).Response()
			c.Abort()
			return
		}
		c.Next()
	}
}
