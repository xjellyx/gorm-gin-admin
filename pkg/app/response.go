package app

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C    *gin.Context
	resp *Response
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
}

// NewGinResponse
func NewGinResponse(c *gin.Context) *Gin {
	return &Gin{
		c,
		&Response{},
	}
}


func (g *Gin) SetCodeAndMessage(code int,message string)*Gin  {
	g.resp.Meta.Code=code
	g.resp.Meta.Message=message
	return g
}

func (g *Gin)SetData(data interface{})*Gin  {
	g.resp.Data=data
	return g
}

// Response setting gin.JSON
func (g *Gin) Response() {
	g.C.JSON(g.resp.Meta.Code, g.resp)
	g.C.Abort()
	return
}
