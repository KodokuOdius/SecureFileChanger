package repository

import (
	"fmt"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/jmoiron/sqlx"
)

// Репозиторий для работ с директориями
type FolderRepository struct {
	db *sqlx.DB
}

func NewFolderRepository(db *sqlx.DB) *FolderRepository {
	return &FolderRepository{db: db}
}

// Создание директории
func (r *FolderRepository) Create(userId int, folder securefilechanger.Folder) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createFolderQuery := fmt.Sprintf("INSERT INTO %s (name, user_id) VALUES ($1, $2) RETURNING id", folderTable)
	row := tx.QueryRow(createFolderQuery, folder.Name, userId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

// Список всех директорий
func (r *FolderRepository) GetAll(userId int) ([]securefilechanger.Folder, error) {
	var folders []securefilechanger.Folder
	query := fmt.Sprintf("SELECT id, name, is_root, is_bin FROM %s WHERE user_id=$1", folderTable)
	err := r.db.Select(&folders, query, userId)

	return folders, err
}

// Данные одной директории
func (r *FolderRepository) GetById(folderId, userId int) (securefilechanger.Folder, error) {
	var folder securefilechanger.Folder
	query := fmt.Sprintf("SELECT id, name, is_root, is_bin FROM %s WHERE folder_id=$1 AND user_id=$2", folderTable)
	err := r.db.Get(&folder, query, folderId, userId)

	return folder, err
}

// Удаление директории
func (r *FolderRepository) Delete(folderId, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2", folderTable)
	_, err := r.db.Exec(query, folderId, userId)

	return err
}

// Изменение директории
func (r *FolderRepository) Update(folderId, userId int, input securefilechanger.UpdateFolder) error {
	if input.Name == nil {
		return nil
	}

	query := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id=$2 AND user_id=$3", folderTable)

	_, err := r.db.Exec(query, input.Name, folderId, userId)
	return err
}
