package repository

import (
	"github.com/ajalck/todo_list/pkg/model"
)

type Repo interface {
	FetchTodo(page, limit int) ([]model.Todo,interface{}, error)
	FetchTodoByID(id uint) (model.Todo, error)
	CreateTodo(todo model.Todo) error
	UpdateTodo(id uint,update model.Todo) error
	DeleteTodo(id uint) error
}
