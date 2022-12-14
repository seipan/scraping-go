package db

import (
	"database/sql"
	"log"

	"github.com/seipan/scraping-go/config"
)

func NewDriver() (*sql.DB, error) {
	dbDSN, err := config.DSN()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", dbDSN)
	if err != nil {
		log.Println("db connect failed")
		panic(err)
	}
	log.Println("db connect success")

	return db, nil
}
