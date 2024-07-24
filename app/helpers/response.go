package helpers

import "github.com/gin-gonic/gin"

type apiResponse struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type apiRes struct {
	Status  bool        `json:"status"`
	Code    string      `json:"code"`
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
	c.JSON(code, apiRes{
		Code:    "SUCCESS",
		Status:  true,
		Message: "success",
		Data:    data,
	})
}

func ApiResponseSuccess1(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, apiRes{
		Code:    "SUCCESS",
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func ApiResponseError(c *gin.Context, code int, err string, message string, data interface{}) {
	c.JSON(code, apiRes{
		Status:  false,
		Code:    err,
		Message: message,
		Data:    data,
	})
}
