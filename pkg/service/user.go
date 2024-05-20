package service

import (
	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/KodokuOdius/SecureFileChanger/pkg/repository"
)

// Сервис Сотрудников
type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Update(userId int, input securefilechanger.UpdateUser) error {
	return s.repo.Update(userId, input)
}

func (s *UserService) SetDisable(userId int) error {
	return s.repo.SetDisable(userId)
}

func (s *UserService) Delete(userId int) error {
	return s.repo.Delete(userId)
}

func (s *UserService) NewPassword(userId int, password string) error {
	return s.repo.NewPassword(userId, password)
}

func (s *UserService) IsApproved(userId int) (bool, error) {
	return s.repo.IsApproved(userId)
}
