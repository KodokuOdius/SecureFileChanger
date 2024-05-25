package securefilechanger

import (
	"errors"
	"time"
)

type Url struct {
	Id       int       `json:"-" db:"id"`
	UUid     string    `json:"uuid" db:"uuid" binding:"required"`
	HourLive int       `json:"hour_live" db:"hour_live"`
	CreateDt time.Time `json:"create_dt" db:"create_dt"`
}

type UrlFile struct {
	Id     int `json:"-" db:"id"`
	FileId int `json:"-" db:"file_id"`
	UrlId  int `json:"-" db:"url_id"`
}

func (u Url) Validate() error {
	if u.HourLive < 0 || u.HourLive > 24 {
		return errors.New("invalid uuid url")
	}

	return nil
}
