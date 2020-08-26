package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/user_base/src/models"
	"github.com/olongfen/user_base/src/pkg/app"
	"github.com/olongfen/user_base/src/pkg/codes"
	"github.com/olongfen/user_base/src/service"
	"github.com/olongfen/user_base/src/utils"
	"strconv"
)

// @tags 管理员
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
		code = codes.CodeProcessingFailed
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			app.NewGinResponse(c).Success(data).Response()
		}
	}()
	if err = c.ShouldBind(&form); err != nil {
		code = codes.CodeParamInvalid
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
