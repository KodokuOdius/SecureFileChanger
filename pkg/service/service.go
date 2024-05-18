package service

import (
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
}

type Service struct {
	Authorization
	Folder
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Folder:        NewFolderService(repos.Folder),
	}
}
