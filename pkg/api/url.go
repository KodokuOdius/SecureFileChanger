package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Обрабочик для операция с временными ссылками

type filesIds struct {
	Ids      []int `json:"file_id" binding:"required"`
	HourLive int   `json:"hour_live"`
}

// Создание временной ссылки
func (h *Handler) createUrl(c *gin.Context) {
	logrus.Info("[Handler createUrl]")
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input filesIds
	err = c.BindJSON(&input)
	if err != nil {
		logrus.Info("[BindJSON]")
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	hex, err := h.services.Url.CreateUrl(userId, input.HourLive, input.Ids)
	if err != nil {
		logrus.Info("[Url.CreateUrl]")
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"url_hex": hex,
	})
}

// Список документов по временной ссылке
func (h *Handler) getFilesUUid(c *gin.Context) {
	urlUUid := c.Param("uuid")

	if len(urlUUid) != 36 {
		newErrorMessage(c, http.StatusBadRequest, "invalid uuid url")
		return
	}

	urlFile, err := h.services.Url.GetUrl(urlUUid)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	if urlFile.Id == 0 {
		newErrorMessage(c, http.StatusNotFound, "url uuid not found")
		return
	}

	expireUrlDt := urlFile.CreateDt.Add(time.Hour * time.Duration(urlFile.HourLive))

	if expireUrlDt.Before(time.Now()) && urlFile.HourLive > 0 {
		newErrorMessage(c, http.StatusNotFound, "url uuid not found")
		err := h.services.Url.DeleteUrl(urlUUid)
		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		return
	}

	files, err := h.services.Url.GetFilesList(urlUUid)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, filesList{Data: files})
}

// Выгрзука документов по временной ссылке
func (h *Handler) downloadFilesUUid(c *gin.Context) {
	urlUUid := c.Param("uuid")

	if len(urlUUid) != 36 {
		newErrorMessage(c, http.StatusBadRequest, "invalid uuid url")
		return
	}

	urlFile, err := h.services.Url.GetUrl(urlUUid)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	if urlFile.Id == 0 {
		newErrorMessage(c, http.StatusNotFound, "url uuid not found")
		return
	}

	expireUrlDt := urlFile.CreateDt.Add(time.Hour * time.Duration(urlFile.HourLive))

	if expireUrlDt.Before(time.Now()) && urlFile.HourLive > 0 {
		newErrorMessage(c, http.StatusNotFound, "url uuid not found")
		err := h.services.Url.DeleteUrl(urlUUid)
		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		return
	}

	files, err := h.services.Url.GetFilesList(urlUUid)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	if urlFile.HourLive == 0 {
		// one time url
		err := h.services.Url.DeleteUrl(urlUUid)
		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if len(files) == 1 {
		h.DownloadOneFile(c, files[0])
	} else {
		err = h.DownloadFileZip(c, files)
		if err != nil {
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
}
