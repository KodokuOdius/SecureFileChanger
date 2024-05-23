package repository

import (
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
	query := fmt.Sprintf("INSERT INTO \"%s\" (email, password, is_admin) values ($1, $2, true) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Email, user.Password)

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
	query := fmt.Sprintf("SELECT id FROM \"%s\" WHERE email=$1 and password=$2", userTable)
	err := r.db.Get(&user, query, email, password)

	if err != nil || user.Id == 0 {
		return user, err
	}
	return user, nil
}
