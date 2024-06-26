package api

import (
	"net/http"
	"strconv"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/gin-gonic/gin"
)

// Обработчик для операция с директориями
// Создание директории
func (h *Handler) createFolder(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input securefilechanger.Folder
	if err := c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	// Вызов метода сервиса
	id, err := h.services.Folder.Create(userId, input)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":          id,
		"folder_name": input.Name,
		"is_root":     false,
	})
}

// Структура для списка директорий
type getAllFolders struct {
	Data []securefilechanger.Folder `json:"data"`
}

// Список всех директорий
func (h *Handler) getAllFolders(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	folders, err := h.services.Folder.GetAll(userId)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllFolders{Data: folders})
}

// Изменение имени директории
func (h *Handler) updateFolder(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	folderId, err := strconv.Atoi(c.Param("folder_id"))
	if err != nil {
		newErrorMessage(c, http.StatusBadRequest, "invalid folder id")
		return
	}

	var input securefilechanger.UpdateFolder
	if err := c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.services.Folder.GetById(userId, folderId)
	if err != nil {
		newErrorMessage(c, http.StatusNotFound, err.Error())
		return
	}

	if err := h.services.Folder.Update(userId, folderId, input); err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{Status: "ok"})
}

// Удаление директории
func (h *Handler) deleteFolder(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	folderId, err := strconv.Atoi(c.Param("folder_id"))
	if err != nil {
		newErrorMessage(c, http.StatusBadRequest, "invalid folder id")
		return
	}

	_, err = h.services.Folder.GetById(userId, folderId)
	if err != nil {
		newErrorMessage(c, http.StatusNotFound, err.Error())
		return
	}

	err = h.services.Folder.Delete(userId, folderId)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{Status: "ok"})
}
