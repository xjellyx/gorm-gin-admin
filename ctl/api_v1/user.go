package api_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/user_base/models"
	"github.com/olongfen/user_base/pkg/app"
	"github.com/olongfen/user_base/pkg/setting"
	"github.com/olongfen/user_base/service/srv_user"
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
			app.NewGin(c).Response(httpCode, err.Error())
		} else {
			app.NewGin(c).Response(200, data)
		}
	}()
	if err = c.Bind(form); err != nil {
		httpCode = http.StatusBadRequest
		err = contrib.ErrParamInvalid
		return
	}
	if data, err = srv_user.AddUser(form); err != nil {
		return
	}
}

// Login 登录
// @tags 用户
// @Summary 用户登录
// @Produce json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/login [post]
func Login(c *gin.Context) {
	var (
		form     = &utils.LoginForm{}
		err      error
		httpCode = http.StatusInternalServerError
		token    string
	)
	defer func() {
		if err != nil {
			app.NewGin(c).Response(httpCode, err.Error())
		} else {
			app.NewGin(c).Response(200, map[string]string{"token": token})
		}
	}()
	form.IP = c.ClientIP()
	if err = c.ShouldBind(form); err != nil {
		return
	}
	if token, err = srv_user.UserLogin(form, false); err != nil {
		return
	}

}

// Login 登出
// @tags 用户
// @Summary 用户登出
// @Produce json
// @Accept json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/login [post]
func Logout(c *gin.Context) {
	var (
		err      error
		httpCode = http.StatusInternalServerError
		s        *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGin(c).Response(httpCode, err.Error())
		} else {
			app.NewGin(c).Response(200, map[string]string{})
		}
	}()
	if s, err = GetSession(c); err != nil {
		return
	}
	if err = srv_user.UserLogout(s.UID); err != nil {
		return
	}

}

// UserUpdate 用户更新基本信息
// @tags 用户
// @Summary 更新用户信息
// @Produce json
// @Param nickname body string false "昵称"
// @Param username body string false "用户名,之可以修改一次"
// @Param Phone body string false "手机号码"
// @Param sign body string false "签名"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/userUpdate [post]
func UserUpdate(c *gin.Context) {
	var (
		err      error
		form     = new(utils.FormEditUser)
		data     *models.UserBase
		httpCode = http.StatusInternalServerError
		s        *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGin(c).Response(httpCode, err.Error())
		} else {
			app.NewGin(c).Response(200, data)
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
	if data, err = srv_user.EditUser(s.UID, form); err != nil {
		return
	}
}

// GetUserProfile 获取用户信息
// @tags 用户
// @Summary 获取个人信息
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/getUserProfile [get]
func GetUserProfile(c *gin.Context) {
	var (
		err      error
		data     = new(models.UserBase)
		httpCode = http.StatusInternalServerError
		s        *session.Session
	)
	defer func() {
		if err != nil {
			app.NewGin(c).Response(httpCode, err.Error())
		} else {
			app.NewGin(c).Response(200, data)
		}
	}()
	if s, err = GetSession(c); err != nil {
		return
	}
	if err = data.GetUserByUId(s.UID); err != nil {
		return
	}

}

// ChangePayPasswd 修改密码
// @tags 用户
// @Summary 修改用户密码
// @Produce json
// @Accept json
// @Param oldPwd query string true "旧密码"
// @Param newPwd query string true "新密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/changeLoginPasswd [post]
func ChangeLoginPwd(c *gin.Context) {
	var (
		err            error
		httpCode       = http.StatusInternalServerError
		s              *session.Session
		oldPwd, newPwd string
	)
	defer func() {
		if err != nil {
			app.NewGin(c).Response(httpCode, err.Error())
		} else {
			app.NewGin(c).Response(200, "")
		}
	}()
	oldPwd = c.Param("oldPwd")
	newPwd = c.Param("newPwd")
	if s, err = GetSession(c); err != nil {
		return
	}
	if err = srv_user.ChangePwd(s.UID, oldPwd, newPwd); err != nil {
		return
	}
}

// ChangePayPasswd 修改密码
// @tags 用户
// @Summary 修改用户密码
// @Produce json
// @Accept json
// @Param oldPwd query string true "旧密码"
// @Param newPwd query string true "新密码"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/changeLoginPasswd [post]
func ChangePayPwd(c *gin.Context) {
	var (
		err            error
		httpCode       = http.StatusInternalServerError
		s              *session.Session
		oldPwd, newPwd string
	)
	defer func() {
		if err != nil {
			app.NewGin(c).Response(httpCode, err.Error())
		} else {
			app.NewGin(c).Response(200, "")
		}
	}()
	oldPwd = c.Param("oldPwd")
	newPwd = c.Param("newPwd")
	if s, err = GetSession(c); err != nil {
		return
	}
	if err = srv_user.ChangePayPwd(s.UID, oldPwd, newPwd); err != nil {
		return
	}
}

// EditHeadIcon 修改用户头像
// @tags 用户
// @Summary 修改用户头像
// @Produce json
// @Accept  json
// @Param headIcon query string true "头像"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/editHeadIcon [post]
func EditHeadIcon(c *gin.Context) {

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
			app.NewGin(c).Response(httpCode, err.Error())
		} else {
			app.NewGin(c).Response(200, gin.H{
				"headIcon": d.HeadIcon,
			})
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

	if err = d.UpdateUserOneColumn(s.UID, "head_icon", dst); err != nil {
		return
	}
	os.Remove(oldDst)

}

// GetHeadIcon 获取用户头像
// @tags 用户
// @Summary 获取用户头像
// @Produce json
// @Accept json
// @Success 200
// @Failure 500 {object} app.Response
// @Router /api/v1/getHeadIcon [get]
func GetHeadIcon(c *gin.Context) {
	var (
		err  error
		s    *session.Session
		data = new(models.UserBase)
	)
	defer func() {
		if err != nil {
			app.NewGin(c).Response(500, err.Error())
		}
	}()
	if s, err = GetSession(c); err != nil {
		return
	}
	if err = data.GetUserByUId(s.UID); err != nil {
		return
	}
	c.File(data.HeadIcon)
}
