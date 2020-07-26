package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/pkg/app"
	"github.com/olongfen/user_base/service"
	"github.com/olongfen/user_base/utils"
	"net/http"
	"strconv"
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
			app.NewGinResponse(c).SetCodeAndMessage(httpCode,err.Error()).Response()
		} else {
			app.NewGinResponse(c).SetCodeAndMessage(200, "success").SetData(map[string]string{"token": token}).Response()
		}
	}()

	form.IP = c.ClientIP()
	if err = c.ShouldBind(form); err != nil {
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
		err      error
		httpCode = http.StatusInternalServerError
		s        *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(httpCode,err.Error()).Response()
		} else {
			app.NewGinResponse(c).SetCodeAndMessage(200,"success").Response()
		}
	}()
	if s, err = GetSession(c); err != nil {
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
		err      error
		httpCode = http.StatusInternalServerError
		form     = new(utils.FormUserList)
		data     []*models.UserBase
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(httpCode,err.Error()).Response()
		}
	}()

	if _, err = GetSession(c); err != nil {
		return
	}
	if err = c.ShouldBindQuery(form); err != nil {
		return
	}
	if data, err = service.GetUserList(form); err != nil {
		return
	}
	app.NewGinResponse(c).SetCodeAndMessage(httpCode,"success").SetData(data).Response()
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
		// s    *session.Session

		code int
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(code, err.Error()).Response()
		}
	}()
	if _, code, err = GetSessionAndBindingForm(form, c); err != nil {
		return
	}
	if _, err = service.EditUser(form); err != nil {
		code = 500
		return
	}
}

// @tags 管理员
// @Title 添加角色接口权限
// @Summary 添加角色接口权限
// @Description 添加角色接口权限
// @Accept json
// @Produce json
// @Param {} body utils.FormRoleAPIPerm true "添加api权限表单"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/admin/addRoleAPIPerm [post]
func AddRoleAPIPerm(c *gin.Context) {
	var (
		f    = &utils.FormRoleAPIPerm{}
		err  error
		code int
		ret  []int64
		s    *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(code, err.Error()).Response()
		}
	}()
	if s, code, err = GetSessionAndBindingForm(f, c); err != nil {
		return
	}
	if s.UID == f.Uid {
		err = utils.ErrActionNotAllow
		return
	}
	if ret, err = service.AddRuleAPI(f); err != nil {
		code = 500
		return
	}
app.NewGinResponse(c).SetCodeAndMessage(200,"success").SetData(ret).Response()
}

// @tags 管理员
// @Title 删除角色接口权限
// @Summary 删除角色接口权限
// @Description 删除角色接口权限
// @Accept json
// @Produce json
// @Param {} body utils.FormRoleAPIPerm true "删除api权限表单"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/admin/removeRoleAPIPerm [delete]
func RemoveRolePermAPI(c *gin.Context) {
	var (
		f    = &utils.FormRoleAPIPerm{}
		err  error
		code int
		ret  []int64
		s    *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(code, err.Error()).Response()
		}
	}()
	if s, code, err = GetSessionAndBindingForm(f, c); err != nil {
		return
	}
	if s.UID == f.Uid {
		err = utils.ErrActionNotAllow
		return
	}
	if ret, err = service.RemoveRuleAPI(f); err != nil {
		code = 500
		return
	}
app.NewGinResponse(c).SetCodeAndMessage(200,"success").SetData(ret).Response()
}

// @tags 管理员
// @Title 获取用户权限
// @Summary 获取用户权限
// @Description 获取用户权限
// @Accept json
// @Produce json
// @Param uid query string false "用户uid,不输入默认返回自己uid"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/getRoleApiList [get]
func GetRoleApiList(c *gin.Context) {
	var (
		err  error
		code = 500
		s    *session.Session
		uid  string
		data []struct {
			Path   string
			Method string
		}
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(code, err.Error()).Response()
		}
	}()
	if s, err = GetSession(c); err != nil {
		code = 401
		return
	}
	uid = c.Query("uid")
	if len(uid) == 0 {
		uid = s.UID
	}
	if data, err = service.GetRuleApiList(uid); err != nil {
		return
	}
	app.NewGinResponse(c).SetCodeAndMessage(200,"success").SetData(data).Response()

}

// @tags 管理员
// @Title 获取全部api
// @Summary 获取全部api
// @Description 获取全部api
// @Accept json
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router  /api/v1/admin/getAllApiGroup [get]
func GetAllAPIGroup(c *gin.Context) {
	var (
		err error
		ret []*models.APIGroup
		code =500
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(code,err.Error()).Response()
		}
	}()
	if _, err = GetSession(c); err != nil {
		code = 401
		return
	}
	if ret, err = service.GetAPIGroupList(); err != nil {
		return
	}
app.NewGinResponse(c).SetCodeAndMessage(200,"success").SetData(ret).Response()
}

// @tags 管理员
// @Title 创建api
// @Summary 创建api
// @Description
// @Accept json
// @Produce json
// @Param {array} body utils.FormAPIGroupAdd true "api数组"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/admin/addApiGroup [post]
func AddApiGroup(c *gin.Context) {
	var (
		err  error
		code int
		f    []*utils.FormAPIGroupAdd
		ret  []*models.APIGroup
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(code, err.Error()).Response()
		}
	}()
	if _, code, err = GetSessionAndBindingForm(f, c); err != nil {
		return
	}
	if ret, err = service.AddAPIGroup(f); err != nil {
		code = 500
		return
	}
app.NewGinResponse(c).SetCodeAndMessage(200,"success").SetData(ret).Response()
}

// @tags 管理员
// @Title 删除api
// @Summary  删除api
// @Description 删除api
// @Accept json
// @Produce json
// @Param id query int true "id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/admin/removeApiGroup [delete]
func RemoveApiGroup(c *gin.Context) {
	var (
		err  error
		code int
		id   string
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(code, err.Error()).Response()
		}
	}()
	id = c.Query("id")
	_id, err_ := strconv.ParseUint(id, 10, 64)
	if err_ != nil {
		code = 404
		err = err_
		return
	}
	if err = service.DelAPIGroup(int64(_id)); err != nil {
		code = 500
		return
	}
	app.NewGinResponse(c).SetCodeAndMessage(200,"success").Response()
}

// @tags 管理员
// @Title 修改api
// @Summary 修改api
// @Description 修改api
// @Accept json
// @Produce json
// @Param {} body utils.FormAPIGroupEdit true "表单"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router  /api/v1/admin/editApiGroup [put]
func EditApiGroup(c *gin.Context) {
	var (
		f    = &utils.FormAPIGroupEdit{}
		err  error
		code int
		ret  *models.APIGroup
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(code, err.Error()).Response()
		}
	}()
	if _, code, err = GetSessionAndBindingForm(f, c); err != nil {
		return
	}

	if ret, err = service.EditAPIGroup(f); err != nil {
		code = 500
		return
	}
app.NewGinResponse(c).SetCodeAndMessage(200,"success").SetData(ret).Response()

}
