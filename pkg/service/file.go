package service

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
	fileId, err := s.GetByName(metaFile.Name, metaFile.FolderId, userId)
	if fileId != 0 {
		return 0, errors.New("file already exists")
	}

	if err != nil {
		return 0, err
	}

	return s.repo.Create(userId, metaFile)
}

// Получение документо по id
func (s *FileService) GetById(fileId, userId int) (securefilechanger.File, error) {
	return s.repo.GetById(fileId, userId)
}

// Получение документо по имени
func (s *FileService) GetByName(fileName string, folderId, userId int) (int, error) {
	return s.repo.GetByName(fileName, folderId, userId)
}

// Список документов в директории
func (s *FileService) GetFilesInFolder(userId int, folderId *int) ([]securefilechanger.File, error) {
	return s.repo.GetFilesInFolder(userId, *folderId)
}

// Получение документов по id
func (s *FileService) GetFilesByIds(userId int, fileIds []int) ([]securefilechanger.File, error) {
	return s.repo.GetFilesByIds(userId, fileIds)
}

// Удаление документа
func (s *FileService) Delete(fileId, userId int) error {
	file, err := s.repo.GetById(fileId, userId)
	if err != nil {
		return err
	}

	err = s.repo.Delete(fileId, userId)
	if err != nil {
		return err
	}

	fullPath := filepath.Join(file.Path, fmt.Sprintf("/document%d.enc", fileId))
	return os.Remove(fullPath)
}

// Шифрование документа
func (s *FileService) FileEncrypt(fileName string, inputFile io.Reader) (string, error) {
	logrus.Info("[FileEncrypt] encrypte file ", fileName)

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

// Дешифрование документа
func (s *FileService) FileDencrypt(key string, encfileName string) (*cipher.StreamReader, *os.File, error) {
	decfileName := strings.ReplaceAll(encfileName, ".enc", "")
	logrus.Info("[FileDecryption] dencrypte file ", decfileName)

	// Открытие исходного файла (защифрованный)
	srcFile, err := os.Open(encfileName)
	if err != nil {
		return nil, nil, err
	}

	// Блоки шифрования
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, nil, err
	}

	// Верктор инициализации
	var iv [aes.BlockSize]byte

	// Поток шифрования
	stream := cipher.NewOFB(block, iv[:])

	reader := &cipher.StreamReader{S: stream, R: srcFile}

	return reader, srcFile, nil
}
