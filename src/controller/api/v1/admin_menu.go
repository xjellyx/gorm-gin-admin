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

// @tags 超级管理员
// @Summary 添加菜单
// @Produce json
// @Param {} body utils.FormAddMenu true "菜单form"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/admin/addMenu [post]
func AddMenu(c *gin.Context) {
	var (
		err  error
		data []*models.Menu
		form []*utils.FormAddMenu
		s    *session.Session
		role = &models.UserBase{}
		code = codes.CodeProcessingFailed
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			app.NewGinResponse(c).Success(data).Response()
		}
	}()
	if s, code, err = GetSessionAndBindingForm(&form, c); err != nil {
		return
	}
	if err = role.GetByUId(s.UID); err != nil {
		return
	}
	if role.Role < models.UserRoleSuperAdmin {
		err = utils.ErrActionNotAllow
		return
	}
	if data, err = service.AddMenu(form); err != nil {
		return
	}
}

// @tags 管理员
// @Summary 获取菜单
// @Description 获取菜单
// @Accept json
// @Produce json
// @Param id query int64 true "菜单id"
// @Success 200 {object} app.Response
// @Failure 500  {object} app.Response
// @router /api/v1/admin/getMenu [get]
func GetMenu(c *gin.Context) {
	var (
		err  error
		id   int
		code = codes.CodeProcessingFailed
		data *models.Menu
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			app.NewGinResponse(c).Success(data).Response()
		}
	}()
	_id := c.Query("id")
	if len(_id) == 0 {
		err = utils.ErrParamInvalid
		return
	}
	if id, err = strconv.Atoi(_id); err != nil {
		code = codes.CodeParamInvalid
		return
	}
	if data, err = service.GetMenu(id); err != nil {
		return
	}

}

// @tags 管理员
// @Summary 获取菜单
// @Description 获取菜单
// @Accept json
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500  {object} app.Response
// @router /api/v1/admin/getMenuList [get]
func GetMenuList(c *gin.Context) {
	var (
		err  error
		code = codes.CodeProcessingFailed
		data []*models.Menu
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			app.NewGinResponse(c).Success(data).Response()
		}
	}()

	if data, err = service.GetMenuList(); err != nil {
		return
	}

}

// @tags 超级管理员
// @Title 删除菜单
// @Summary 删除菜单
// @Description 删除菜单
// @Accept json
// @Produce json
// @Param id param int true "id"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Path /api/v1/admin/delMenu [delete]
func DelMenu(c *gin.Context) {
	var (
		_id  int
		err  error
		s    *session.Session
		code = codes.CodeProcessingFailed
		role = new(models.UserBase)
	)
	id := c.Query("id")
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
	if err = role.GetByUId(s.UID); err != nil {
		return
	}
	if role.Role < models.UserRoleSuperAdmin {
		err = utils.ErrActionNotAllow
		return
	}
	if _id, err = strconv.Atoi(id); err != nil {
		code = codes.CodeParamInvalid
		err = utils.ErrParamInvalid
		return
	}
	if err = service.DelMenu(_id); err != nil {
		return
	}
}

// @tags 超级管理员
// @Title 更新菜单
// @Summary 更新菜单
// @Description 更新菜单
// @Accept json
// @Produce json
// @Param {} body utils.FormUpdateMenu true "菜单form"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Path  /api/v1/admin/editMenu [put]
func EditMenu(c *gin.Context) {
	var (
		err  error
		code = codes.CodeProcessingFailed
		s    *session.Session
		data *models.Menu
		role = new(models.UserBase)
		form = new(utils.FormUpdateMenu)
	)

	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			app.NewGinResponse(c).Success(data).Response()
		}
	}()
	if s, code, err = GetSessionAndBindingForm(form, c); err != nil {
		return
	}
	if err = role.GetByUId(s.UID); err != nil {
		return
	}
	if role.Role < models.UserRoleSuperAdmin {
		err = utils.ErrActionNotAllow
		return
	}
	if data, err = service.UpdateMenu(form); err != nil {
		return
	}
}

