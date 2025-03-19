package main

import (
	"log"
	"os"

	"go1f/pkg/db"
	"go1f/pkg/server"
)

func main() {
	dbFile := "scheduler.db"
	_, err := os.Stat(dbFile)
	install := err != nil

	if err = db.Open(dbFile); err != nil {
		log.Fatal(err)
	}
	if install {
		if err = db.Install(); err != nil {
			log.Fatal(err)
		}
	}

	if err = server.Run("web"); err != nil {
		log.Fatal(err)
	}
}
