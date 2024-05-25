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
	CheckAdminIsExists() (bool, error)
}

// Обработчик операций с директориями
type Folder interface {
	Create(userId int, folder securefilechanger.Folder) (int, error)
	GetAll(userId int) ([]securefilechanger.Folder, error)
	GetById(userId, folderId int) (securefilechanger.Folder, error)
	GetByName(folderName string, userId int) (int, error)
	Delete(userId, folderId int) error
	Update(userId, folderId int, input securefilechanger.UpdateFolder) error
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
	Delete(userId int) error
	NewPassword(userId int, changePass securefilechanger.ChangePass) error
	IsApproved(userId int) (bool, error)
	IsAdmin(userId int) (bool, error)
	GetAll(adminId int) ([]securefilechanger.User, error)
	SetApprove(userId int) error
	SetDisable(userId int) error
	GetInfo(userId int) (securefilechanger.UserInfo, error)
	CheckPassword(userId int, password string) (bool, error)
}

// Обработчик операций с временными ссылками
type Url interface {
	CreateUrl(userId int, url securefilechanger.Url, filesIds []int) error
	CheckFileIds(userId int, fileIds []int) ([]int, error)
	GetByUUid(uuid string) (securefilechanger.Url, error)
	GetFilesByUrlUUid(uuid string) ([]securefilechanger.File, error)
	DeleteUrl(uuid string) error
}

// Структура репозитория
type Repository struct {
	Authorization
	Folder
	File
	User
	Url
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Folder:        NewFolderRepository(db),
		File:          NewFileRepository(db),
		User:          NewUserRepository(db),
		Url:           NewUrlRepository(db),
	}
}
