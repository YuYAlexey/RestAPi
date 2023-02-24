package app

import (
	"github.com/adYushinW/RestAPi/internal/db"
	"github.com/adYushinW/RestAPi/internal/model"
)

type App struct {
	db db.Database
}

func New(db db.Database) *App {
	return &App{
		db: db,
	}
}

func (a *App) GetTodo(state string, date1 string, date2 string, sort string, limit string) ([]*model.Todo, error) {
	return a.db.GetTodo(state, date1, date2, sort, limit)
}

func (a *App) GetOnlyOne(id int) ([]*model.Todo, error) {
	return a.db.GetOnlyOne(id)
}

func (a *App) AddNew(state bool, date string, name string) ([]*model.Todo, error) {
	return a.db.AddNew(state, date, name)
}

func (a *App) ChangeStatus(id int, state bool) ([]*model.Todo, error) {
	return a.db.ChangeStatus(id, state)
}

func (a *App) Delete(id int) (bool, error) {
	return a.db.Delete(id)
}
