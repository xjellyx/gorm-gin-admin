package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/olongfen/gorm-gin-admin/src/controller/api/v1"
)

func apiRouterAPI(apiAdmin *gin.RouterGroup) {
	apiAdmin.GET("getApiGroupList", v1.GetAPIGroupList)
	apiAdmin.GET("getApiGroupListAll", v1.GetAPIGroupListAll)
	apiAdmin.POST("addApiGroup", v1.AddApiGroup)
	apiAdmin.DELETE("removeApiGroup", v1.RemoveApiGroup)
	apiAdmin.PUT("editApiGroup", v1.EditApiGroup)
}
