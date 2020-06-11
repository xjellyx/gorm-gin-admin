package resp

import "github.com/gin-gonic/gin"

// Gin
type Gin struct {
	C *gin.Context
}

// Response
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON todo
func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  "",
		Data: data,
	})
	return
}
