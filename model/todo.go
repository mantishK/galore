package model

import (
	"time"

	"github.com/mantishK/galore/config"
)

type Todo struct {
	ID       int       `json:"id"`
	Content  string    `json:"content"`
	UserID   int       `json:"user_id"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func (t *Todo) Get() error {
	err := config.DB.QueryRow("SELECT * FROM todos WHERE id = $1", t.ID).Scan(&t.ID, &t.Content, &t.UserID, &t.Created, &t.Modified)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todo) GetUserTodos() ([]Todo, error) {
	todos := make([]Todo, 0)
	rows, err := config.DB.Query("SELECT * FROM todos WHERE user_id = $1 ORDER BY created", t.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.Created, &todo.Modified); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (t *Todo) Insert() error {
	t.Created = time.Now()
	t.Modified = time.Now()
	return insertTodo(t)
}

func (t *Todo) Update() error {
	t.Modified = time.Now()
	return updateTodo(t)
}

func (t *Todo) Delete() error {
	_, err := config.DB.Exec("DELETE FROM todos WHERE id = $1", t.ID)
	return err
}

func insertTodo(t *Todo) error {
	err := config.DB.QueryRow("INSERT INTO todos (content, user_id, created, modified )VALUES ($1,$2,$3,$4) returning id", t.Content, t.UserID, t.Created, t.Modified).Scan(&t.ID)
	if err != nil {
		return err
	}
	return nil
}

func updateTodo(t *Todo) error {
	_, err := config.DB.Exec("UPDATE todos SET content=$1, user_id=$2, created=$3, modified=$4 WHERE id = $5", t.Content, t.UserID, t.Created, t.Modified, t.ID)
	return err
}
