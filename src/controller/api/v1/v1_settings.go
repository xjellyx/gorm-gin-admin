package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/olongfen/gorm-gin-admin/src/pkg/app"
	"github.com/olongfen/gorm-gin-admin/src/pkg/codes"
	"github.com/olongfen/gorm-gin-admin/src/setting"
	"github.com/spf13/viper"
)

// @tags 管理员
// @Summary　获取系统配置
// @Description 获取系统配置
// @Accept json
// @Produce json
// @Success 200  {object} app.Response
// @Failure 500  {object} app.Response
// @router /api/v1/admin/getSettings  [get]
func GetSettings(ctx *gin.Context) {

	var (
		err  error
		code = codes.CodeProcessingFailed
		data = new(setting.Project)
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(ctx).Fail(code, err.Error()).Response()
		} else {
			app.NewGinResponse(ctx).Success(data).Response()
		}
	}()
	if _, code, err = GetSession(ctx); err != nil {
		return
	}
	d := viper.Get("project")
	if err = mapstructure.Decode(d, data); err != nil {
		return
	}
	return
}
