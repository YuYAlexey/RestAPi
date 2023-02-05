package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres
)

// Database структура для подключения к БД
type ConfigDatabase struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLmode  string
}

// String ...
func (d ConfigDatabase) String() string {
	return fmt.Sprintf(
		`host=%s port=%s user=%s password=%s database=%s sslmode=%s`,
		d.Host, d.Port, d.User, d.Password, d.Database, d.SSLmode,
	)
}

func constDB() ConfigDatabase {
	return ConfigDatabase{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Database: "todo",
		SSLmode:  "disable",
	}
}

// ConnectDB создание подключения к БД
func newConnect() (*sql.DB, error) {
	db, err := sql.Open("postgres", constDB().String())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
