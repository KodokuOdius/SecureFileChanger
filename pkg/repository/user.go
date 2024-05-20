package repository

import (
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

// Ограничение доступа Сотрудника
func (r *UserRepository) SetDisable(userId int) error {
	query := fmt.Sprintf("UPDATE \"%s\" SET is_approved=false WHERE id=$1", userTable)
	_, err := r.db.Exec(query, userId)

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
	isApproved := true
	query := fmt.Sprintf("SELECT is_approved from \"%s\" WHERE id=$1", userTable)
	err := r.db.Get(&isApproved, query, userId)

	return isApproved, err
}
