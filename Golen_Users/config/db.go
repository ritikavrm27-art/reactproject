package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {

//connString := "postgres://postgres:password@localhost:5432/databasename?sslmode=disable"
	connString := "postgres://postgres:postgres@blog-db:5432/users_golen?sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connString)

	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("PostgreSQL Database Connected")
}
func GetDB() *sql.DB {
	return DB
}