package helper

import "github.com/gin-gonic/gin"

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Errors     interface{} `json:"errors,omitempty"`
}

func NewResponse(statusCode int, message string, data interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}

func NewErrorResponse(statusCode int, message string, errors interface{}) ErrorResponse {
	return ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
		Errors:     errors,
	}
}

func Success(c *gin.Context, statusCode int, msg string, data ...interface{}) {
	resp := NewResponse(statusCode, msg, data)
	c.JSON(statusCode, resp)
}

func Error(c *gin.Context, statusCode int, msg string, data ...interface{}) {
	resp := NewErrorResponse(statusCode, msg, data)
	c.AbortWithStatusJSON(statusCode, resp)
}
