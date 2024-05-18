package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	userTable    = "user"
	fileTable    = "file"
	folderTable  = "folder"
	urlTable     = "upload_url"
	fileUrlTable = "file_url"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Password))

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
