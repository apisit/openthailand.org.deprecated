package database

import (
	"database/sql"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var (
	DB_URL = os.Getenv("DB_URL")
)
var db *sql.DB
var once sync.Once

func Connect() *sql.DB {
	once.Do(func() {
		var err error
		if DB_URL == "" {
			DB_URL = "host=localhost dbname=openthailand sslmode=disable"
		}
		db, err = sql.Open("postgres", DB_URL)
		if err != nil {
			log.Printf("cannot open database %v", err)
			db = nil
		}
	})
	return db
}
