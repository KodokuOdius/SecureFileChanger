package repository

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

// Основные таблицы БД
const (
	userTable    = "cloud_user"
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

// Подключение к базе данных
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

	path := filepath.Join(".", "schema/000001_init.up.sql")
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	logrus.Info("initalization database")
	sql := string(content)
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	_, err = tx.Exec(sql)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return db, nil
}
