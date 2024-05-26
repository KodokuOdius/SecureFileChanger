package service

import (
	"errors"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/KodokuOdius/SecureFileChanger/pkg/repository"
	uuid "github.com/nu7hatch/gouuid"
)

// Сервис временных ссылок
type UrlService struct {
	repo repository.Url
}

func NewUrlService(repo repository.Url) *UrlService {
	return &UrlService{repo: repo}
}

// Создание временной ссылки
func (s *UrlService) CreateUrl(userId, hourLive int, filesIds []int) (string, error) {
	cleadIds, err := s.repo.CheckFileIds(userId, filesIds)
	if err != nil {
		return "", err
	}

	if len(cleadIds) == 0 {
		return "", errors.New("files not found")
	}

	uuidBytes, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	url := securefilechanger.Url{
		UUid:     uuidBytes.String(),
		HourLive: hourLive,
	}

	err = s.repo.CreateUrl(userId, url, cleadIds)
	if err != nil {
		return "", err
	}

	return url.UUid, nil
}

// Полчение данных о временной ссылки
func (s *UrlService) GetUrl(uuid string) (securefilechanger.Url, error) {
	return s.repo.GetByUUid(uuid)
}

// Список документов по временной ссылки
func (s *UrlService) GetFilesList(uuid string) ([]securefilechanger.File, error) {
	return s.repo.GetFilesByUrlUUid(uuid)
}

// Удаление временной ссылки
func (s *UrlService) DeleteUrl(uuid string) error {
	return s.repo.DeleteUrl(uuid)
}

// Проверка id документов
func (s *UrlService) CheckFileIds(userId int, fileIds []int) ([]int, error) {
	return s.repo.CheckFileIds(userId, fileIds)
}
