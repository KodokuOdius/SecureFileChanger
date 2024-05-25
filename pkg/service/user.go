package service

import (
	"errors"

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

func (s *UserService) ToggleApprove(userId int) error {
	isApproved, err := s.repo.IsApproved(userId)
	if err != nil {
		return err
	}

	if isApproved {
		return s.repo.SetDisable(userId)
	}

	return s.repo.SetApprove(userId)
}

func (s *UserService) Delete(userId int) error {
	return s.repo.Delete(userId)
}

func (s *UserService) NewPassword(userId int, changePass securefilechanger.ChangePass) error {
	if err := changePass.Validate(); err != nil {
		return err
	}

	changePass.OldPass = hashPassword(changePass.OldPass)
	changePass.NewPass = hashPassword(changePass.NewPass)

	ok, err := s.repo.CheckPassword(userId, changePass.OldPass)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("wrong password")
	}

	return s.repo.NewPassword(userId, changePass)
}

func (s *UserService) IsApproved(userId int) (bool, error) {
	return s.repo.IsApproved(userId)
}

func (s *UserService) IsAdmin(userId int) (bool, error) {
	return s.repo.IsAdmin(userId)
}

func (s *UserService) GetAll(adminId int) ([]securefilechanger.User, error) {
	return s.repo.GetAll(adminId)
}

func (s *UserService) GetInfo(userId int) (securefilechanger.UserInfo, error) {
	return s.repo.GetInfo(userId)
}
