package log

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

func Error(r *http.Request, level string, code int, errr error) {
	file, err := os.OpenFile(
		"RESTAPI.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	logger.Error().
		Str("level", level).
		Str("URL", r.RequestURI).
		Str("method", r.Method).
		Int("code", code).
		Err(errr)
}

func Info(r *http.Request, level string, code int, errr error) {
	file, err := os.OpenFile(
		"RESTAPI.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	logger.Info().
		Str("level", level).
		Str("URL", r.RequestURI).
		Str("method", r.Method).
		Int("code", code).
		Err(errr)
}
