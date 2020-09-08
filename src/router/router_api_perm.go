package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/olongfen/gorm-gin-admin/src/controller/api/v1"
)

func apiRolePermRouterAPI(apiAdmin *gin.RouterGroup)  {
	apiAdmin.POST("addRoleApiPerm", v1.AddRoleAPIPerm)
	apiAdmin.DELETE("removeRoleApiPerm", v1.RemoveRolePermAPI)
	apiAdmin.GET("getRoleApiList", v1.GetRoleApiList)
}
