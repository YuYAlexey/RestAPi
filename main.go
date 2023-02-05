package main

import (
	"log"
	"os"
	"test/internal/app"
	"test/internal/db"
	"test/internal/transport/http"
)

func main() {
	db, err := db.New()
	if err != nil {
		log.Fatalln(err)
		os.Exit(2)
	}

	app := app.New(db)

	if err := http.Service(app); err != nil {
		log.Fatalln(err)
		os.Exit(2)
	}
}
