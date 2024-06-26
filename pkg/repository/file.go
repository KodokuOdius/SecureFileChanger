package repository

import (
	"database/sql"
	"errors"
	"fmt"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/jmoiron/sqlx"
)

// Репозиторий для работ с документами
type FileRepository struct {
	db *sqlx.DB
}

func NewFileRepository(db *sqlx.DB) *FileRepository {
	return &FileRepository{db: db}
}

// Создание документа в БД
func (r *FileRepository) Create(userId int, metaFile securefilechanger.File) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, path, size_bytes, type, user_id, folder_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", fileTable)
	row := r.db.QueryRow(query, metaFile.Name, metaFile.Path, metaFile.SizeBytes, metaFile.Type, userId, metaFile.FolderId)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// Изменение имени документа
func (r *FileRepository) Update(userId, fileId int, input securefilechanger.UpdateFile) error {
	if input.Name == nil {
		return nil
	}

	query := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id=$2 AND user_id=$3", fileTable)

	_, err := r.db.Exec(query, input.Name, fileId, userId)
	return err
}

// Получение документа по id
func (r *FileRepository) GetById(fileId, userId int) (securefilechanger.File, error) {
	var file securefilechanger.File
	query := fmt.Sprintf("SELECT id, name, path, size_bytes, type, folder_id, user_id FROM %s WHERE id=$1 AND user_id=$2", fileTable)
	err := r.db.Get(&file, query, fileId, userId)

	if err == sql.ErrNoRows {
		return file, errors.New("file not found")
	}

	return file, err
}

// Получение документа по имени
func (r *FileRepository) GetByName(fileName, fileType string, folderId, userId int) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE user_id=$1 AND folder_id=$2 AND name=$3 AND type=$4", fileTable)
	err := r.db.Get(&id, query, userId, folderId, fileName, fileType)

	if err == sql.ErrNoRows {
		return id, nil
	}

	return id, err
}

// Список документов в директории
func (r *FileRepository) GetFilesInFolder(userId int, folderId int) ([]securefilechanger.File, error) {
	var files []securefilechanger.File
	query := fmt.Sprintf("SELECT id, name, path, size_bytes, type, folder_id, user_id FROM %s WHERE user_id=$1 and folder_id=$2", fileTable)
	err := r.db.Select(&files, query, userId, folderId)

	return files, err
}

// Удаление документа
func (r *FileRepository) Delete(fileId, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND id=$2", fileTable)
	_, err := r.db.Exec(query, userId, fileId)

	return err
}

// Полчение списка документов по id
func (r *FileRepository) GetFilesByIds(userId int, fileIds []int) ([]securefilechanger.File, error) {
	var files []securefilechanger.File
	query := fmt.Sprintf("SELECT id, name, path, size_bytes, type, user_id FROM %s WHERE user_id=%d and id IN(?)", fileTable, userId)
	query, args, err := sqlx.In(query, fileIds)
	if err != nil {
		return files, err
	}

	query = r.db.Rebind(query)

	err = r.db.Select(&files, query, args...)
	if err == sql.ErrNoRows {
		return files, nil
	}

	return files, err
}
