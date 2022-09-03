package auth

import (
	"gotaskapp/app/entities"
	fail "gotaskapp/app/failures"
	"gotaskapp/app/helpers"
	"gotaskapp/app/repositories/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signin struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6,max=16"`
}

// Sign in
func SignIn(c *gin.Context) {

	var form signin

	if err := c.ShouldBind(&form); err != nil {
		helpers.ApiResponse(c, false, http.StatusBadRequest, "invalid form fields", err.Error())
		return
	}

	credential, err := auth.SigIn(
		entities.Auth{
			Email:    form.Email,
			Password: form.Password,
		},
	)

	if err != nil {
		switch err.(type) {
		case *fail.DatabaseConnectFailure:
		case *fail.SqlSelectFailure:
		case *fail.GenerateJwtTokenFailure:
			helpers.ApiResponse(c, false, http.StatusInternalServerError, "Server internal error", nil)
			return
		case *fail.SqlSelectNotFoundFailure:
		case *fail.SignInFailure:
			helpers.ApiResponse(c, false, http.StatusUnauthorized, "Email or password invalid", nil)
			return
		}
	}

	helpers.ApiResponse(c, true, http.StatusOK, "success", credential)
}
