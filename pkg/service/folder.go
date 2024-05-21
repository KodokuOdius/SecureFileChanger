package service

import (
	"fmt"
	"os"
	"path/filepath"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/KodokuOdius/SecureFileChanger/pkg/repository"
)

// Сервис директорий
type FolderService struct {
	repo repository.Folder
}

func NewFolderService(repo repository.Folder) *FolderService {
	return &FolderService{repo: repo}
}

// Создание новой директории
func (s *FolderService) Create(userId int, folder securefilechanger.Folder) (int, error) {
	return s.repo.Create(userId, folder)
}

// Список директорий
func (s *FolderService) GetAll(userId int) ([]securefilechanger.Folder, error) {
	return s.repo.GetAll(userId)
}

// Данные директории
func (s *FolderService) GetById(folderId, userId int) (securefilechanger.Folder, error) {
	return s.repo.GetById(folderId, userId)
}

// Удаление директории
func (s *FolderService) Delete(folderId, userId int) error {
	return s.repo.Delete(folderId, userId)
}

// Изменения директории
func (s *FolderService) Update(folderId, userId int, input securefilechanger.UpdateFolder) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(folderId, userId, input)
}

// Создание начальных директорий
func (s *FolderService) CreateDefaultFolder(userId int) error {
	err := s.repo.CreateDefaultFolder(userId)
	if err != nil {
		return err
	}

	path := filepath.Join(".", fmt.Sprintf("files/user%d/bin", userId))
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
