package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/olongfen/gorm-gin-admin/src/controller/api/v1"
)

func roleRouterSettings(g *gin.RouterGroup) {
	g.GET("getSettings", v1.GetSettings)
}
