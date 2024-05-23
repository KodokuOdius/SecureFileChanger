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
	GetByName(folderName string, userId int) (int, error)
	Delete(folderId, userId int) error
	Update(folderId, userId int, input securefilechanger.UpdateFolder) error
	GetRoot(userId int) (int, error)
	GetBin(userId int) (int, error)
}

// Обработчик операций с документами
type File interface {
	Create(userId int, metaFile securefilechanger.File) (int, error)
	GetFilesInFolder(userId int, folderId int) ([]securefilechanger.File, error)
	Delete(fileId, userId int) error
	GetByName(fileName string, folderId, userId int) (int, error)
	GetById(fileId, userId int) (securefilechanger.File, error)
}

// Обработчик операций с Сотрудниками
type User interface {
	Update(userId int, input securefilechanger.UpdateUser) error
	SetDisable(userId int) error
	Delete(userId int) error
	NewPassword(userId int, password string) error
	IsApproved(userId int) (bool, error)
}

// Структура репозитория
type Repository struct {
	Authorization
	Folder
	File
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Folder:        NewFolderRepository(db),
		File:          NewFileRepository(db),
		User:          NewUserRepository(db),
	}
}
