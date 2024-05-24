package api

import (
	"net/http"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Регистрация
func (h *Handler) register(c *gin.Context) {
	var input securefilechanger.User

	if err := c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	// Create user
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		logrus.Infoln("[CreateUser]")
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// Структура для авторизации
type sighInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Авторизация
func (h *Handler) logIn(c *gin.Context) {
	var input sighInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorMessage(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
