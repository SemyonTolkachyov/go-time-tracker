package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}

func (e ErrorResponse) String() string {
	return fmt.Sprintf("{message: %s}", e.Message)
}

func (e StatusResponse) String() string {
	return fmt.Sprintf("{status: %s}", e.Status)
}
