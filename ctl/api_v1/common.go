package api_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib"
	"github.com/olongfen/contrib/log"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/pkg/setting"
)

var (
	logCtl = log.NewLogFile("./log/log_ctrl", setting.ProjectSetting.IsProduct)
)

// getSession 获取会话信息
func getSession(c *gin.Context) (ret *session.Session, err error) {
	var (
		ok bool
		s  interface{}
	)
	if s, ok = c.Get("sessionTag"); !ok {
		err = contrib.ErrSessionUndefined
		return
	}

	ret = s.(*session.Session)
	return
}
