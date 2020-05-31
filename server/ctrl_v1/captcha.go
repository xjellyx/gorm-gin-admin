package ctrl

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (c *ControlServe)Captcha(ctx *gin.Context) {
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
	 Serve(ctx.Writer, ctx.Request, id, d.Ext, d.Lang, d.IsDownload, captcha.StdWidth, captcha.StdHeight)
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Captcha-ID", id)
	var content bytes.Buffer
	switch ext {
	case "png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case "wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

