package app

import (
	"github.com/adYushinW/RestAPi/internal/db"
	"github.com/adYushinW/RestAPi/internal/model"
)

type App struct {
	db *db.Database
}

func New(db *db.Database) *App {
	return &App{
		db: db,
	}
}

func (a *App) GetAllTodo() ([]*model.Todo, error) {
	return a.db.GetAllTodo()
}

func (a *App) GetAllSort(sort string) ([]*model.Todo, error) {
	return a.db.GetAllSort(sort)
}

func (a *App) GetAllSortUndone() ([]*model.Todo, error) {
	return a.db.GetAllSortUndone()
}

func (a *App) GetAllState(sort string) ([]*model.Todo, error) {
	return a.db.GetAllState(sort)
}

func (a *App) GetOnlyOne(id string) ([]*model.Todo, error) {
	return a.db.GetOnlyOne(id)
}
