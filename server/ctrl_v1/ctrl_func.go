package ctrl

import (
	"github.com/gin-gonic/gin"
	userBase "github.com/srlemon/userDetail"
	"github.com/srlemon/userDetail/model"
	"net/http"
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

func (s *ControlServe) CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", allowHeaders)
		c.Writer.Header().Set("Access-Control-Allow-Methods", allowMethods)

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (s *ControlServe) Common(c *gin.Context) {
	header := c.Writer.Header()
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
func (s *ControlServe) Register(c *gin.Context) {
	var (
		data *model.UserDetail
		err  error
	)
	defer PubCheckError(&err, c)
	var (
		f = &userBase.FormRegister{}
	)
	if err = c.ShouldBind(f); err != nil {
		return
	}

	if data, err = model.PubUserAdd(f); err != nil {
		// 转换数据库错误
		err = userBase.TransformGORMErr(err)
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, data)
}

//
func (s *ControlServe) Login(c *gin.Context) {
	var (
		data  *model.UserDetail
		token string
		f     = new(userBase.LoginForm)
		err   error
	)
	defer PubCheckError(&err, c)
	if err = c.ShouldBind(f); err != nil {
		return
	}
	if data, token, err = model.Login(f); err != nil {
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"data":  data,
		"token": token,
	})
}
