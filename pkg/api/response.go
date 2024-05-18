package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Структура ответов сервера
type errorResponce struct {
	Message string `json:"message"`
}

type statusResponce struct {
	Status string `json:"status"`
}

func newErrorMessage(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	// Блокарует выполнение обработчиков
	c.AbortWithStatusJSON(statusCode, errorResponce{message})
}
