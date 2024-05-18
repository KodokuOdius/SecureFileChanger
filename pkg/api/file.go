package api

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) uploadFile(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	name := c.Request.Form.Get("name")
	// Проверка наличия имени файла
	if name == "" {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	f, handler, err := c.Request.FormFile("file")
	// Проверка наличия самого файла
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer f.Close()

	fileExtension := strings.ToLower(filepath.Ext(handler.Filename))
	path := filepath.Join(".", "files")
	_ = os.MkdirAll(path, os.ModePerm)
	fullPath := path + "/" + name + fileExtension

	// Пересоздание файла
	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	// Копирование файоа
	_, err = io.Copy(file, f)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":  id,
		"msg": "file upload",
	})
}

func (h *Handler) downloadFile(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	name := c.Request.URL.Query().Get("name")

	directory := filepath.Join("files", name)

	_, err := os.Open(directory)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(directory))
	http.ServeFile(c.Writer, c.Request, directory)
}
