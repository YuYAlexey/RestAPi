package log

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

type Logger interface {
	Error(r *http.Request, level string, code int, errr error)
	Info(r *http.Request, level string, code int, errr error)
	Close()
}

type logger struct {
	file      *os.File
	loggerErr zerolog.Logger
	loggerOut zerolog.Logger
}

func New(file string) (Logger, error) {
	var err error
	l := new(logger)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if file != "" {
		l.file, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}

		l.loggerOut = zerolog.New(l.file).With().Timestamp().Logger()
		l.loggerErr = l.loggerOut
	} else {
		l.loggerOut = zerolog.New(os.Stdout).With().Timestamp().Logger()
		l.loggerErr = zerolog.New(os.Stderr).With().Timestamp().Logger()
	}

	return l, nil
}

func (l *logger) Error(r *http.Request, level string, code int, errr error) {

	l.loggerErr.Error().
		Str("level", level).
		Str("method", r.Method).
		Str("URL", r.RequestURI).
		Str("user_agent", r.UserAgent()).
		Int("code", code).
		Err(errr).
		Send()

}

func (l *logger) Info(r *http.Request, level string, code int, errr error) {
	l.loggerOut.Info().
		Str("level", level).
		Str("method", r.Method).
		Str("URL", r.RequestURI).
		Str("user_agent", r.UserAgent()).
		Int("code", code).
		Err(errr).
		Send()
}

func (l *logger) Close() {
	if l.file != nil {
		l.file.Close()
	}
}
