package repository

import (
	"database/sql"
	"errors"
	"fmt"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/jmoiron/sqlx"
)

// Репозиторий для работ с Сотрудниками
type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Изменение Сотрудника
func (r *UserRepository) Update(userId int, input securefilechanger.UpdateUser) error {
	query := fmt.Sprintf("UPDATE \"%s\" SET name=$1 AND surname=$2 WHERE id=$3", userTable)
	_, err := r.db.Exec(query, input.Name, input.SurName, userId)

	return err
}

func (r *UserRepository) Delete(userId int) error {
	query := fmt.Sprintf("DELETE FROM \"%s\" WHERE id=$1", userTable)
	_, err := r.db.Exec(query, userId)

	return err
}

func (r *UserRepository) NewPassword(userId int, password string) error {
	return nil
}

func (r *UserRepository) IsApproved(userId int) (bool, error) {
	var isApproved bool
	query := fmt.Sprintf("SELECT is_approved FROM \"%s\" WHERE id=$1", userTable)
	err := r.db.Get(&isApproved, query, userId)

	if err == sql.ErrNoRows {
		return isApproved, errors.New("no such user")
	}

	return isApproved, err
}

func (r *UserRepository) IsAdmin(userId int) (bool, error) {
	var isAdmin bool
	query := fmt.Sprintf("SELECT is_admin FROM \"%s\" WHERE id=$1", userTable)
	err := r.db.Get(&isAdmin, query, userId)

	if err == sql.ErrNoRows {
		return isAdmin, errors.New("no such user")
	}

	return isAdmin, err
}

func (r *UserRepository) GetAll(adminId int) ([]securefilechanger.User, error) {
	var users []securefilechanger.User
	query := fmt.Sprintf("SELECT id, email, name, surname, is_approved FROM \"%s\" WHERE id != $1", userTable)
	err := r.db.Select(&users, query, adminId)

	if err == sql.ErrNoRows {
		return users, nil
	}

	return users, err
}

// Выдача доступа Сотруднику
func (r *UserRepository) SetApprove(userId int) error {
	query := fmt.Sprintf("UPDATE \"%s\" SET is_approved=true WHERE id=$1", userTable)
	_, err := r.db.Exec(query, userId)

	return err
}

// Ограничение доступа Сотруднику
func (r *UserRepository) SetDisable(userId int) error {
	query := fmt.Sprintf("UPDATE \"%s\" SET is_approved=false WHERE id=$1", userTable)
	_, err := r.db.Exec(query, userId)

	return err
}
