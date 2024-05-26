package service

import (
	"errors"
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
	folderId, err := s.GetByName(folder.Name, userId)
	if folderId != 0 {
		return 0, errors.New("folder already exists")
	}

	if err != nil {
		return 0, err
	}

	return s.repo.Create(userId, folder)
}

// Полчение директории по имени
func (s *FolderService) GetByName(folderName string, userId int) (int, error) {
	return s.repo.GetByName(folderName, userId)
}

// Список директорий
func (s *FolderService) GetAll(userId int) ([]securefilechanger.Folder, error) {
	return s.repo.GetAll(userId)
}

// Данные директории
func (s *FolderService) GetById(userId, folderId int) (securefilechanger.Folder, error) {
	folder, err := s.repo.GetById(userId, folderId)

	if folder.Id == 0 {
		return folder, errors.New("folder not found")
	}

	return folder, err
}

// id корневой директории
func (s *FolderService) GetRoot(userId int) (int, error) {
	return s.repo.GetRoot(userId)
}

// Удаление директории
func (s *FolderService) Delete(userId, folderId int) error {
	err := s.repo.Delete(userId, folderId)
	if err != nil {
		return err
	}

	fullPath := filepath.Join(".", fmt.Sprintf("files/user%d/folder%d", userId, folderId))
	return os.RemoveAll(fullPath)
}

// Изменения директории
func (s *FolderService) Update(userId, folderId int, input securefilechanger.UpdateFolder) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, folderId, input)
}
