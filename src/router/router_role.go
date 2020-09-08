package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/olongfen/gorm-gin-admin/src/controller/api/v1"
)

func roleRouterAPI(g *gin.RouterGroup)  {
	g.GET("getRoleList",v1.GetRoleList)
	g.GET("getRoleLevel",v1.GetRoleLevel)
	g.PUT("editRole",v1.EditRole)
	g.POST("addRole",v1.AddRole)
	g.DELETE("removeRole",v1.RemoveRole)
}
