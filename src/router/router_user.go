package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/olongfen/gorm-gin-admin/src/controller/api/v1"
)

func userBaseRouterAPI(g *gin.RouterGroup)  {
	g.POST("logout", v1.UserLogout)
	g.POST("verified", v1.Verified)
	g.POST("setPayPwd", v1.SetPayPwd)
	g.PUT("modifyProfile", v1.ModifyProfile)
	g.PUT("modifyLoginPwd", v1.ModifyLoginPwd)
	g.PUT("modifyPayPwd", v1.ModifyPayPwd)
	g.PUT("modifyHeadIcon", v1.ModifyHeadIcon)
	g.GET("getHeadIcon", v1.GetHeadIcon)
	g.GET("profile", v1.GetUserProfile)
}

func adminActionUserRouterAPI(apiAdmin *gin.RouterGroup)  {
	// apiAdmin.POST("/editUser", v1.EditUser)
	apiAdmin.GET("listUser", v1.ListUser)
	apiAdmin.PUT("editUser", v1.EditUser)
	apiAdmin.DELETE("deleteUser", v1.DeleteUser)
	apiAdmin.GET("getUserKV",v1.GetUserKV)
	apiAdmin.GET("userTotal", v1.UserTotal)
}