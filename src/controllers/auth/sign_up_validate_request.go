package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password," json:"password" binding:"required,min=2"`
}

func SignUpValidateRequest(c *gin.Context) {

	// data := map[string]interface{}{}

	var form User
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	errors:{
		password: {
			"required":"Fields required",
			"min":"require min 6 characteres"
		}
	}

	// SignUp(c, data)
}
