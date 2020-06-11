package api_v1

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/olongfen/userDetail/service/svc_captcha"
	"net/http"
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
	 svc_captcha.Serve(ctx.Writer, ctx.Request, id, d.Ext, d.Lang, d.IsDownload, captcha.StdWidth, captcha.StdHeight)
}



