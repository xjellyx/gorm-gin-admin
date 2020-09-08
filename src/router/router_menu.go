package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/olongfen/gorm-gin-admin/src/controller/api/v1"
)

func menuRouterAPI(apiAdmin *gin.RouterGroup)  {
	apiAdmin.POST("addMenu", v1.AddMenu)
	apiAdmin.GET("getMenu", v1.GetMenu)
	apiAdmin.GET("getMenuList", v1.GetMenuList)
	apiAdmin.DELETE("delMenu", v1.DelMenu)
	apiAdmin.PUT("editMenu", v1.EditMenu)
}
