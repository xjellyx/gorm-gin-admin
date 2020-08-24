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
// @Failure 200 {object} app.Response
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
// @Param id query int64 false "菜单id"
// @Success 200 {object} app.Response
// @Failure 200  {object} app.Response
// @router /api/v1/admin/getMenu [get]
func GetMenu(c *gin.Context) {
	var (
		err  error
		id   int
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
	if id, err = strconv.Atoi(c.Query("id")); err != nil {
		code = codes.CodeParamInvalid
		return
	}
	if id == 0 {
		id = 1
	}
	if data, err = service.GetMenu(id); err != nil {
		return
	}

}
