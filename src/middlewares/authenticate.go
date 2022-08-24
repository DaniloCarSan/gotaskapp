package middlewares

import (
	"gotaskapp/src/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Verify user authenticated
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		if err := security.ValidateJwtToken(c.Request); err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}

		c.Next()
	}
}
