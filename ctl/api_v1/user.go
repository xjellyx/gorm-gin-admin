package api_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/pkg/app"
	"github.com/olongfen/user_base/service/srv_user"
	"github.com/olongfen/user_base/utils"
	"net/http"
)

// UserRegister 用户注册
// @tags 用户
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
		form     = new(utils.AddUserForm)
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

// Login 登录
// @tags 用户
// @Summary 用户登录
// @Produce json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/login [post]
func Login(c *gin.Context) {
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
	if token, err = srv_user.UserLogin(form, false); err != nil {
		return
	}

}

// UserUpdate 用户更新基本信息
// @tags 用户
// @Summary 更新用户信息
// @Produce json
// @Param nickname body string false "昵称"
// @Param username body string false "用户名,之可以修改一次"
// @Param Phone body string false "手机号码"
// @Param sign body string false "签名"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/userUpdate [post]
func UserUpdate(c *gin.Context) {
	var (
		err      error
		form     = new(utils.FormEditUser)
		data     *models.UserBase
		httpCode = http.StatusInternalServerError
		s        *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGin(c).Response(httpCode, err.Error())
		} else {
			app.NewGin(c).Response(200, data)
		}
	}()
	if s, err = getSession(c); err != nil {
		return
	}
	if err = c.ShouldBind(form); err != nil {
		httpCode = http.StatusBadRequest
		err = contrib.ErrParamInvalid
		return
	}
	if data, err = srv_user.EditUser(s.UID, form); err != nil {
		return
	}
}

// GetUserProfile 获取用户信息
// @tags 用户
// @Summary 获取个人信息
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/getUserProfile [get]
func GetUserProfile(c *gin.Context) {
	var (
		err      error
		data     = new(models.UserBase)
		httpCode = http.StatusInternalServerError
		s        *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGin(c).Response(httpCode, err.Error())
		} else {
			app.NewGin(c).Response(200, data)
		}
	}()
	if s, err = getSession(c); err != nil {
		return
	}
	if err = data.GetUserByUId(s.UID); err != nil {
		return
	}

}

// ChangePayPasswd 修改密码
// @tags 用户
// @Summary 修改用户密码
// @Produce json
// @Accept json
// @Param oldPasswd query string true "旧密码"
// @Param newPasswd query string true "新密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/changeLoginPasswd [post]
func ChangeLoginPasswd(c *gin.Context) {
	var (
		err            error
		httpCode       = http.StatusInternalServerError
		s              *session.Session
		oldPwd, newPwd string
	)
	defer func() {
		if err != nil {
			app.NewGin(c).Response(httpCode, err.Error())
		} else {
			app.NewGin(c).Response(200, "")
		}
	}()
	oldPwd = c.Param("oldPasswd")
	newPwd = c.Param("newPasswd")
	if s, err = getSession(c); err != nil {
		return
	}
	if err = srv_user.ChangePasswd(s.UID, oldPwd, newPwd); err != nil {
		return
	}
}

// ChangePayPasswd 修改密码
// @tags 用户
// @Summary 修改用户密码
// @Produce json
// @Accept json
// @Param oldPasswd query string true "旧密码"
// @Param newPasswd query string true "新密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/changeLoginPasswd [post]
func ChangePayPasswd(c *gin.Context) {
	var (
		err            error
		httpCode       = http.StatusInternalServerError
		s              *session.Session
		oldPwd, newPwd string
	)
	defer func() {
		if err != nil {
			app.NewGin(c).Response(httpCode, err.Error())
		} else {
			app.NewGin(c).Response(200, "")
		}
	}()
	oldPwd = c.Param("oldPasswd")
	newPwd = c.Param("newPasswd")
	if s, err = getSession(c); err != nil {
		return
	}
	if err = srv_user.ChangePayPasswd(s.UID, oldPwd, newPwd); err != nil {
		return
	}
}

// EditHeadIcon 修改用户头像
// @tags 用户
// @Summary 修改用户头像
// @Produce json
// @Accept  json
// @Param headIcon query string true '头像'
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
func EditHeadIcon(c *gin.Context) {

	var (
		err      error
		s        *session.Session
		headIcon string
	)
	headIcon = c.Query("headIcon")
	if s, err = getSession(c); err != nil {
		return
	}
	_ = s
	_ = headIcon

}
