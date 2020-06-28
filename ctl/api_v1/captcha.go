package api_v1

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/olongfen/user_base/pkg/app"
	"net/http"
	"strconv"
	"time"
)

// Captcha
func Captcha(ctx *gin.Context) {
	var d = struct {
		Ext        string `json:"ext" form:"ext" binding:"required"`
		Lang       string `json:"lang" form:"lang" binding:"required"`
		IsDownload bool   `json:"isDownload" form:"isDownload"`
	}{}
	if err := ctx.BindQuery(&d); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	l := captcha.DefaultLen
	id := captcha.NewLen(l)
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", strconv.Itoa(int(captcha.Expiration/time.Second))+"s")
	ctx.Header("Captcha-ID", id)
	var content bytes.Buffer
	switch d.Ext {
	case "png":
		ctx.Header("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, captcha.StdWidth, captcha.StdHeight)
	case "wav":
		ctx.Header("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, d.Lang)
	default:

		ctx.JSON(500, app.Response{
			Meta: app.Meta{
				Status: 500,
				Msg:    "fail",
				Error:  captcha.ErrNotFound,
			},
			Data: nil,
		})
	}

	if d.IsDownload {
		ctx.Header("Content-Type", "application/octet-stream")
	}
	ctx.Data(200, ctx.GetHeader("Content-Type"), content.Bytes())

	//http.ServeContent(ctx.Writer, ctx.Request, id+"."+d.Ext, time.Time{}, bytes.NewReader(content.Bytes()))

}

// VerifyCaptcha
func VerifyCaptcha(c *gin.Context) {

	digits := c.Query("digits")
	id := c.Query("captchaId")
	verify := captcha.VerifyString(id, digits)
	if verify {
		app.NewGin(c).Response(200, gin.H{"verify": verify})
	} else {
		app.NewGin(c).Response(200, gin.H{"verify": verify})
	}

}
