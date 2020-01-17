package ctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PubCheckError
func PubCheckError(err *error, c *gin.Context) {
	if err != nil && *err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": c.Error(*err).Error(),
		})
	}
}
