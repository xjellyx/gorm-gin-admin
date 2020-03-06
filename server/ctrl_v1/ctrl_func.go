package ctrl

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/contrib/session"
	userBase "github.com/olongfen/userDetail"
	"github.com/olongfen/userDetail/model"
	"net/http"
	"strconv"
	"strings"
)

const (
	PayPasswordHeader = "X-Pay-Password"
	SignatureHeader   = "X-Signature"
	Token             = "X-Token"
)

var (
	allowHeaders = strings.Join([]string{
		"accept",
		"origin",
		"Authorization",
		"Content-Type",
		"Content-Length",
		"Content-Length",
		"Accept-Encoding",
		"Cache-Control",
		"X-CSRF-Token",
		"X-Requested-With",
		Token,
		SignatureHeader,    // 接受签名的 Header
		PayPasswordHeader,  // 接收交易密码的 Header
		"X-Wechat-Binding", // 激活微信帐号
	}, ",")
	allowMethods = strings.Join([]string{
		http.MethodOptions,
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
	}, ",")
)

func (c *ControlServe) CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		origin := ctx.GetHeader("Origin")
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", allowHeaders)
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", allowMethods)

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}

func (c *ControlServe) Common(ctx *gin.Context) {
	header := ctx.Writer.Header()
	// alone dns prefect
	header.Set("X-DNS-Prefetch-Control", "on")
	// IE No Open
	header.Set("X-Download-Options", "noopen")
	// not cache
	header.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
	header.Set("Expires", "max-age=0")
	// Content Security Policy
	header.Set("Content-Security-Policy", "default-src 'self'")
	// xss protect
	// it will caught some problems is old IE
	header.Set("X-XSS-Protection", "1; mode=block")
	// Referrer Policy
	header.Set("Referrer-Header", "no-referrer")
	// cros frame, allow same origin
	header.Set("X-Frame-Options", "SAMEORIGIN")
	// HSTS
	header.Set("Strict-Transport-Security", "max-age=5184000;includeSubDomains")
	// no sniff
	header.Set("X-Content-Type-Options", "nosniff")
}

// Register 注册
func (c *ControlServe) Register(ctx *gin.Context) {
	var (
		data *model.UserDetail
		err  error
	)
	defer PubCheckError(&err, ctx)
	var (
		f = &userBase.FormRegister{}
	)
	if err = ctx.ShouldBind(f); err != nil {
		return
	}

	if data, err = model.PubUserAdd(f); err != nil {
		// 转换数据库错误
		err = userBase.TransformGORMErr(err)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, data)
}

//
func (c *ControlServe) Login(ctx *gin.Context) {
	var (
		data  *model.UserDetail
		token string
		f     = new(userBase.LoginForm)
		err   error
	)
	defer PubCheckError(&err, ctx)
	if err = ctx.ShouldBind(f); err != nil {
		return
	}
	if data, token, err = model.Login(f); err != nil {
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"data":  data,
		"token": token,
	})
}

// GetUserDetailSelf 用户获取自己的信息
func (c *ControlServe) GetUserDetailSelf(ctx *gin.Context) {
	var (
		data *model.UserDetail
		sn   *session.Session
		err  error
	)
	defer PubCheckError(&err, ctx)
	if sn, err = model.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}

	if data, err = model.PubUserGet(sn.UID); err != nil {
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
		form = new(userBase.UpdateUserProfile)
		data *model.UserDetail
		sn   *session.Session
	)
	defer PubCheckError(&err, ctx)
	if sn, err = model.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if err = ctx.ShouldBindJSON(form); err != nil {
		return
	}
	if data, err = model.PubUserUpdate(sn.UID, form); err != nil {
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
		form = new(userBase.FormIDCard)
		data *model.IDCard
	)
	defer PubCheckError(&err, ctx)
	if sn, err = model.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}

	if err = ctx.ShouldBindJSON(form); err != nil {
		return
	}
	if data, err = model.PubIDCardAdd(sn.UID, form); err != nil {
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
		data *model.IDCard
	)
	defer PubCheckError(&err, ctx)
	if sn, err = model.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if data, err = model.PubGetIDCard(sn.UID); err != nil {
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
		form = new(userBase.FormBankCard)
		data *model.BankCard
		s    *session.Session
	)
	defer PubCheckError(&err, ctx)
	if err = ctx.ShouldBindJSON(form); err != nil {
		return
	}
	if s, err = model.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if data, err = model.PubBankCardAdd(s.UID, form); err != nil {
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

		data []*model.BankCard
		s    *session.Session
	)
	defer PubCheckError(&err, ctx)

	if s, err = model.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if data, err = model.PubBankCardGetList(s.UID); err != nil {
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
	if s, err = model.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	number = ctx.Param("number")
	if err = model.PubBankCardDel(s.UID, number); err != nil {
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
		form = new(userBase.FormAddr)
		data *model.AddressDetail
	)
	defer PubCheckError(&err, ctx)
	if err = ctx.ShouldBindJSON(form); err != nil {
		return
	}
	if s, err = model.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if data, err = model.PubAddressAdd(s.UID, form); err != nil {
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
		data []*model.AddressDetail
		s    *session.Session
	)
	defer PubCheckError(&err, ctx)
	if s, err = model.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if data, err = model.PubAddressGetList(s.UID); err != nil {
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
		data *model.AddressDetail
		form *userBase.FormAddr
	)
	defer PubCheckError(&err, ctx)
	form = new(userBase.FormAddr)
	if err = ctx.ShouldBindJSON(form); err != nil {
		return
	}
	if s, err = model.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if data, err = model.PubAddressUpdate(s.UID, form); err != nil {
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

	if s, err = model.TokenDecodeSession(ctx.Request, false); err != nil {
		return
	}
	if err = model.PubAddressDelete(s.UID, _id); err != nil {
		return
	}

	ctx.AbortWithStatusJSON(200, gin.H{
		"data": "success",
	})
}
