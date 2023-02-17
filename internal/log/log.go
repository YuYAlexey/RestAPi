package log

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

func Error(r *http.Request, level string, code int, errr error) {
	file, err := os.OpenFile("restapi.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(file).With().Timestamp().Logger()
	logger.Error().
		Str("level", level).
		Str("method", r.Method).
		Str("URL", r.RequestURI).
		Str("method", r.Method).
		Str("user_agent", r.UserAgent()).
		Int("code", code).
		Err(errr).
		Send()

}

func Info(r *http.Request, level string, code int, errr error) {
	file, err := os.OpenFile("restapi.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(file).With().Timestamp().Logger()
	logger.Info().
		Str("level", level).
		Str("method", r.Method).
		Str("URL", r.RequestURI).
		Str("method", r.Method).
		Str("user_agent", r.UserAgent()).
		Int("code", code).
		Err(errr).
		Send()
}
