package db

import (
	"database/sql"

	"github.com/adYushinW/RestAPi/internal/model"
)

type Database struct {
	conn *sql.DB
}

func New() (*Database, error) {
	conn, err := newConnect()
	if err != nil {
		return nil, err
	}

	return &Database{
		conn: conn,
	}, nil
}

func (db *Database) GetAllTodo() ([]*model.Todo, error) {
	query := "SELECT id, state, date, name FROM info"

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*model.Todo, 0)

	for rows.Next() {
		todo := new(model.Todo)

		if err := rows.Scan(&todo.ID, &todo.State, &todo.Date, &todo.Name); err != nil {
			continue
		}

		result = append(result, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (db *Database) GetAllSortByDate() ([]*model.Todo, error) {
	query := "SELECT id, state, date, name FROM info ORDER BY date"

	rows, err := db.conn.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*model.Todo, 0)

	for rows.Next() {
		todo := new(model.Todo)

		if err := rows.Scan(&todo.ID, &todo.State, &todo.Date, &todo.Name); err != nil {
			continue
		}

		result = append(result, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// TODO: Ебать тебя за *. Исправить на поля
// TODO: ASC по дефолту, можно убрать
func (db *Database) GetAllSortUndone() ([]*model.Todo, error) {
	query := "SELECT * FROM info ORDER BY state ASC"

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*model.Todo, 0)

	for rows.Next() {
		todo := new(model.Todo)

		if err := rows.Scan(&todo.ID, &todo.State, &todo.Date, &todo.Name); err != nil {
			continue
		}

		result = append(result, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// TODO: Заменить регистр в запросе
func (db *Database) GetAllState(sort string) ([]*model.Todo, error) {
	query := "select * from info where state = $1"

	rows, err := db.conn.Query(query, sort)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*model.Todo, 0)

	for rows.Next() {
		todo := new(model.Todo)

		if err := rows.Scan(&todo.ID, &todo.State, &todo.Date, &todo.Name); err != nil {
			continue
		}

		result = append(result, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// TODO: сделать получение по  ID
// TODO: Заменить регистр в запросе
// TODO: Зачем тут массив в return? убрать
func (db *Database) GetOnlyOne() ([]*model.Todo, error) {

	query := "select * from info limit 1"

	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make([]*model.Todo, 0)

	for rows.Next() {
		todo := new(model.Todo)

		if err := rows.Scan(&todo.ID, &todo.State, &todo.Date, &todo.Name); err != nil {
			continue
		}

		result = append(result, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
