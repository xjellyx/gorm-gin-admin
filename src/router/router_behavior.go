package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/olongfen/gorm-gin-admin/src/controller/api/v1"
)

func apiBehavior(apiAdmin *gin.RouterGroup) {
	apiAdmin.DELETE("removeBehaviors", v1.RemoveBehaviors)
	apiAdmin.GET("getBehaviorList", v1.GetBehaviorList)
}
