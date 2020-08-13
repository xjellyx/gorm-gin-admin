package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/pkg/app"
	"github.com/olongfen/user_base/pkg/codes"
	"github.com/olongfen/user_base/service"
	"github.com/olongfen/user_base/utils"
)

// Verified 用户实名认证
// @tags 用户
// @Summary 用户实名认证
// @Produce json
// @Param form body utils.FormIDCard true "实名认证Form"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/user/verified [post]
func Verified(c *gin.Context) {
	var (
		s    *session.Session
		form = new(utils.FormIDCard)
		err  error
		code = codes.CodeProcessingFailed
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		}
	}()
	if s, err = GetSession(c); err != nil {
		return
	}
	if err = c.ShouldBind(form); err != nil {
		code = codes.CodeParamInvalid
		return
	}
	if _d, _err := service.AddIDCard(s.UID, form); _err != nil {
		err = _err
		return
	} else {
		app.NewGinResponse(c).Success(_d).Response()
	}
}
