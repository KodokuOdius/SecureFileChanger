package service

import (
	"crypto/cipher"
	"io"
	"os"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/KodokuOdius/SecureFileChanger/pkg/repository"
)

// Уровень сервисов (бизнес логика)
// Сервис авториазции
type Authorization interface {
	CreateUser(user securefilechanger.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

// Сервис работ с директориями
type Folder interface {
	Create(userId int, folder securefilechanger.Folder) (int, error)
	GetAll(userId int) ([]securefilechanger.Folder, error)
	GetById(folderId, userId int) (securefilechanger.Folder, error)
	Delete(folderId, userId int) error
	Update(folderId, userId int, input securefilechanger.UpdateFolder) error
	CreateDefaultFolder(userId int) error
}

// Сервис работ с документами
type File interface {
	Create(userId int, metaFile securefilechanger.File) (int, error)
	GetFilesInFolder(userId int, folderId *int) ([]securefilechanger.File, error)
	Delete(fileId, userId int) error
	FileEncrypt(fileName string, inputFile io.Reader) (string, error)
	FileDencrypt(key string, encfileName string) (*cipher.StreamReader, *os.File, error)
}

type User interface {
	Update(userId int, input securefilechanger.UpdateUser) error
	SetDisable(userId int) error
	Delete(userId int) error
	NewPassword(userId int, password string) error
	IsApproved(userId int) (bool, error)
}

// Структура сервиса
type Service struct {
	Authorization
	Folder
	File
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Folder:        NewFolderService(repos.Folder),
		File:          NewFileService(repos.File),
		User:          NewUserService(repos.User),
	}
}
