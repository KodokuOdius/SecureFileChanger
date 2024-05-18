package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

// middlware для проверка токена авторизации
func (h *Handler) userIdentity(c *gin.Context) {
	// Проверка на наличие заголовков
	header := c.GetHeader(authHeader)
	if header == "" {
		newErrorMessage(c, http.StatusUnauthorized, "Empty header")
		return
	}

	// Проверка на корректность заголовков
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorMessage(c, http.StatusUnauthorized, "Invalid auth header")
		return
	}
	token := headerParts[1]

	// parse token
	userId, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		newErrorMessage(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

// Получения id пользователя из токена
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorMessage(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorMessage(c, http.StatusInternalServerError, "invalid user id")
		return 0, errors.New("invalid user id")
	}

	return idInt, nil
}
