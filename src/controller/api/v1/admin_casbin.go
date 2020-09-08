package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/gorm-gin-admin/src/pkg/app"
	"github.com/olongfen/gorm-gin-admin/src/pkg/codes"
	"github.com/olongfen/gorm-gin-admin/src/service"
	"github.com/olongfen/gorm-gin-admin/src/utils"
)

// @tags 管理员
// @Title 添加角色接口权限
// @Summary 添加角色接口权限
// @Description 添加角色接口权限
// @Accept json
// @Produce json
// @Param {} body utils.FormRoleAPIPerm true "添加api权限表单"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Path /api/v1/admin/addRoleAPIPerm [post]
func AddRoleAPIPerm(c *gin.Context) {
	var (
		f    = &utils.FormRoleAPIPerm{}
		err  error
		code = codes.CodeProcessingFailed
		ret  []struct{
			Method string `json:"method"`
			Path string `json:"path"`
		}
		s    *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		}
	}()
	if s, code, err = GetSessionAndBindingForm(f, c); err != nil {
		return
	}

	if ret, err = service.AddRoleAPI(s.UID,f); err != nil {
		return
	}
	app.NewGinResponse(c).Success(ret).Response()
}

// @tags 管理员
// @Title 添加角色组
// @Summary 添加角色组
// @Description 添加角色组
// @Accept json
// @Produce json
// @Param {} body utils.FormRoleAPIPerm true "添加角色组"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Path /api/v1/admin/addRoleGroup [post]
func AddRoleGroup(c *gin.Context)  {
	
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
// @Path /api/v1/admin/removeRoleAPIPerm [delete]
func RemoveRolePermAPI(c *gin.Context) {
	var (
		f    = &utils.FormRoleAPIPerm{}
		err  error
		code = codes.CodeProcessingFailed

		s    *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		}
	}()
	if s, code, err = GetSessionAndBindingForm(f, c); err != nil {
		return
	}
	if  err = service.RemoveRoleAPI(s.UID,f); err != nil {
		return
	}
	app.NewGinResponse(c).Success(nil).Response()
}

// @tags 管理员
// @Title 获取用户权限
// @Summary 获取用户权限
// @Description 获取用户权限
// @Accept json
// @Produce json
// @Param role query string true "角色名称"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Path /api/v1/admin/getRoleApiList [get]
func GetRoleApiList(c *gin.Context) {
	var (
		err  error
		code = codes.CodeProcessingFailed
		role string
		data  []service.RoleApiResp
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		}else {
			app.NewGinResponse(c).Success(data).Response()
		}
	}()
	if _, code, err = GetSession(c); err != nil {
		return
	}
	role = c.Query("role")
	if len(role) == 0 {
		return
	}
	if data, err = service.GetRoleApiList(role); err != nil {
		return
	}


}



