package auth

import (
	"gotaskapp/app/database"
	"gotaskapp/app/entities"
	"gotaskapp/app/helpers"
	"gotaskapp/app/security"
	"net/http"
	"strings"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

type passwordReset struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6,max=16"`
}

// Password reset
func PasswordReset(c *gin.Context) {

	var form passwordReset

	if err := c.ShouldBind(&form); err != nil {
		helpers.ApiResponseError(c, http.StatusBadRequest, "FORM_FIELDS_INVALID", err.Error(), nil)
		return
	}

	tokenString := c.Param("token")

	var token, err = security.ValidateJwtToken(tokenString)
	var id uint64

	if err != nil {
		helpers.ApiResponseError(c, http.StatusBadRequest, "INVALID_TOKEN", "Error, invalid confirmation token or it has expired, request another", nil)
		return
	}

	id, err = security.ExtractUserIdOfJwtToken(token)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		helpers.ApiResponseError(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error(), nil)
		return
	}

	repository, err := database.Repository()

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		helpers.ApiResponseError(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", err.Error(), nil)
		return
	}

	var user entities.User

	if user, err = repository.User.ById(id); err != nil {
		helpers.ApiResponseError(c, http.StatusUnauthorized, "USER_NOT_FOUND", "error, invalid confirmation token or it has expired, request another", nil)
		return
	}

	if !strings.EqualFold(user.Email, form.Email) {
		helpers.ApiResponseError(c, http.StatusUnauthorized, "EMAIL_NOT_FOUND", "the email provided does not correspond to this password reset request.", nil)
		return
	}

	user.Password = form.Password

	err = user.PasswordToHash()

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		helpers.ApiResponseError(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "an internal error occurred, please try again later.", nil)
		return
	}

	err = repository.User.PasswordReset(user.ID, user.Password)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		helpers.ApiResponseError(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "an internal error occurred, please try again later.", nil)
		return
	}

	helpers.ApiResponseSuccess1(c, http.StatusOK, "Password successfully changed, sign in now and enjoy our features.", nil)
}
