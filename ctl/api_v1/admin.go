package api_v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/pkg/app"
	"github.com/olongfen/user_base/service/srv_user"
	"github.com/olongfen/user_base/utils"
	"net/http"
)

// @tags 管理员
// @Summary 管理员登录
// @Produce json
// @Accept json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/admin/login [post]
// AdminLogin
func AdminLogin(c *gin.Context) {
	var (
		form     = &utils.LoginForm{}
		err      error
		httpCode = http.StatusInternalServerError
		token    string
	)
	defer func() {
		if err != nil {
			app.NewGin(c).Response(httpCode, err.Error())
		} else {
			app.NewGin(c).Response(200, map[string]string{"token": token})
		}
	}()

	form.IP = c.ClientIP()
	if err = c.ShouldBind(form); err != nil {
		return
	}
	if token, err = srv_user.UserLogin(form, true); err != nil {
		return
	}
}

// AdminLogout 登出
// @tags 管理员
// @Summary 管理员登出
// @Produce json
// @Accept json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/admin/logout [post]
func AdminLogout(c *gin.Context) {
	var (
		err      error
		httpCode = http.StatusInternalServerError
		s        *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGin(c).Response(httpCode, err.Error())
		} else {
			app.NewGin(c).Response(200, map[string]string{})
		}
	}()
	if s, err = GetSession(c); err != nil {
		return
	}
	if err = srv_user.UserLogout(s.UID); err != nil {
		return
	}

}

// @tags 管理员
// @Title 获取用户列表
// @Summary 获取用户列表
// @Description 获取用户列表
// @Accept json
// @Produce json
// @Param {} body utils.FormUserList true "查询数据"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/admin/userList [get]
func UserList(c *gin.Context) {
	var (
		err      error
		httpCode = http.StatusInternalServerError
		s        *session.Session
		form     = new(utils.FormUserList)
		data     []*models.UserBase
	)
	defer func() {
		if err != nil {
			app.NewGin(c).Response(httpCode, err.Error())
		}
	}()

	if s, err = GetSession(c); err != nil {
		return
	}
	if err = c.ShouldBindQuery(form); err != nil {
		fmt.Println("sssssssssss", err)
		return
	}
	if data, err = srv_user.GetUserList(s.UID, form); err != nil {
		return
	}
	app.NewGin(c).Response(200, data)
}
