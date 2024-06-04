package repository

import (
	"database/sql"
	"errors"
	"fmt"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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
	query := fmt.Sprintf("UPDATE %s SET name=$1, surname=$2 WHERE id=$3", userTable)
	_, err := r.db.Exec(query, *input.Name, *input.SurName, userId)

	return err
}

// Удаление Сотрудника
func (r *UserRepository) Delete(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", userTable)
	_, err := r.db.Exec(query, userId)

	return err
}

// Смена пароля
func (r *UserRepository) NewPassword(userId int, changePass securefilechanger.ChangePass) error {
	query := fmt.Sprintf("UPDATE %s SET password=$1 WHERE id=$2 AND password=$3", userTable)
	_, err := r.db.Exec(query, changePass.NewPass, userId, changePass.OldPass)

	return err
}

// Проверка пароля
func (r *UserRepository) CheckPassword(userId int, password string) (bool, error) {
	var ok bool
	query := fmt.Sprintf("SELECT true from %s WHERE id=$1 AND password=$2", userTable)
	err := r.db.Get(&ok, query, userId, password)

	if err == sql.ErrNoRows {
		return ok, nil
	}

	return ok, err
}

// Объём загруженных документов
func (r *UserRepository) GetUsedBytes(userId int) (int, error) {
	var usedBytes int
	query := fmt.Sprintf("SELECT COALESCE(SUM(size_bytes), 0) FROM %s WHERE user_id=$1", fileTable)
	err := r.db.Get(&usedBytes, query, userId)

	return usedBytes, err
}

// Проверка доступа
func (r *UserRepository) IsApproved(userId int) (bool, error) {
	var isApproved bool
	query := fmt.Sprintf("SELECT is_approved FROM %s WHERE id=$1", userTable)
	err := r.db.Get(&isApproved, query, userId)

	if err == sql.ErrNoRows {
		return isApproved, errors.New("no such user")
	}

	return isApproved, err
}

// Проверка на доступ к админ панели
func (r *UserRepository) IsAdmin(userId int) (bool, error) {
	var isAdmin bool
	query := fmt.Sprintf("SELECT is_admin FROM %s WHERE id=$1", userTable)
	err := r.db.Get(&isAdmin, query, userId)

	if err == sql.ErrNoRows {
		return isAdmin, errors.New("no such user")
	}

	return isAdmin, err
}

// Список всех сотрудников
func (r *UserRepository) GetAll(adminId int) ([]securefilechanger.User, error) {
	var users []securefilechanger.User
	query := fmt.Sprintf("SELECT id, email, name, surname, is_approved FROM %s WHERE id != $1", userTable)
	err := r.db.Select(&users, query, adminId)

	if err == sql.ErrNoRows {
		return users, nil
	}

	return users, err
}

// Поиск сотрудников
func (r *UserRepository) GetLike(adminId int, queryName string) ([]securefilechanger.User, error) {
	var users []securefilechanger.User
	query := fmt.Sprintf("SELECT id, email, name, surname, is_approved FROM %s WHERE id!=$1 AND email LIKE $2", userTable)
	logrus.Info("[GetLike] ", query)
	err := r.db.Select(&users, query, adminId, queryName+"%")

	if err == sql.ErrNoRows {
		return users, nil
	}

	return users, err
}

// Выдача доступа Сотруднику
func (r *UserRepository) SetApprove(userId int) error {
	query := fmt.Sprintf("UPDATE %s SET is_approved=true WHERE id=$1", userTable)
	_, err := r.db.Exec(query, userId)

	return err
}

// Ограничение доступа Сотруднику
func (r *UserRepository) SetDisable(userId int) error {
	query := fmt.Sprintf("UPDATE %s SET is_approved=false WHERE id=$1", userTable)
	_, err := r.db.Exec(query, userId)

	return err
}

// Информация о Сотруднике
func (r *UserRepository) GetInfo(userId int) (securefilechanger.UserInfo, error) {
	var user securefilechanger.UserInfo
	query := fmt.Sprintf(`
		SELECT email, c.name name, surname, is_admin, is_approved, COALESCE(SUM(size_bytes), 0) used_bytes
		FROM %s c LEFT JOIN %s f
		ON c.id = f.user_id
		WHERE c.id=$1
		GROUP BY email, c.name, surname, is_admin, is_approved`,
		userTable, fileTable,
	)
	err := r.db.Get(&user, query, userId)

	return user, err
}
