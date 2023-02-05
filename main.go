package main

import (
	"log"
	"os"

	"github.com/adYushinW/RestAPi/internal/app"
	"github.com/adYushinW/RestAPi/internal/db"
	"github.com/adYushinW/RestAPi/internal/transport/http"
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
