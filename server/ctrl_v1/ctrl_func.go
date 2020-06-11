package ctrl
/*

import (
	"github.com/dchest/svc_captcha"
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/session"
	"github.com/olongfen/userDetail/models"
	"github.com/olongfen/userDetail/pkg/setting"
	"github.com/olongfen/userDetail/utils"
	"net/http"
	"strconv"
	"strings"
)



// Register 注册
func (c *ControlServe) Register(ctx *gin.Context) {
	if setting.ProjectSetting.IsCaptcha{
		digits := ctx.Query("digits")
		id := ctx.Query("id")
		if !verifyString(id,digits){
			ctx.AbortWithStatusJSON(200,gin.H{"msg":"verify svc_captcha failed"})
			return
		}
	}
	var (
		data *models.UserDetail
		err  error
	)
	defer PubCheckError(&err, ctx)
	var (
		f = &utils.FormRegister{}
	)
	if err = ctx.ShouldBind(f); err != nil {
		return
	}

	if data, err = models.PubUserAdd(f); err != nil {
		// 转换数据库错误
		err = utils.TransformGORMErr(err)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, data)
}

func verifyString(id,digits string) bool  {

	if !svc_captcha.VerifyString(id, digits) {
		return false
	}
	return true
}

//
func (c *ControlServe) Login(ctx *gin.Context) {
	if setting.ProjectSetting.IsCaptcha{
		digits := ctx.Query("digits")
		id := ctx.Query("id")
		if !verifyString(id,digits){
			ctx.AbortWithStatusJSON(200,gin.H{"msg":"verify svc_captcha failed"})
			return
		}
	}
	var (
		token string
		f     = new(utils.LoginForm)
		err   error
	)
	defer PubCheckError(&err, ctx)
	if err = ctx.ShouldBind(f); err != nil {
		return
	}
	f.IP = ctx.ClientIP()
	if token, err = models.Login(f); err != nil {
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"data": token,
	})
}

//  params: token isAdmin
func (c *ControlServe) Logout(ctx *gin.Context) {

	var (
		sn  *session.Session
		err error
	)
	defer PubCheckError(&err, ctx)
	if sn, err = models.TokenDecodeSession(ctx.Request, false); err != nil {
		if sn, err = models.TokenDecodeSession(ctx.Request, true); err != nil {
			return
		}
	}

	if err = models.Logout(sn.UID); err != nil {
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"data": "success",
	})
}

// GetUserDetailSelf 用户获取自己的信息
func (c *ControlServe) GetUserDetailSelf(ctx *gin.Context) {
	var (
		data *models.UserDetail
		sn   *session.Session
		err  error
	)
	defer PubCheckError(&err, ctx)
	if sn, err = models.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}

	if data, err = models.PubUserGet(sn.UID); err != nil {
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"data": data,
	})

}

// UpdateUserDetail 用户更新自己的信息
func (c *ControlServe) UpdateUserDetail(ctx *gin.Context) {
	var (
		err  error
		form = new(utils.UpdateUserProfile)
		data *models.UserDetail
		sn   *session.Session
	)
	defer PubCheckError(&err, ctx)
	if sn, err = models.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if err = ctx.ShouldBindJSON(form); err != nil {
		return
	}
	if data, err = models.PubUserUpdate(sn.UID, form); err != nil {
		return
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"data": data,
	})
}

// AddUserIDCard 实名验证
func (c *ControlServe) AddUserIDCard(ctx *gin.Context) {
	var (
		err  error
		sn   *session.Session
		form = new(utils.FormIDCard)
		data *models.IDCard
	)
	defer PubCheckError(&err, ctx)
	if sn, err = models.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}

	if err = ctx.ShouldBindJSON(form); err != nil {
		return
	}
	if data, err = models.PubIDCardAdd(sn.UID, form); err != nil {
		return
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"data": data,
	})
}



// GetUserIDCard 获取用户身份证信息
func (c *ControlServe) GetUserIDCard(ctx *gin.Context) {
	var (
		err  error
		sn   *session.Session
		data *models.IDCard
	)
	defer PubCheckError(&err, ctx)
	if sn, err = models.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if data, err = models.PubGetIDCard(sn.UID); err != nil {
		return
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"data": data,
	})
}

// AddUserBankCard 添加一张银行卡信息
func (c *ControlServe) AddUserBankCard(ctx *gin.Context) {
	var (
		err  error
		form = new(utils.FormBankCard)
		data *models.BankCard
		s    *session.Session
	)
	defer PubCheckError(&err, ctx)
	if err = ctx.ShouldBindJSON(form); err != nil {
		return
	}
	if s, err = models.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if data, err = models.PubBankCardAdd(s.UID, form); err != nil {
		return
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"data": data,
	})

}

// GetUserBankCardList 获取用户银行卡信息
func (c *ControlServe) GetUserBankCardList(ctx *gin.Context) {
	var (
		err error

		data []*models.BankCard
		s    *session.Session
	)
	defer PubCheckError(&err, ctx)

	if s, err = models.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if data, err = models.PubBankCardGetList(s.UID); err != nil {
		return
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"data": data,
	})

}

// DeleteUserBankCard 删除银行卡信息
func (c *ControlServe) DeleteUserBankCard(ctx *gin.Context) {
	var (
		err    error
		number string
		s      *session.Session
	)
	defer PubCheckError(&err, ctx)
	if s, err = models.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	number = ctx.Param("number")
	if err = models.PubBankCardDel(s.UID, number); err != nil {
		return
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"data": "success",
	})

}

// AddUserAddress 添加地址
func (c *ControlServe) AddUserAddress(ctx *gin.Context) {
	var (
		err  error
		s    *session.Session
		form = new(utils.FormAddr)
		data *models.AddressDetail
	)
	defer PubCheckError(&err, ctx)
	if err = ctx.ShouldBindJSON(form); err != nil {
		return
	}
	if s, err = models.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if data, err = models.PubAddressAdd(s.UID, form); err != nil {
		return
	}
	ctx.AbortWithStatusJSON(200, gin.H{
		"data": data,
	})

}

// GetUserAddressList 用户获取自己的地址
func (c *ControlServe) GetUserAddressList(ctx *gin.Context) {
	var (
		err  error
		data []*models.AddressDetail
		s    *session.Session
	)
	defer PubCheckError(&err, ctx)
	if s, err = models.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if data, err = models.PubAddressGetList(s.UID); err != nil {
		return
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"data": data,
	})
}

// UpdateUserAddress 用户更新自己的地址
func (c *ControlServe) UpdateUserAddress(ctx *gin.Context) {
	var (
		err  error
		s    *session.Session
		data *models.AddressDetail
		form *utils.FormAddr
	)
	defer PubCheckError(&err, ctx)
	form = new(utils.FormAddr)
	if err = ctx.ShouldBindJSON(form); err != nil {
		return
	}
	if s, err = models.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if data, err = models.PubAddressUpdate(s.UID, form); err != nil {
		return
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"data": data,
	})
}

// DeleteUserAddress 用户删除地址
func (c *ControlServe) DeleteUserAddress(ctx *gin.Context) {
	var (
		s   *session.Session
		id  string
		_id int64
		err error
	)
	defer PubCheckError(&err, ctx)
	id = ctx.Param("id")
	if _id, err = strconv.ParseInt(id, 10, 64); err != nil {
		return
	}

	if s, err = models.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if err = models.PubAddressDelete(s.UID, _id); err != nil {
		return
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"data": "success",
	})
}
*/