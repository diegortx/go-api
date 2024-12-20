package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "go_db"
	// host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "go_db"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected! " + dbname)

	return db, nil
}
