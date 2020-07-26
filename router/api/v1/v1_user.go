package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/pkg/app"
	"github.com/olongfen/user_base/pkg/setting"
	"github.com/olongfen/user_base/service"
	"github.com/olongfen/user_base/utils"
	uuid "github.com/satori/go.uuid"
	"image"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// UserRegister 用户注册
// @tags 用户
// @Summary 用户注册
// @Produce json
// @Param phone body string true "Phone"
// @Param password body string true "Password"
// @Param code body string  false "Code"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/register [post]
func UserRegister(c *gin.Context) {
	var (
		form     = new(utils.AddUserForm)
		data     *models.UserBase
		httpCode = http.StatusInternalServerError
		err      error
	)

	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(httpCode,err.Error()).Response()
		} else {
			app.NewGinResponse(c).SetCodeAndMessage(httpCode,"success").SetData(data).Response()
		}
	}()
	if err = c.Bind(form); err != nil {
		httpCode = http.StatusBadRequest
		err = contrib.ErrParamInvalid
		return
	}
	if data, err = service.AddUser(form); err != nil {
		return
	}
}

// UserLogin 登录
// @tags 用户
// @Summary 用户登录
// @Produce json
// @Param {} body utils.LoginForm true "form"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/user/login [post]
func UserLogin(c *gin.Context) {
	var (
		form     = &utils.LoginForm{}
		err      error
		httpCode = 500
		token    string
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(httpCode,err.Error()).Response()
		} else {
			app.NewGinResponse(c).SetCodeAndMessage(200,"success").SetData(map[string]string{"token": token}).Response()
		}
	}()
	form.IP = c.ClientIP()
	if err = c.ShouldBind(form); err != nil {
		return
	}
	if token, err = service.UserLogin(form, false); err != nil {
		return
	}

}

// UserLogin 登出
// @tags 用户
// @Summary 用户登出
// @Produce json
// @Accept json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/user/logout [post]
func UserLogout(c *gin.Context) {
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

// ModifyProfile 用户更新基本信息
// @tags 用户
// @Summary 更新用户信息
// @Produce json
// @Param nickname body string false "昵称"
// @Param Phone body string false "手机号码"
// @Param sign body string false "签名"
// @Param email body string false "邮箱"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/user/modifyProfile [put]
func ModifyProfile(c *gin.Context) {
	var (
		err      error
		form     = new(utils.FormEditUser)
		data     *models.UserBase
		httpCode = http.StatusInternalServerError
		s        *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(httpCode,err.Error()).Response()
		} else {
			app.NewGinResponse(c).SetCodeAndMessage(httpCode,"success").SetData(data).Response()
		}
	}()
	if s, err = GetSession(c); err != nil {
		return
	}
	if err = c.ShouldBind(form); err != nil {
		httpCode = http.StatusBadRequest
		err = contrib.ErrParamInvalid
		return
	}
	form.Uid = s.UID
	if data, err = service.EditUser(form); err != nil {
		return
	}
}

// GetUserProfile 获取用户信息
// @tags 用户
// @Summary 获取个人信息
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/user/getUserProfile [get]
func GetUserProfile(c *gin.Context) {
	var (
		err      error
		data     = new(models.UserBase)
		httpCode = http.StatusInternalServerError
		s        *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(httpCode,err.Error()).Response()
		} else {
			app.NewGinResponse(c).SetCodeAndMessage(httpCode,"success").SetData(data).Response()
		}
	}()
	if s, err = GetSession(c); err != nil {
		return
	}
	if err = data.GetUserByUId(s.UID); err != nil {
		return
	}

}

// ModifyLoginPwd 修改密码
// @tags 用户
// @Summary 修改用户密码
// @Produce json
// @Accept json
// @Param oldPwd body string true "旧密码"
// @Param newPwd body string true "新密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/user/modifyLoginPwd [put]
func ModifyLoginPwd(c *gin.Context) {
	var (
		err      error
		httpCode = http.StatusInternalServerError
		s        *session.Session
		f        struct {
			OldPwd string `form:"oldPwd" binding:"required"`
			NewPwd string `form:"newPwd" binding:"required"`
		}
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(httpCode,err.Error()).Response()
		} else {
			app.NewGinResponse(c).SetCodeAndMessage(200,"message").Response()
		}
	}()
	if err = c.ShouldBind(&f); err != nil {
		httpCode = 400
		return
	}
	if s, err = GetSession(c); err != nil {
		return
	}
	if err = service.ChangePwd(s.UID, f.OldPwd, f.NewPwd); err != nil {
		return
	}
}

// ModifyPayPwd 修改密码
// @tags 用户
// @Summary 修改用户密码
// @Produce json
// @Accept json
// @Param oldPwd body string true "旧密码"
// @Param newPwd body string true "新密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/user/modifyPayPwd [put]
func ModifyPayPwd(c *gin.Context) {
	var (
		err      error
		httpCode = http.StatusInternalServerError
		s        *session.Session
		f        struct {
			OldPwd string `form:"oldPwd" binding:"required"`
			NewPwd string `form:"newPwd" binding:"required"`
		}
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(httpCode,err.Error()).Response()
		} else {
			app.NewGinResponse(c).SetCodeAndMessage(200,"message").Response()
		}
	}()
	if err = c.ShouldBind(&f); err != nil {
		httpCode = 400
		return
	}
	if s, err = GetSession(c); err != nil {
		return
	}
	if err = service.ChangePayPwd(s.UID, f.OldPwd, f.NewPwd); err != nil {
		return
	}
}

// ModifyHeadIcon 修改用户头像
// @tags 用户
// @Summary 修改用户头像
// @Produce json
// @Accept  json
// @Param headIcon body string true "头像"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/modifyHeadIcon [put]
func ModifyHeadIcon(c *gin.Context) {

	var (
		err      error
		s        *session.Session
		headIcon *multipart.FileHeader
		httpCode = http.StatusInternalServerError
		d        = new(models.UserBase)
		img      image.Image
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(httpCode,err.Error()).Response()
		} else {
			app.NewGinResponse(c).SetCodeAndMessage(200,"message").SetData(gin.H{
				"headIcon": d.HeadIcon,
			}).Response()
		}
	}()
	if headIcon, err = c.FormFile("headIcon"); err != nil {
		httpCode = http.StatusBadRequest
		return
	}
	_f, _ := headIcon.Open()
	if img, _, err = image.Decode(_f); err != nil {
		return
	}
	b := img.Bounds()
	if b.Max.X > 300 || b.Max.Y > 300 {
		err = utils.ErrImagePixelToBig
		return
	}
	// 最高能够保存500kb的头像
	if headIcon.Size > 2<<20 {
		err = utils.ErrImageSizeToBig
		return
	}

	if s, err = GetSession(c); err != nil {
		return
	}
	if err = d.GetUserByUId(s.UID); err != nil {
		return
	}
	oldDst := d.HeadIcon
	//
	arr := strings.Split(headIcon.Filename, ".")
	dst := setting.ProjectSetting.HeadIconDir + uuid.NewV4().String() + "." + arr[len(arr)-1]
	if err = c.SaveUploadedFile(headIcon, dst); err != nil {
		return
	}

	if err = d.UpdateUserOne(s.UID, "head_icon", dst); err != nil {
		return
	}
	os.Remove(oldDst)

}

// GetHeadIcon 获取用户头像
// @tags 用户
// @Summary 获取用户头像
// @Produce json
// @Accept json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/getHeadIcon [get]
func GetHeadIcon(c *gin.Context) {
	var (
		err  error
		s    *session.Session
		data = new(models.UserBase)
		httpCode =500
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(httpCode,err.Error()).Response()
		}
	}()
	if s, err = GetSession(c); err != nil {
		httpCode =401
		return
	}
	if err = data.GetUserByUId(s.UID); err != nil {
		return
	}
	c.File(data.HeadIcon)
}

// @tags 用户
// @Title 用户设置支付密码
// @Summary 用户设置支付密码
// @Description 用户设置支付密码
// @Accept json
// @Produce json
// @Param pwd body string true "密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/user/setPayPwd/ [post]
func SetPayPwd(c *gin.Context) {
	var (
		sess     *session.Session
		err      error
		httpCode = 500
		pwd      string
	)
	defer func() {
		if err != nil {
			app.NewGinResponse(c).SetCodeAndMessage(httpCode,err.Error()).Response()
		}
	}()
	if pwd = c.PostForm("pwd"); len(pwd) == 0 {
		httpCode = 400
		err = contrib.ErrParamInvalid
		return
	}
	if sess, err = GetSession(c); err != nil {
		return
	}
	if err = service.SetUserPayPwd(sess.UID, pwd); err != nil {
		return
	}
	app.NewGinResponse(c).SetCodeAndMessage(200,"success").Response()
}
