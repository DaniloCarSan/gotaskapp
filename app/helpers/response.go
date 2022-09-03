package helpers

import "github.com/gin-gonic/gin"

type apiResponse struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ApiResponse(c *gin.Context, Status bool, code int, message string, data interface{}) {
	c.JSON(code, apiResponse{
		Status:  Status,
		Code:    code,
		Message: message,
		Data:    data,
	})
}
