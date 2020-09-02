package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/gorm-gin-admin/src/models"
	"github.com/olongfen/gorm-gin-admin/src/pkg/app"
	"github.com/olongfen/gorm-gin-admin/src/pkg/codes"
	"github.com/olongfen/gorm-gin-admin/src/service"
	"github.com/olongfen/gorm-gin-admin/src/utils"
	"strconv"
)

// @tags 管理员
// @Title 获取全部api
// @Summary 获取全部api
// @Description 获取全部api
// @Param {} body utils.ApiListForm true "获取api列表"
// @Accept json
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Path  /api/v1/admin/getAllApiGroup [get]
func GetAPIGroupList(c *gin.Context) {
	var (
		err  error
		code = codes.CodeProcessingFailed
		ret  []*models.APIGroup
		form =new(utils.ApiListForm)
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		}
	}()
	if _, code, err = GetSessionAndBindingForm(form,c); err != nil {
		return
	}
	if ret, err = service.GetAPIGroupList(form); err != nil {
		return
	}
	app.NewGinResponse(c).Success(ret).Response()
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
// @Path /api/v1/admin/addApiGroup [post]
func AddApiGroup(c *gin.Context) {
	var (
		err  error
		code = codes.CodeProcessingFailed
		f    []*utils.FormAPIGroupAdd
		ret  []*models.APIGroup
		s *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		}
	}()
	if s, code, err = GetSessionAndBindingForm(&f, c); err != nil {
		return
	}
	if ret, err = service.AddAPIGroup(s.UID,f); err != nil {

		return
	}
	app.NewGinResponse(c).Success(ret).Response()
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
// @Path /api/v1/admin/removeApiGroup [delete]
func RemoveApiGroup(c *gin.Context) {
	var (
		err  error
		s *session.Session
		code = codes.CodeProcessingFailed
		id   string
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		}
	}()
	id = c.Query("id")
	_id, err_ := strconv.ParseUint(id, 10, 64)
	if err_ != nil {
		code = codes.CodeParamInvalid
		err = err_
		return
	}
	if s,code,err = GetSession(c);err!=nil{
		return
	}
	if err = service.DelAPIGroup(s.UID,int64(_id)); err != nil {
		return
	}
	app.NewGinResponse(c).Success(nil).Response()
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
// @Path  /api/v1/admin/editApiGroup [put]
func EditApiGroup(c *gin.Context) {
	var (
		f    = &utils.FormAPIGroupEdit{}
		err  error
		code = codes.CodeProcessingFailed
		ret  *models.APIGroup
		s *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		}
	}()
	if s, code, err = GetSessionAndBindingForm(f, c); err != nil {
		return
	}

	if ret, err = service.EditAPIGroup(s.UID,f); err != nil {

		return
	}
	app.NewGinResponse(c).Success(ret).Response()

}
