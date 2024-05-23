package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Список документов в директории
type getAllFiles struct {
	Data []securefilechanger.File `json:"data"`
}

func (h *Handler) getFilesInFolder(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	folderId, err := strconv.Atoi(c.Param("folder_id"))
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, "invalid folder id")
		return
	}

	files, err := h.services.File.GetFilesInFolder(userId, &folderId)
	if err != nil {
		logrus.Infoln("[GetFilesInFolder]")
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllFiles{Data: files})
}

// Удаление документа
func (h *Handler) deleteFile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	fileId, err := strconv.Atoi(c.Param("file_id"))
	if err != nil {
		newErrorMessage(c, http.StatusBadRequest, "invalid file id")
		return
	}

	err = h.services.File.Delete(fileId, userId)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{
		Status: "ok",
	})
}

// Перемещение в корзину
func (h *Handler) toBinFile(c *gin.Context) {
}

// Загрузка документов
func (h *Handler) uploadFile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	srcFile, handler, err := c.Request.FormFile("file")
	// Проверка наличия самого файла
	if err != nil {
		logrus.Infoln("[FormFile]")
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer srcFile.Close()

	c.Writer.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	// 10 MB
	err = c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		logrus.Infoln("[ParseMultipartForm]")
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	folderReq := c.Request.FormValue("folder_id")
	folderId := 0
	if folderReq != "" {
		folderId, err = strconv.Atoi(folderReq)
		if err != nil {
			logrus.Infoln("[Atoi]")
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	var folder securefilechanger.Folder
	if folderId == 0 {
		folderId, err = h.services.Folder.GetRoot(userId)
		if err != nil {
			logrus.Infoln("[GetRoot]")
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		folder, err = h.services.Folder.GetById(folderId, userId)
		if err != nil {
			logrus.Infoln("[GetById]")
			newErrorMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
	}

	var path string
	if folder.Name == "" {
		path = filepath.Join(".", fmt.Sprintf("files/user%d", userId))
	} else {
		path = filepath.Join(".", fmt.Sprintf("files/user%d/folder%d", userId, folder.Id))
	}

	// Создание директорий
	_ = os.MkdirAll(path, os.ModePerm)

	fileExtension := strings.ToLower(filepath.Ext(handler.Filename))
	metaFile := securefilechanger.File{
		Name:      strings.ReplaceAll(handler.Filename, fileExtension, ""),
		Path:      path,
		SizeBytes: int(handler.Size),
		Type:      fileExtension,
		FolderId:  folderId,
	}

	fileId, err := h.services.File.Create(userId, metaFile)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	// fileExtension
	fullPath := path + fmt.Sprintf("/document%d", fileId)
	dstFile, err := h.services.File.FileEncrypt(fullPath, srcFile)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"msg":       "file upload",
		"file_path": dstFile,
		"file_id":   fileId,
		"folder_id": metaFile.FolderId,
	})
}

// Выгрузка документов
func (h *Handler) downloadFile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	c.Writer.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	fileId, err := strconv.Atoi(c.Param("file_id"))
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, "invalid file id")
		return
	}

	file, err := h.services.File.GetById(fileId, userId)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	directory := filepath.Join(file.Path, fmt.Sprintf("document%d.enc", fileId))

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

	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(file.Name+file.Type))
	// gzip writer
	_, err = io.Copy(c.Writer, reader)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
