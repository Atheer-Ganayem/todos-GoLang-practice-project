package models

import (
	"errors"
	"fmt"
	"time"

	db "todo-golang.com/DB"
)

type Todo struct {
	ID   int64 `json:"id"`
	Text string `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	UserId int64  `json:"userId"`
}

func (todo *Todo) Save() error {
	if todo.Text == "" {
		return errors.New("todo's text is empty")
	}
	query := "INSERT INTO todos(text, createdAt, user_id) VALUES (?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(todo.Text, todo.CreatedAt, todo.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {return err}

	todo.ID = id

	return nil
}

func GetTodos() ([]Todo, error) {
	query := "SELECT * FROM todos"
	rows, err := db.DB.Query(query)
	if err != nil {return nil, err}
	defer rows.Close()

	var todos []Todo

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Text, &todo.CreatedAt, &todo.UserId)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func GetTodo(todoId int64) (*Todo, error) {
	query := "SELECT * FROM todos WHERE id = ?"
	row := db.DB.QueryRow(query, todoId)

	var todo Todo
	if err := row.Scan(&todo.ID, &todo.Text, &todo.CreatedAt, &todo.UserId); err != nil {
		return nil, err
	}

	return &todo, nil
}

func (todo *Todo) Delete() error {
	query := "DELETE FROM todos WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {return err}
	defer stmt.Close()

	_, err = stmt.Exec(todo.ID)

	return err
}



func (todo Todo) Update() error {
	fmt.Println(todo)
	query := `
	UPDATE todos
	SET text = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {return nil}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Text, todo.ID)
	return err
}