package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func GetDbConnection() *sqlx.DB {
	if db == nil {
		createDbConnection()
	}
	return db
}

func createDbConnection() {
	log.Printf("Opening Database connection to postgres")
	dbConnectionStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))
	var err error
	db, err = sqlx.Open("postgres", dbConnectionStr)
	_ = db
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Database connection opened")
}
