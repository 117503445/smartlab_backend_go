package middleware

import (
	"net/http"
	"smartlab/model"
	"smartlab/serializer"
	"smartlab/service"

	"github.com/gin-gonic/gin"
)

func HasRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := service.CurrentUser(c)
		if user == nil {
			c.JSON(http.StatusProxyAuthRequired, serializer.Err(http.StatusProxyAuthRequired, "has role failed: need auth", nil))
			c.Abort()
			return
		}
		model.DB.Preload("Roles").Find(&user)
		roles := user.Roles

		isFound := false
		for _, r := range roles {
			if r.Name == role {
				isFound = true
				c.Next()
			}
		}
		if !isFound {
			c.JSON(http.StatusProxyAuthRequired, serializer.Err(http.StatusProxyAuthRequired, "has role failed: don't have role "+role, nil))
			c.Abort()
		}
	}
}
