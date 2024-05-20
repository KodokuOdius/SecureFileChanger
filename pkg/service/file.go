package service

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"os"
	"strings"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/KodokuOdius/SecureFileChanger/pkg/repository"
	"github.com/sirupsen/logrus"
)

// Сервис документов
type FileService struct {
	repo repository.File
}

func NewFileService(repo repository.File) *FileService {
	return &FileService{repo: repo}
}

// Создание нового документа
func (s *FileService) Create(userId int, metaFile securefilechanger.File) (int, error) {
	return s.repo.Create(userId, metaFile)
}

// Список документов в директории
func (s *FileService) GetFilesInFolder(userId int, folderId *int) ([]securefilechanger.File, error) {
	return s.repo.GetFilesInFolder(userId, *folderId)
}

// Удаление документа
func (s *FileService) Delete(fileId, userId int) error {
	return s.repo.Delete(fileId, userId)
}

func (s *FileService) FileEncrypt(fileName string, inputFile io.Reader) (string, error) {
	logrus.Info("[FileEncrypt] encrypte file", fileName)

	encfileName := fileName + ".enc"

	// Создание блоков шифрования
	block, err := aes.NewCipher([]byte(os.Getenv("AES_KEY")))
	if err != nil {
		return "", err
	}

	// Инициализация вектора (вставляется в начале файла)
	var iv [aes.BlockSize]byte

	// Поток для шифрования
	stream := cipher.NewOFB(block, iv[:])

	// Создание файла на сервере
	outFile, err := os.OpenFile(encfileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	// Поток шифрования
	writer := &cipher.StreamWriter{S: stream, W: outFile}

	if _, err := io.Copy(writer, inputFile); err != nil {
		return "", err
	}

	return encfileName, nil
}

func FileDecryption(key string, encfileName string) (string, error) {
	decfileName := strings.ReplaceAll(encfileName, ".enc", "")
	inFile, err := os.Open(encfileName)
	if err != nil {
		return "", err
	}
	defer inFile.Close()

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])
	outFile, err := os.OpenFile(decfileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return "", err
	}
	defer outFile.Close()
	reader := &cipher.StreamReader{S: stream, R: inFile}

	if _, err := io.Copy(outFile, reader); err != nil {
		return "", err
	}
	return decfileName, nil
}
