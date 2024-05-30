package service

import (
	"archive/zip"
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
	CheckAdminIsExists() (bool, error)
}

// Сервис работ с директориями
type Folder interface {
	Create(userId int, folder securefilechanger.Folder) (int, error)
	GetAll(userId int) ([]securefilechanger.Folder, error)
	GetById(userId, folderId int) (securefilechanger.Folder, error)
	GetByName(folderName string, userId int) (int, error)
	Delete(userId, folderId int) error
	Update(userId, folderId int, input securefilechanger.UpdateFolder) error
	GetRoot(userId int) (int, error)
}

// Сервис работ с документами
type File interface {
	Create(userId int, metaFile securefilechanger.File) (int, error)
	Update(userId, fileId int, input securefilechanger.UpdateFile) error
	GetFilesInFolder(userId int, folderId int) ([]securefilechanger.File, error)
	Delete(fileId, userId int) error
	FileEncrypt(fileName string, inputFile io.Reader) (string, error)
	FileDencrypt(key string, encfileName string) (*cipher.StreamReader, *os.File, error)
	GetByName(fileName, fileType string, folderId, userId int) (int, error)
	GetById(fileId, userId int) (securefilechanger.File, error)
	GetFilesByIds(userId int, fileIds []int) ([]securefilechanger.File, error)
}

// Сервис работ с Сотрудниками
type User interface {
	Update(userId int, input securefilechanger.UpdateUser) error
	ToggleApprove(userId int) error
	Delete(userId int) error
	NewPassword(userId int, changePass securefilechanger.ChangePass) error
	IsApproved(userId int) (bool, error)
	IsAdmin(userId int) (bool, error)
	GetAll(adminId int) ([]securefilechanger.User, error)
	GetInfo(userId int) (securefilechanger.UserInfo, error)
	GetLike(adminId int, queryName string) ([]securefilechanger.User, error)
	GetUsedBytes(userId int) (int, error)
}

// Сервис работ с временными ссылками
type Url interface {
	CreateUrl(userId, hourLive int, filesIds []int) (string, error)
	GetUrl(uuid string) (securefilechanger.Url, error)
	GetFilesList(uuid string) ([]securefilechanger.File, error)
	DeleteUrl(uuid string) error
	CheckFileIds(userId int, fileIds []int) ([]int, error)
}

type ZipService interface {
	AddFileToArchive(fileName string, srcFile io.Reader, zipw *zip.Writer) error
}

// Структура сервиса
type Service struct {
	Authorization
	Folder
	File
	User
	Url
	ZipService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Folder:        NewFolderService(repos.Folder),
		File:          NewFileService(repos.File),
		User:          NewUserService(repos.User),
		Url:           NewUrlService(repos.Url),
		ZipService:    NewZipService(),
	}
}
