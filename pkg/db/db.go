package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

func Open(path string) error {
	var err error
	db, err = sqlx.Connect("sqlite3", path)
	return err
}
