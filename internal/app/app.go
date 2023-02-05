package app

import (
	"test/internal/db"
	"test/internal/model"
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

func (a *App) GetAllSortByDate() ([]*model.Todo, error) {
	return a.db.GetAllSortByDate()
}

func (a *App) GetAllSortUnDone() ([]*model.Todo, error) {
	return a.db.GetAllSortUnDone()
}

func (a *App) GetAllState(sort string) ([]*model.Todo, error) {
	return a.db.GetAllState(sort)
}

func (a *App) GetOnlyOne() ([]*model.Todo, error) {
	return a.db.GetOnlyOne()
}
