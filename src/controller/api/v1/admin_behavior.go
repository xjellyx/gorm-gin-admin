package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/gorm-gin-admin/src/models"
	"github.com/olongfen/gorm-gin-admin/src/pkg/app"
	"github.com/olongfen/gorm-gin-admin/src/pkg/codes"
	"github.com/olongfen/gorm-gin-admin/src/service"
	"github.com/olongfen/gorm-gin-admin/src/utils"
)

// @tags 管理员
// @Summary 获取操作记录
// @Description 获取操作记录
// @Accept json
// @Produce json
// @Param {object} query utils.BehaviorQueryForm  true "請求form"
// @Success 200  {object}  app.Response
// @Failure 500  {object}  app.Response
// @router /api/v1/admin/getBehaviorList  [get]
func GetBehaviorList(c *gin.Context) {
	var (
		param = new(utils.BehaviorQueryForm)
		err   error
		code  = codes.CodeProcessingFailed
		ret   []*models.BehaviorRecord
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		}
	}()
	if _, code, err = GetSessionAndBindingForm(param, c); err != nil {
		return
	}
	if ret, err = service.GetBehaviorList(param); err != nil {
		return
	}

	app.NewGinResponse(c).Success(ret).Response()
}

// @tags
// @Summary
// @Description
// @Accept json
// @Produce json
// @Param {}  body utils.BehaviorRemoveForm  true "id list"
// @Success 200  {object} app.Response
// @Failure 500  {object} app.Response
// @router /api/v1/admin/removeBehaviors  [delete]
func RemoveBehaviors(c *gin.Context) {
	var (
		form = new(utils.BehaviorRemoveForm)
		err  error
		code = codes.CodeProcessingFailed

		s *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		}
	}()
	if s, code, err = GetSessionAndBindingForm(form, c); err != nil {
		return
	}
	if err = service.RemoveBehavior(s.UID, form.Ids); err != nil {
		return
	}
	_ = models.NewActionRecord(s, c, fmt.Sprintf(`remove  behavior record  %v `, form.Ids)).Insert()
	app.NewGinResponse(c).Success(nil).Response()
}

// @tags
// @Summary
// @Description
// @Accept json
// @Produce json
// @Success 200  {object} app.Response
// @Failure 500  {object} app.Response
// @router /api/v1/admin/behaviorCount  [get]
func BehaviorCount(c *gin.Context) {
	var (
		err   error
		code  = codes.CodeProcessingFailed
		count int64
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).Fail(code, err.Error()).Response()
		}
	}()
	if count, err = models.BehaviorCount(); err != nil {
		return
	}
	app.NewGinResponse(c).Success(count).Response()
}
