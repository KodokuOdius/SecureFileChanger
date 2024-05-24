package repository

import (
	"database/sql"
	"errors"
	"fmt"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/jmoiron/sqlx"
)

// Репозиторий для авториазции
type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user securefilechanger.User) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var userId int
	query := fmt.Sprintf("INSERT INTO \"%s\" (email, password, is_admin, is_approved) values ($1, $2, $3, $4) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Email, user.Password, user.IsAdmin, user.IsApproved)

	if err := row.Scan(&userId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createFolderQuery := fmt.Sprintf("INSERT INTO %s (name, user_id, is_root, is_bin) VALUES ($1, $2, $3, $4) RETURNING id", folderTable)
	_, err = tx.Exec(createFolderQuery, "root", userId, true, false)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec(createFolderQuery, "bin", userId, false, true)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return userId, tx.Commit()
}

func (r *AuthRepository) GetUser(email, password string) (securefilechanger.User, error) {
	var user securefilechanger.User
	query := fmt.Sprintf("SELECT id, is_approved FROM \"%s\" WHERE email=$1 and password=$2", userTable)
	err := r.db.Get(&user, query, email, password)

	if err == sql.ErrNoRows {
		return user, errors.New("no such email, or invalid password")
	}

	return user, nil
}

func (r *AuthRepository) CheckAdminIsExists() (bool, error) {
	var adminExist bool
	query := fmt.Sprintf("SELECT is_admin FROM \"%s\" WHERE is_admin", userTable)
	err := r.db.Get(&adminExist, query)

	if err == sql.ErrNoRows {
		return adminExist, nil
	}

	return adminExist, err
}
