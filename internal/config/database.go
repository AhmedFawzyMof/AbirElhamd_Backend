package config

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/sqlite"
)

func Database() *sql.DB {
	db, err := sql.Open("sqlite", "./abirelhamd.db")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
