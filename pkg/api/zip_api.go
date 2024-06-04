package api

import (
	"archive/zip"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Выгрузка несольких документов
func (h *Handler) downloadManyFiles(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input filesIds
	err = c.BindJSON(&input)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	clearFileIds, err := h.services.Url.CheckFileIds(userId, input.Ids)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(clearFileIds) == 0 {
		newErrorMessage(c, http.StatusBadRequest, "files not found")
		return
	}

	files, err := h.services.File.GetFilesByIds(userId, clearFileIds)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.DownloadFileZip(c, files)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
}

// Выгрузка zip архива
func (h *Handler) DownloadFileZip(c *gin.Context, files []securefilechanger.File) error {
	c.Writer.Header().Set("Content-Type", "application/zip")
	c.Writer.Header().Set("Content-Disposition", "application; filename=archive.zip")
	zipWriter := zip.NewWriter(c.Writer)
	defer zipWriter.Close()

	err := h.addFileToArchive(files, zipWriter)
	if err != nil {
		return err
	}

	return nil
}

// Добавление файла в архив
func (h *Handler) addFileToArchive(files []securefilechanger.File, zipWriter *zip.Writer) error {
	for _, file := range files {
		fileName := file.Name + file.Type
		logrus.Info("[CreateArchive] ", fileName)
		srcPath := filepath.Join(file.Path, fmt.Sprintf("document%d.enc", file.Id))

		_, err := os.Stat(srcPath)
		if err != nil {
			return errors.New("files already deleted")
		}

		reader, srcFile, err := h.services.File.FileDencrypt(os.Getenv("AES_KEY"), srcPath)
		defer srcFile.Close()
		if err != nil {
			return err
		}

		err = h.services.ZipService.AddFileToArchive(file.Name+file.Type, reader, zipWriter)
		if err != nil {
			return err
		}

		logrus.Info("[AddFileToArchive Success] ", fileName)
	}
	return nil
}
