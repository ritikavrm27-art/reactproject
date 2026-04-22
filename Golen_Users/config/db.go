package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {

//connString := "postgres://postgres:password@localhost:5432/databasename?sslmode=disable"
	connString := "postgres://postgres:postgres@user-db:5432/users_golen?sslmode=disable"

	var err error

	// Retry logic
	for i := 0; i < 20; i++ {
		DB, err = sql.Open("postgres", connString)
		if err != nil {
			log.Println("Error opening DB:", err)
		} else {
			err = DB.Ping()
			if err == nil {
				log.Println("PostgreSQL Database Connected")
				return
			}
			log.Println("DB not ready yet, retrying...")
		}

		time.Sleep(3 * time.Second)
	}

	log.Fatal("Could not connect to DB after multiple attempts")
}
func GetDB() *sql.DB {
	return DB
}