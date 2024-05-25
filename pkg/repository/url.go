package repository

import (
	"database/sql"
	"errors"
	"fmt"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/jmoiron/sqlx"
)

// Репозиторий для работ с документами
type UrlRepository struct {
	db *sqlx.DB
}

func NewUrlRepository(db *sqlx.DB) *UrlRepository {
	return &UrlRepository{db: db}
}

func (r *UrlRepository) CreateUrl(userId int, url securefilechanger.Url, filesIds []int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var urlId int
	queryUrl := fmt.Sprintf("INSERT INTO %s (uuid, hour_live) VALUES ($1, $2) RETURNING id", urlTable)
	row := tx.QueryRow(queryUrl, url.UUid, url.HourLive)
	if err := row.Scan(&urlId); err != nil {
		tx.Rollback()
		return err
	}

	for _, fileId := range filesIds {
		var fileUrlId int
		queryFileUrl := fmt.Sprintf("INSERT INTO %s (file_id, url_id) VALUES ($1, $2) RETURNING id", fileUrlTable)
		row := tx.QueryRow(queryFileUrl, fileId, urlId)
		if err := row.Scan(&fileUrlId); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *UrlRepository) CheckFileIds(userId int, fileIds []int) ([]int, error) {
	var clearIds []int

	query := fmt.Sprintf("SELECT id FROM %s WHERE user_id=%d AND id IN (?)", fileTable, userId)

	query, args, err := sqlx.In(query, fileIds)
	if err != nil {
		return clearIds, err
	}

	query = r.db.Rebind(query)

	err = r.db.Select(&clearIds, query, args...)
	if err == sql.ErrNoRows {
		return clearIds, nil
	}

	return clearIds, err
}

func (r *UrlRepository) GetByUUid(uuid string) (securefilechanger.Url, error) {
	var url securefilechanger.Url
	query := fmt.Sprintf("SELECT id, uuid, hour_live, create_dt FROM %s WHERE uuid=$1", urlTable)
	err := r.db.Get(&url, query, uuid)

	if err == sql.ErrNoRows {
		return url, nil
	}

	return url, err
}

func (r *UrlRepository) GetFilesByUrlUUid(uuid string) ([]securefilechanger.File, error) {
	var files []securefilechanger.File

	query := fmt.Sprintf(`
		SELECT file_id AS id, name, path, size_bytes, type
		FROM %s u
		JOIN %s fu
		ON u.id = fu.url_id
		JOIN %s f
		ON fu.file_id = f.id
		WHERE u.uuid = $1
	`, urlTable, fileUrlTable, fileTable)

	err := r.db.Select(&files, query, uuid)

	if err == sql.ErrNoRows {
		return files, errors.New("files not found")
	}

	return files, nil
}

func (r *UrlRepository) DeleteUrl(uuid string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE uuid=$1", urlTable)
	_, err := r.db.Exec(query, uuid)

	return err
}
