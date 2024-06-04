package securefilechanger

import (
	"errors"
	"net/mail"
)

type User struct {
	Id         int     `json:"id" db:"id"`
	Email      string  `json:"email" binding:"required" db:"email"`
	Password   string  `json:"-"`
	Name       *string `json:"user_name" db:"name"`
	SurName    *string `json:"user_surname" db:"surname"`
	IsAdmin    bool    `json:"is_admin" db:"is_admin"`
	IsApproved bool    `json:"is_approved" db:"is_approved"`
}

// Структура для информации о Сотруднике
type UserInfo struct {
	Email      string  `json:"email" db:"email"`
	Name       *string `json:"name" db:"name"`
	SurName    *string `json:"surname" db:"surname"`
	IsAdmin    bool    `json:"is_admin" db:"is_admin"`
	IsApproved bool    `json:"is_approved" db:"is_approved"`
	UsedBytes  int     `json:"used_bytes" db:"used_bytes"`
}

// Структура для смены пароля
type ChangePass struct {
	OldPass string `json:"old_password"`
	NewPass string `json:"new_password"`
}

func (c ChangePass) Validate() error {
	if c.NewPass == "" || c.OldPass == "" {
		return errors.New("no change data")
	}

	if c.OldPass == c.NewPass {
		return errors.New("new password equale old password")
	}

	return nil
}

// Структура для авторизации и регистрации
type AuthInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a AuthInput) Validate() error {
	if len(a.Password) < 8 {
		return errors.New("password is less 8 letters")
	}

	_, err := mail.ParseAddress(a.Email)
	return err
}

// Структура для изменения Сотрудника
type UpdateUser struct {
	Name    *string `json:"user_name"`
	SurName *string `json:"user_surname"`
}

func (u UpdateUser) Validate() error {
	if u.Name == nil && u.SurName == nil {
		return errors.New("update has no values")
	}

	return nil
}
