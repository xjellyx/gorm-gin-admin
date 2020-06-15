package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Msg   string      `json:"msg"`
	Error interface{} `json:"error"`
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
		g.C.JSON(httpCode, ResponseError{
			Msg:   "fail",
			Error: data,
		})
	} else {
		g.C.JSON(httpCode, Response{
			Msg:  "success",
			Data: data,
		})
	}
	g.C.Abort()
	return
}
