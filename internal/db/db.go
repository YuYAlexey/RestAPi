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

func (db *Database) GetTodo(state string, date1 string, date2 string, sort string, limit string) ([]*model.Todo, error) {

	query := "SELECT id, state, date, name FROM info"

	where := ""
	if state != "" {
		where = fmt.Sprintf("WHERE state = %s", state)
	}

	date := ""
	switch {
	case date1 != "" && date2 == "":
		date = fmt.Sprintf("date = '%s'", date1)
	case date1 != "" && date2 != "":
		date = fmt.Sprintf("date BETWEEN '%s' AND '%s'", date1, date2)
	}

	if date != "" {
		if where != "" {
			where = fmt.Sprintf("%s AND %s", where, date)
		} else {
			where = fmt.Sprintf("WHERE %s", date)
		}
	}

	query = fmt.Sprintf("%s %s", query, where)

	if sort != "" {
		query = fmt.Sprintf("%s ORDER BY %s", query, sort)
	}

	if limit != "" {
		query = fmt.Sprintf("%s LIMIT %s", query, limit)
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

func (db *Database) GetOnlyOne(id int) ([]*model.Todo, error) {

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

func (db *Database) AddNew(state bool, date string, name string) ([]*model.Todo, error) {

	query := "INSERT INTO info (state, date, name) VALUES ($1, $2, $3) RETURNING id, state, date, name"

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

func (db *Database) ChangeStatus(id int, state bool) ([]*model.Todo, error) {

	query := "UPDATE info SET state = $2 WHERE id = $1 RETURNING id, state, date, name"

	row := db.conn.QueryRow(query, id, state)

	result := make([]*model.Todo, 0)

	todo := new(model.Todo)

	err := row.Scan(&todo.ID, &todo.State, &todo.Date, &todo.Name)

	if err != nil {
		return result, err
	}

	result = append(result, todo)

	return result, err
}

func (db *Database) Delete(id int) (bool, error) {

	query := "DELETE FROM info WHERE id = $1"

	row := db.conn.QueryRow(query, id)

	todo := new(model.Todo)

	err := row.Scan(&todo.ID)

	if err == sql.ErrNoRows {
		return true, err
	}

	return false, err
}
