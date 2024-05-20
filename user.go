package securefilechanger

import "errors"

type User struct {
	Id         int    `json:"-" db:"id"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	IsAdmin    bool   `json:"is_admin" db:"is_admin"`
	IsApproved bool   `json:"is_approved" db:"is_approved"`
}

type UpdateUser struct {
	Name    *string `json:"user_name"`
	SurName *string `json:"user_surname"`
}

func (u UpdateUser) Validate() error {
	if u.Name == nil && u.SurName == nil {
		return errors.New("update has no values")
	}

	return nil
}
