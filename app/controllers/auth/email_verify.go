package auth

import (
	"gotaskapp/app/database"
	"gotaskapp/app/entities"
	"gotaskapp/app/helpers"
	"gotaskapp/app/security"
	"net/http"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

// Email verify
func EmailVerify(c *gin.Context) {

	tokenString := c.Param("token")

	var token, err = security.ValidateJwtToken(tokenString)
	var id uint64
	var user entities.User

	if err != nil {
		helpers.ApiResponseError(c, http.StatusUnauthorized, "INVALID_TOKEN", "Error, invalid confirmation token or it has expired, request another", nil)
		return
	}

	id, err = security.ExtractUserIdOfJwtToken(token)

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		helpers.ApiResponseError(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Server internal error", nil)
		return
	}

	repository, err := database.Repository()

	if err != nil {
		if hub := sentrygin.GetHubFromContext(c); hub != nil {
			hub.CaptureException(err)
		}
		helpers.ApiResponseError(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Server internal error", nil)
		return
	}

	if user, err = repository.User.ById(id); err != nil {
		helpers.ApiResponseError(c, http.StatusUnauthorized, "INVALID_TOKEN", "Error, invalid confirmation token or it has expired, request another", nil)
		return
	}

	if !user.IsEmailVerified() {

		if err := repository.User.SetEmailVerified(id); err != nil {
			if hub := sentrygin.GetHubFromContext(c); hub != nil {
				hub.CaptureException(err)
			}
			helpers.ApiResponseError(c, http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Server internal error", nil)
			return
		}
	}

	helpers.ApiResponseSuccess1(c, http.StatusOK, "Email successfully confirmed, enjoy the app.", nil)
}
