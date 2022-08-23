package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Sign up
func SignUp(c *gin.Context, data map[string]interface{}) {
	c.JSON(http.StatusOK, data)
}
