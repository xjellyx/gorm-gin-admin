package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/pkg/app"
	"github.com/olongfen/user_base/service"
	"github.com/olongfen/user_base/utils"
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
		err      error
		data     []*models.Menu
		form     []*utils.FormAddMenu
		httpCode = 500
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(httpCode, err.Error()).Response()
		} else {
			app.NewGinResponse(c).Success(data).Response()
		}
	}()
	if err = c.ShouldBind(&form); err != nil {
		httpCode = 400
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
// @Param id int64 true "菜单id"
// @Success 200 {object} app.Response
// @Failure 500  {object} app.Response
// @router /api/v1/admin/getMenu [get]
func GetMenu(c *gin.Context) {
	var (
		err  error
		id   int
		code = 500
		data *models.Menu
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		} else {
			app.NewGinResponse(c).Success(data).Response()
		}
	}()
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		code = 400
		return
	}
	if data, err = service.GetMenu(id); err != nil {
		return
	}

}
