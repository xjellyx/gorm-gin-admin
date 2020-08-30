package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib"
	"github.com/olongfen/contrib/log"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/src/pkg/codes"
	"github.com/olongfen/user_base/src/pkg/setting"
)

var (
	logCtl = log.NewLogFile(log.ParamLog{Path: "./log/log_ctrl", Stdout: setting.Setting.IsProduct, P: setting.Setting.LogPatent})
)

// GetSession 获取会话信息
func GetSession(c *gin.Context) (ret *session.Session,code int, err error) {
	var (
		ok bool
		s  interface{}
	)
	if s, ok = c.Get("sessionTag"); !ok {
		err = contrib.ErrSessionUndefined
		code = codes.CodeTokenInvalid
		return
	}

	ret = s.(*session.Session)
	return
}

func GetSessionAndBindingForm(form interface{}, c *gin.Context) (ret *session.Session, code int, err error) {
	if ret,code, err = GetSession(c); err != nil {
		return
	}
	if err = c.ShouldBind(form); err != nil {
		return nil, codes.CodeParamInvalid, err
	}
	return
}
