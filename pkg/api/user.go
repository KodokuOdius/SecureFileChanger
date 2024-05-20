package api

import (
	"net/http"
	"strconv"

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

	c.Redirect(http.StatusOK, "/sing-up")
}

func (h *Handler) disableUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorMessage(c, http.StatusBadRequest, "Invalid user id")
		return
	}

	err = h.services.User.SetDisable(id)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{
		Status: "ok",
	})
}

func (h *Handler) newPassword(c *gin.Context) {
	// NewPassword(userId int, password string) error\
}
