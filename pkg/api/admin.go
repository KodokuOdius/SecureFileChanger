package api

import (
	"net/http"
	"strconv"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type userList struct {
	Data []securefilechanger.User `json:"data"`
}

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

func (h *Handler) toggleUserApprove(c *gin.Context) {
	// only admin

	logrus.Info("[toggleUserApprove]")
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

func (h *Handler) adminExist(c *gin.Context) {
	adminExist, err := h.services.Authorization.CheckAdminIsExists()

	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"admin_exist": adminExist})
}
