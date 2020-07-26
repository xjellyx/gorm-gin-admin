package v1

import (
	"bytes"
	"encoding/base64"
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
		Ext        string `json:"ext" form:"ext" `
		Lang       string `json:"lang" form:"lang" `
		IsDownload bool   `json:"isDownload" form:"isDownload"`
	}{}
	if err := ctx.BindQuery(&d); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	if len(d.Ext) == 0 {
		d.Ext = "png"
	}
	if len(d.Lang) == 0 {
		d.Lang = "en"
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
		//ctx.Header("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, captcha.StdWidth, captcha.StdHeight)
	case "wav":
		//ctx.Header("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, d.Lang)
	default:
		app.NewGinResponse(ctx).SetCodeAndMessage(500,captcha.ErrNotFound.Error()).Response()
	}

	if d.IsDownload {
		ctx.Header("Content-Type", "application/octet-stream")
	}
	data := make(map[string]interface{})
	data["id"] = id
	data["img"] = base64.StdEncoding.EncodeToString(content.Bytes())
	app.NewGinResponse(ctx).SetCodeAndMessage(200,"success").SetData(data).Response()
	// ctx.Data(200, ctx.GetHeader("Content-Type"), content.Bytes())
	//http.ServeContent(ctx.Writer, ctx.Request, id+"."+d.Ext, time.Time{}, bytes.NewReader(content.Bytes()))

}

// VerifyCaptcha
func VerifyCaptcha(c *gin.Context) {

	digits := c.Query("digits")
	id := c.Query("captchaId")
	verify := captcha.VerifyString(id, digits)
	app.NewGinResponse(c).SetCodeAndMessage(200,"success").SetData( gin.H{"verify": verify}).Response()

}
