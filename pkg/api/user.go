package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/gin-gonic/gin"
)

// Обрабочик для операция с Сотрудниками
func (h *Handler) updateUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input securefilechanger.UpdateUser
	if err = c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = input.Validate(); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.User.Update(userId, input)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{Status: "ok"})
}

func (h *Handler) infoUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	user, err := h.services.User.GetInfo(userId)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
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
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input securefilechanger.ChangePass
	if err = c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.User.NewPassword(userId, input)
	if err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{Status: "ok"})
}
