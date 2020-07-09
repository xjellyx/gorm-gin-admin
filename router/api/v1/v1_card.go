package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/pkg/app"
	"github.com/olongfen/user_base/service"
	"github.com/olongfen/user_base/utils"
	"net/http"
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
		s        *session.Session
		form     = new(utils.FormIDCard)
		err      error
		httpCode = http.StatusInternalServerError
	)
	defer func() {
		if err != nil {
			app.NewGin(c).Response(httpCode, err.Error())
		}
	}()
	if s, err = GetSession(c); err != nil {
		return
	}
	if err = c.ShouldBind(form); err != nil {
		httpCode = 404
		return
	}
	if _d, _err := service.AddIDCard(s.UID, form); _err != nil {
		err = _err
		return
	} else {
		app.NewGin(c).Response(200, _d)
	}
}
