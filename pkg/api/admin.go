package api

import (
	"net/http"
	"strconv"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/gin-gonic/gin"
)

type userList struct {
	Data []securefilechanger.User `json:"data"`
}

// Список сотрудников
func (h *Handler) userList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	users, err := h.services.User.GetAll(userId)
	if err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, userList{Data: users})
}

func (h *Handler) userListSearch(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	query := c.Query("q")
	if query == "" {
		newErrorMessage(c, http.StatusBadRequest, "no query params")
		return
	}

	users, err := h.services.User.GetLike(userId, query)
	if err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, userList{Data: users})
}

// Ограничение доступа Сотруднику
func (h *Handler) toggleUserApprove(c *gin.Context) {
	adminId, err := getUserId(c)
	if err != nil {
		return
	}

	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorMessage(c, http.StatusBadRequest, "invalid user id")
		return
	}

	if userId == adminId {
		newErrorMessage(c, http.StatusBadRequest, "admin connot disable yourself")
		return
	}

	err = h.services.User.ToggleApprove(userId)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{
		Status: "ok",
	})
}

// Проверка на существование Администратора
func (h *Handler) adminExist(c *gin.Context) {
	adminExist, err := h.services.Authorization.CheckAdminIsExists()

	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"admin_exist": adminExist})
}
