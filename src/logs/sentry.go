package logs

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func SentryCaptureAndSendException(ctx *gin.Context, err error) {
	if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
		hub.CaptureException(err)
	}
}
