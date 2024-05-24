package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Обрабочик для операция с Сотрудниками
func (h *Handler) updateUser(c *gin.Context) {
	// Update(userId int, input securefilechanger.UpdateUser) error
}

func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	err = h.services.User.Delete(userId)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusMovedPermanently, map[string]interface{}{"redirect": "/api/auth/register"})

	path := filepath.Join(".", fmt.Sprintf("files/user%d", userId))

	os.RemoveAll(path)
}

func (h *Handler) newPassword(c *gin.Context) {
	// NewPassword(userId int, password string) error\
}
