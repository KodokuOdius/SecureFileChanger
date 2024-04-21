package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "cloud"
)

type PG struct {
	DB *sql.DB
}

func NewPG() (*PG, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("no database connection")
		return nil, err
	}

	return &PG{DB: db}, nil
}

func (pg *PG) Close() error {
	return pg.DB.Close()
}

func (pg *PG) createNewUser(email string, password string, isAdmin bool) (int, error) {
	_, err := pg.DB.Exec("insert into \"user\" (email, password, is_admin) values ($1, $2, $3)", email, password, isAdmin)
	if err != nil {
		return 0, err
	}

	rows, _ := pg.DB.Query("select max(id) from \"user\"")

	defer rows.Close()
	var maxID int
	for rows.Next() {
		if err := rows.Scan(&maxID); err != nil {
			log.Fatal("Error parsing row:", err)
			return 0, nil
		}
	}

	return maxID, nil
}

func (pg *PG) getUserByID(user *User, id int) error {
	rows, err := pg.DB.Query("select id, email, created_at, is_admin from \"user\" where id = $1", id)
	if err != nil {
		log.Fatal("Error quering data:", err)
		return err
	}

	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Email, &user.CreateDt, &user.IsAdmin); err != nil {
			log.Fatal("Error parsing row:", err)
			return err
		}
	}

	return nil
}
