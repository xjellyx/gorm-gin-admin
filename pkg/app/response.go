package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Error  interface{} `json:"error"`
}

// NewGin
func NewGin(c *gin.Context) *Gin {
	return &Gin{
		c,
	}
}

// Response setting gin.JSON
func (g *Gin) Response(httpCode int, data interface{}) {
	if httpCode != http.StatusOK {
		g.C.JSON(httpCode, Response{
			Meta: Meta{
				Status: httpCode,
				Msg:    "fail",
				Error:  data,
			},
		})
	} else {
		g.C.JSON(httpCode, Response{
			Meta: Meta{
				Status: httpCode,
				Msg:    "success",
			},
			Data: data,
		})
	}
	g.C.Abort()
	return
}
