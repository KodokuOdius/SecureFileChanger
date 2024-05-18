package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message"`
}

func newErrorMessage(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	// Блокарует выполнение обработчиков
	c.AbortWithStatusJSON(statusCode, Error{message})
}
