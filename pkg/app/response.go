package app

import (
	"github.com/gin-gonic/gin"
	"github.com/olongfen/user_base/pkg/codes"
)

type Gin struct {
	c      *gin.Context
	resp   *Response
	status int
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewGinResponse
func NewGinResponse(c *gin.Context) *Gin {
	return &Gin{
		c,
		&Response{},
		200,
	}
}

func (g *Gin) Fail(code int, message string) *Gin {
	g.resp.Meta.Code = code
	g.resp.Meta.Message = message
	return g
}

func (g *Gin) SetStatus(status int) *Gin {
	g.status = status
	return g
}

func (g *Gin) Success(data interface{}) *Gin {
	g.resp.Meta.Code = codes.CodeSuccess
	g.resp.Meta.Message = "success"
	g.resp.Data = data
	return g
}

// Response setting gin.JSON
func (g *Gin) Response() {
	g.c.JSON(g.status, g.resp)
	g.c.Abort()
	return
}
