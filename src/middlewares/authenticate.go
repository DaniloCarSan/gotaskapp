package middlewares

import (
	"gotaskapp/src/security"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Verify user authenticated
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {

		var tokenString string
		var err error

		if tokenString, err = security.ExtractJwtTokenFromHeaderAuthorization(c.Request); err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}

		if _, err := security.ValidateJwtToken(tokenString); err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}

		c.Next()
	}
}
