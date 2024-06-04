package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		newErrorMessage(c, http.StatusUnauthorized, "empty header")
		return
	}

	// Проверка на корректность заголовков
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorMessage(c, http.StatusUnauthorized, "invalid auth header")
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
	c.Next()
}

func (h *Handler) CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusAccepted)
		return
	}

	c.Next()
}

// middlware для логгирования запросов
func (h *Handler) logMiddleware(c *gin.Context) {
	method := c.Request.Method
	url := c.Request.URL

	logrus.Infoln(fmt.Sprintf("[%s] %s", method, url))

	c.Next()
}

// middlware для проверки доступа Сотрудника
func (h *Handler) userCheckApprove(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	isApproved, err := h.services.User.IsApproved(userId)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !isApproved {
		newErrorMessage(c, http.StatusForbidden, "user not approved")
		return
	}

	c.Next()
}

// middlware для проверки доступа к Админ панели
func (h *Handler) adminIdentify(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	isAdmin, err := h.services.User.IsAdmin(userId)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !isAdmin {
		newErrorMessage(c, http.StatusForbidden, "user not admin")
		return
	}

	c.Next()
}

// middlware ограничение на объём файла
func (h *Handler) maxFileLimit(c *gin.Context) {
	// 15 Mb
	maxBytes := 1024 * 1024 * 15
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(maxBytes))

	c.Next()
}

// Получения id пользователя из токена
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorMessage(c, http.StatusUnauthorized, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorMessage(c, http.StatusInternalServerError, "invalid user id")
		return 0, errors.New("invalid user id")
	}

	return idInt, nil
}
