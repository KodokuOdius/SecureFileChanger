package repository

import (
	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user securefilechanger.User) (int, error)
	GetUser(email, password string) (securefilechanger.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
