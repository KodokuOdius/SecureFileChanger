package api

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
	err = c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	serverFile, handler, err := c.Request.FormFile("file")
	// Проверка наличия самого файла
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer serverFile.Close()

	// fileExtension := strings.ToLower(filepath.Ext(handler.Filename))
	// logrus.Info(fileExtension)

	path := filepath.Join(".", fmt.Sprintf("files/user%d", userId))

	// Создание директорий
	_ = os.MkdirAll(path, os.ModePerm)

	fullPath := path + "/" + handler.Filename

	newFile, err := h.services.File.FileEncrypt(fullPath, serverFile)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"msg":     "file upload",
		"encFile": newFile,
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

	_, err = os.Open(directory)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	decrFile, err := FileDecryption(os.Getenv("AES_KEY"), directory)

	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	key := os.Getenv("AES_KEY")
	decfileName := strings.ReplaceAll(directory, ".enc", "")
	logrus.Info("[FileDecryption] dencrypte file", decfileName)

	// Открытие исходного файла
	srcFile, err := os.Open(directory)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer srcFile.Close()

	// Блоки шифрования
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Верктор инициализации
	var iv [aes.BlockSize]byte

	// Поток шифрования
	stream := cipher.NewOFB(block, iv[:])

	// Открытие или создание конечного файла
	outFile, err := os.OpenFile(decfileName, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer outFile.Close()

	// Поток дешифрования
	reader := &cipher.StreamReader{S: stream, R: srcFile}

	if _, err := io.Copy(c.Writer, reader); err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(decrFile))
	http.ServeFile(c.Writer, c.Request, decrFile)
}

func FileDecryption(key string, encfileName string) (string, error) {
	decfileName := strings.ReplaceAll(encfileName, ".enc", "")
	logrus.Info("[FileDecryption] dencrypte file", decfileName)

	// Открытие исходного файла
	inFile, err := os.Open(encfileName)
	if err != nil {
		return "", err
	}
	defer inFile.Close()

	// Блоки шифрования
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Верктор инициализации
	var iv [aes.BlockSize]byte

	// Поток шифрования
	stream := cipher.NewOFB(block, iv[:])

	// Открытие или создание конечного файла
	outFile, err := os.OpenFile(decfileName, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	// Поток дешифрования
	reader := &cipher.StreamReader{S: stream, R: inFile}

	if _, err := io.Copy(outFile, reader); err != nil {
		return "", err
	}
	return decfileName, nil
}
