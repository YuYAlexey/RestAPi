package main

import (
	"github.com/adYushinW/RestAPi/internal/app"
	"github.com/adYushinW/RestAPi/internal/db"
	"github.com/adYushinW/RestAPi/internal/log"
	"github.com/adYushinW/RestAPi/internal/transport/http"
)

func main() {
	logger, err := log.New("")
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Error(nil, "Error", 404, nil)
	logger.Info(nil, "Info", 200, nil)

	db, err := db.New()
	if err != nil {
		panic(err)
	}

	app := app.New(db)

	if err := http.Service(app, logger); err != nil {
		panic(err)
	}
}
