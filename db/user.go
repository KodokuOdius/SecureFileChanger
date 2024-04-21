package db

import "log"

type User struct {
	Id       uint64 `json:"id"`
	Email    string `json:"email"`
	CreateDt string `json:"create_dt"`
	IsAdmin  bool   `json:"is_admin"`
}

func NewUser(email string, password string, isAdmin bool) *User {
	pg, err := NewPG()
	if err != nil {
		log.Fatal("no pg connection")
		return nil
	}
	userId, err := pg.createNewUser(email, password, isAdmin)
	if err != nil {
		log.Fatal("error create user")
		return nil
	}

	newUser := User{}

	err = pg.getUserByID(&newUser, userId)
	if err != nil {
		log.Fatal("User not found")
		return nil
	}

	defer pg.Close()

	return &newUser
}

func (u *User) CheckPassword(pass string) {

}
