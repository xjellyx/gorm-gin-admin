package cors

import (
	"github.com/gin-gonic/gin"
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

// CORS
func  CORS() gin.HandlerFunc {
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

// Common
func Common() gin.HandlerFunc {
	return func(context *gin.Context) {
		header := context.Writer.Header()
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
}

