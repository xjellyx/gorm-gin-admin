package api_v1

import "github.com/gin-gonic/gin"

// UserRegister 用户注册
// @Summary 用户注册
// @Produce json
// @Param phone string true 'phone'
// @param password true 'phone'
// @param code   false 'code'
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /api/v1/register
func UserRegister(c *gin.Context) {

}
