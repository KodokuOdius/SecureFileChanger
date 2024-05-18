package repository

import (
	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/jmoiron/sqlx"
)

// Уровень репозитория (работа с данными)
// Обработчик операций авторизации
type Authorization interface {
	CreateUser(user securefilechanger.User) (int, error)
	GetUser(email, password string) (securefilechanger.User, error)
}

// Обработчик операций с директориями
type Folder interface {
	Create(userId int, folder securefilechanger.Folder) (int, error)
	GetAll(userId int) ([]securefilechanger.Folder, error)
	GetById(folderId, userId int) (securefilechanger.Folder, error)
	Delete(folderId, userId int) error
	Update(folderId, userId int, input securefilechanger.UpdateFolder) error
}

// Структура репозитория
type Repository struct {
	Authorization
	Folder
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Folder:        NewFolderRepository(db),
	}
}
