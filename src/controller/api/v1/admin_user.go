package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/gorm-gin-admin/src/models"
	"github.com/olongfen/gorm-gin-admin/src/pkg/app"
	"github.com/olongfen/gorm-gin-admin/src/pkg/codes"
	"github.com/olongfen/gorm-gin-admin/src/service"
	"github.com/olongfen/gorm-gin-admin/src/utils"
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
		form  = &utils.LoginForm{}
		err   error
		code  = codes.CodeProcessingFailed
		token string
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			app.NewGinResponse(c).Success(map[string]string{"token": token}).Response()
		}
	}()

	form.IP = c.ClientIP()
	if err = c.ShouldBind(form); err != nil {
		code = codes.CodeParamInvalid
		return
	}
	if token, err = service.UserLogin(form, true); err != nil {
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
		err  error
		code = codes.CodeProcessingFailed
		s    *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			app.NewGinResponse(c).Success(nil).Response()
		}
	}()
	if s, code, err = GetSession(c); err != nil {
		return
	}
	if err = service.UserLogout(s.UID); err != nil {
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
// @Router /api/v1/admin/listUser [get]
func ListUser(c *gin.Context) {
	var (
		err  error
		code = codes.CodeProcessingFailed
		form = new(utils.FormUserList)
		data []*models.UserBase
		s    *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		}
	}()

	if s, code, err = GetSession(c); err != nil {
		return
	}
	if err = c.ShouldBindQuery(form); err != nil {
		code = codes.CodeParamInvalid
		return
	}
	if data, err = service.GetUserList(s.UID, form); err != nil {
		return
	}
	app.NewGinResponse(c).Success(data).Response()
}

// @tags 管理员
// @Title 获取用户总数
// @Summary 获取用户总数
// @Description 获取用户总数
// @Accept json
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router   /api/v1/admin/userTotal  [get]
func UserTotal(c *gin.Context) {
	var (
		err  error
		code = codes.CodeProcessingFailed
		data int64
		s    *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		}
	}()

	if s, code, err = GetSession(c); err != nil {
		return
	}

	if data, err = service.GetUserCount(s.UID); err != nil {
		return
	}
	app.NewGinResponse(c).Success(data).Response()
}

// @tags 管理员
// @Title
// @Summary
// @Description
// @Accept json
// @Produce json
// @Param uid path string true "用户uid"
// @Param {} body utils.FormEditUser true "修改用户信息"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/admin/editUser [post]
func EditUser(c *gin.Context) {
	var (
		err  error
		form = new(utils.FormEditUser)
		s    *session.Session

		code = codes.CodeProcessingFailed
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		}else {
			app.NewGinResponse(c).Success(nil).Response()
		}
	}()
	if s, code, err = GetSessionAndBindingForm(form, c); err != nil {
		return
	}
	if s.UID==form.Uid{
		err = utils.ErrActionNotAllow.SetMeta("")
		return
	}
	if _, err = service.EditUserByRole(s.UID, form); err != nil {
		return
	}
}

// @tags 管理员
// @Summary 删除用户
// @Description 删除用户
// @Accept json
// @Produce json
// @Param uid query string true "用户UID"
// @Success 200  {object} app.Response
// @Failure 500  {object} app.Response
// @router /api/v1/admin/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	var (
		err error
		uid string
		s   *session.Session

		code = codes.CodeProcessingFailed
	)
	uid = c.Query("uid")
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			app.NewGinResponse(c).Success(nil).Response()
		}
	}()
	if s, code, err = GetSession(c); err != nil {
		return
	}
	if s.UID == uid {
		err = utils.ErrActionNotAllow
		return
	}
	if err = service.DeleteUser(s.UID, uid); err != nil {
		return
	}
}

// @tags 管理员
// @Summary 获取用户表状态信息
// @Description 获取用户表状态信息
// @Accept json
// @Produce json
// @Success 200  {object} app.Response
// @Failure 500  {object} app.Response
// @router  /api/v1/admin/getUserKV [get]
func GetUserKV(c *gin.Context)  {
	app.NewGinResponse(c).Success(service.GetUserKV()).Response()
}

// @tags 管理员
// @Summary 添加用户
// @Description 添加用户
// @Accept json
// @Produce json
// @Param {object} body utils.AddUserForm true "添加用户form"
// @Success 200  {object} app.Response
// @Failure 500  {object} app.Response
// @router  /api/v1/admin/addUser [post]
func AddUser(c *gin.Context)  {
	var (
		form = new(utils.AdminAddUserForm)
		data *models.UserBase
		code = codes.CodeProcessingFailed
		s *session.Session
		err  error
	)

	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			app.NewGinResponse(c).Success(data).Response()
		}
	}()
	if s,code,err = GetSessionAndBindingForm(form,c);err!=nil{
		return
	}
	if data,err = service.AdminAddUser(s.UID,form);err!=nil{
		return
	}

}