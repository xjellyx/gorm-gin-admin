package api_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib"
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/pkg/app"
	"github.com/olongfen/user_base/service/srv_user"
	"net/http"
)

// UserRegister 用户注册
// @Summary 用户注册
// @Produce json
// @Param phone body string true "Phone"
// @Param password body string true "Password"
// @Param code body string  false "Code"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/register [post]
func UserRegister(c *gin.Context) {
	var (
		form     = new(srv_user.AddUserForm)
		data     *models.UserBase
		httpCode = http.StatusInternalServerError
		err      error
	)

	defer func() {
		if err != nil {
			app.NewGin(c).Response(httpCode, err.Error())
		} else {
			app.NewGin(c).Response(200, data)
		}
	}()

	if err = c.Bind(form); err != nil {
		httpCode = http.StatusBadRequest
		err = contrib.ErrParamInvalid
		return
	}
	if data, err = srv_user.AddUser(form); err != nil {
		return
	}
}
