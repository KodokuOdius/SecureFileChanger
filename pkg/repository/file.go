package repository

import (
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

// Создание директории
func (r *FileRepository) Create(userId int, metaFile securefilechanger.File) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, path, user_id, folder_id) VALUES ($1, $2, $3, $4)", fileTable)
	row := r.db.QueryRow(query, metaFile.Name, metaFile.Path, userId, metaFile.FolderId)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// Список документов в директории
func (r *FileRepository) GetFilesInFolder(userId int, folderId int) ([]securefilechanger.File, error) {
	var files []securefilechanger.File
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1 and folder_id=$2", fileTable)
	err := r.db.Get(&files, query, userId, folderId)

	return files, err
}

// Удаление документа
func (r *FileRepository) Delete(fileId, userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND id=$2", fileTable)
	_, err := r.db.Exec(query, userId, fileId)

	return err
}
