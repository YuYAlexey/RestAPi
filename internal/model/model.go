package model

import (
	"time"
)

type Todo struct {
	ID    int64     `json:"id"`
	State bool      `json:"state"`
	Date  time.Time `json:"date"`
	Name  string    `json:"name"`
}
