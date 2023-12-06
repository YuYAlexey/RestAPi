package main

import (
	"github.com/YuYAlexey/RestAPi/internal/app"
	"github.com/YuYAlexey/RestAPi/internal/db"
	"github.com/YuYAlexey/RestAPi/internal/log"
	"github.com/YuYAlexey/RestAPi/internal/transport/http"
)

func main() {
	logger, err := log.New("")
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	db, err := db.New()
	if err != nil {
		panic(err)
	}

	app := app.New(db)

	if err := http.Service(app, logger); err != nil {
		panic(err)
	}
}
