package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Создание документа
func (h *Handler) createFile(c *gin.Context) {
}

// Список документов в директории
func (h *Handler) getFilesInFolder(c *gin.Context) {
}

// Удаление документа
func (h *Handler) deleteFile(c *gin.Context) {
}

// Обработчик для операция с документами
// Загрузка документов
func (h *Handler) uploadFile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	c.Writer.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	// 10 MB
	err = c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	srcFile, handler, err := c.Request.FormFile("file")
	// Проверка наличия самого файла
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer srcFile.Close()

	// fileExtension := strings.ToLower(filepath.Ext(handler.Filename))
	// logrus.Info(fileExtension)

	path := filepath.Join(".", fmt.Sprintf("files/user%d", userId))

	// Создание директорий
	_ = os.MkdirAll(path, os.ModePerm)

	fullPath := path + "/" + handler.Filename

	dstFile, err := h.services.File.FileEncrypt(fullPath, srcFile)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"msg":     "file upload",
		"encFile": dstFile,
	})
}

// Выгрузка документов
func (h *Handler) downloadFile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	c.Writer.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	name := c.Request.URL.Query().Get("name")
	path := filepath.Join(".", fmt.Sprintf("files/user%d", userId))
	directory := filepath.Join(path, name+".enc")

	// If Exists
	_, err = os.Open(directory)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Поток дешифрования
	reader, srcFile, err := h.services.File.FileDencrypt(os.Getenv("AES_KEY"), directory)
	defer srcFile.Close()
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(name))
	// gzip writer
	_, err = io.Copy(c.Writer, reader)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
}
