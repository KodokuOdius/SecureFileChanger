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

type filesList struct {
	Data []securefilechanger.File `json:"data"`
}

// Список документов в директории
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

	_, err = h.services.Folder.GetById(userId, folderId)
	if err != nil {
		newErrorMessage(c, http.StatusNotFound, err.Error())
		return
	}

	files, err := h.services.File.GetFilesInFolder(userId, folderId)
	if err != nil {
		logrus.Infoln("[GetFilesInFolder]")
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, filesList{Data: files})
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

	c.JSON(http.StatusOK, statusResponce{Status: "ok"})
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

	fileExtension := strings.ToLower(filepath.Ext(handler.Filename))
	if !securefilechanger.IsAsseptedExtension(fileExtension) {
		newErrorMessage(c, http.StatusBadRequest, fmt.Sprintf("file extension %s does not support", fileExtension))
		return
	}

	c.Writer.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	usedBytes, err := h.services.User.GetUsedBytes(userId)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 5Gb
	userLimit := 1024 * 1024 * 1024 * 5
	if int(handler.Size)+usedBytes > userLimit {
		newErrorMessage(c, http.StatusBadRequest, "exceeded limit (5Gb)")
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
		folder, err = h.services.Folder.GetById(userId, folderId)
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

	h.DownloadOneFile(c, file)

	c.JSON(http.StatusOK, statusResponce{Status: "ok"})
}

// Выгрзка одного документа
func (h *Handler) DownloadOneFile(c *gin.Context, file securefilechanger.File) {
	srcPath := filepath.Join(file.Path, fmt.Sprintf("document%d.enc", file.Id))

	_, err := os.Stat(srcPath)
	if err != nil {
		h.services.File.Delete(file.Id, file.OwnerId)
		newErrorMessage(c, http.StatusNotFound, "file already deleted")
		return
	}

	// Поток дешифрования
	reader, srcFile, err := h.services.File.FileDencrypt(os.Getenv("AES_KEY"), srcPath)
	defer srcFile.Close()
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Writer.Header().Set("Content-Type", "application/%s"+strings.ReplaceAll(file.Type, ".", ""))
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(file.Name+file.Type))

	_, err = io.Copy(c.Writer, reader)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
}

// Изменение имени документа
func (h *Handler) updateFile(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	fileId, err := strconv.Atoi(c.Param("file_id"))
	if err != nil {
		newErrorMessage(c, http.StatusBadRequest, "invalid file id")
		return
	}

	var input securefilechanger.UpdateFile
	if err := c.BindJSON(&input); err != nil {
		newErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.services.File.GetById(fileId, userId)
	if err != nil {
		newErrorMessage(c, http.StatusNotFound, err.Error())
		return
	}

	if err := h.services.File.Update(userId, fileId, input); err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponce{Status: "ok"})
}

// Список документов в корневой директории
func (h *Handler) getFilesInRoot(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	rootId, err := h.services.Folder.GetRoot(userId)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	files, err := h.services.File.GetFilesInFolder(userId, rootId)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, filesList{Data: files})
}
