package db

import (
	"context"
	_ "embed"
)

//go:embed schema.sql
var schema string

func Install() error {
	_, err := db.ExecContext(context.Background(), schema)
	return err
}
