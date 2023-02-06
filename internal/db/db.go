package db

import (
	"database/sql"
	"fmt"

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

func (db *Database) GetAllSort(sort string) ([]*model.Todo, error) {

	if len(sort) == 0 {
		return nil, nil
	}

	query := fmt.Sprintf("SELECT id, state, date, name FROM info ORDER BY %s ", sort)

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

func (db *Database) GetAllSortUndone() ([]*model.Todo, error) {

	query := "SELECT id, state, date, name FROM info ORDER BY state"

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

func (db *Database) GetAllState(sort string) ([]*model.Todo, error) {
	query := "SELECT id, state, date, name FROM info WHERE state = $1"

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

func (db *Database) GetOnlyOne(id string) ([]*model.Todo, error) {

	query := "SELECT id, state, date, name FROM info WHERE ID = $1"

	row := db.conn.QueryRow(query, id)

	result := make([]*model.Todo, 0)

	todo := new(model.Todo)

	err := row.Scan(&todo.ID, &todo.State, &todo.Date, &todo.Name)

	if err != nil {
		return result, nil
	}

	result = append(result, todo)

	return result, nil
}

func (db *Database) GetSomeRow(state string, date1 string, date2 string, limit string) ([]*model.Todo, error) {

	query := "SELECT id, state, date, name FROM info"

	switch {
	case limit != "" && state == "":
		query = fmt.Sprintf(query+" LIMIT %s", limit)
	case date1 != "" && limit != "" && date2 == "":
		query = fmt.Sprintf(query+" WHERE date = '%s' LIMIT '%s'", date1, limit)
	case date1 != "" && date2 != "" && limit == "" && state == "":
		query = fmt.Sprintf(query+" WHERE date BETWEEN '%s' AND '%s' LIMIT %s", date1, date2, limit)
	case limit != "" && state != "":
		query = fmt.Sprintf(query+" WHERE state = %s LIMIT '%s'", state, limit)
	default:
		return nil, nil
	}

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
		return result, err
	}

	return result, nil
}

func (db *Database) AddNew(state bool, date string, name string) ([]*model.Todo, error) {
	fmt.Println(state, date, name)
	query := "INSERT INTO info (state, date, name) VALUES ($1, $2, $3) RETURNING id, state, date, name"

	fmt.Println(query)

	row := db.conn.QueryRow(query, state, date, name)

	result := make([]*model.Todo, 0)

	todo := new(model.Todo)

	err := row.Scan(&todo.ID, &todo.State, &todo.Date, &todo.Name)

	if err != nil {
		return result, err
	}

	result = append(result, todo)

	return result, err
}
