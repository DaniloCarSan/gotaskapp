package helpers

import "github.com/gin-gonic/gin"

type apiResponse struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type apiResponseError struct {
	Error   string      `json:"error"`
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

func ApiResponseSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(code, data)
}

func ApiResponseError(c *gin.Context, code int, error string, message string, data interface{}) {
	c.JSON(code, apiResponseError{
		Error:   error,
		Message: message,
		Data:    data,
	})
}
