package api

import (
	"net/http"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/gin-gonic/gin"
)

// Регистрация
func (h *Handler) signUp(c *gin.Context) {
	var input securefilechanger.User

	if err := c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	// Create user
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Create default bin
	err = h.services.Folder.CreateDefaultFolder(id)
	if err != nil {
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
func (h *Handler) signIn(c *gin.Context) {
	var input sighInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
