package securefilechanger

type User struct {
	Id       int    `json:"-" db:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Is_admin bool   `json:"is_admin"`
}
