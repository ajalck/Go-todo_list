package repository

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/ajalck/todo_list/pkg/model"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	fmt.Println("NewRepo fn called", db)
	return &repo{db}
}

func (r *repo) FetchTodo(page, limit int) (model.Todo, interface{}, error) {
	var totalRecords int64
	r.DB.Model(&model.Todo{}).Count(&totalRecords)
	if totalRecords == 0 {
		return model.Todo{}, nil, errors.New("no records found")
	}
	offset, metadata := func(currentPage int) (int, interface{}) {
		offset := (currentPage - 1) * limit
		totalPages := math.Ceil(float64(totalRecords) / float64(limit))
		metadata := struct {
			Page         int
			Limit        int
			TotalRecords int64
			TotalPages   float64
		}{
			Page:         currentPage,
			Limit:        limit,
			TotalRecords: totalRecords,
			TotalPages:   totalPages,
		}
		return offset, metadata
	}(page)
	list := model.Todo{}
	result := r.DB.Model(&model.Todo{}).Select("id", "title", "description", "due").Offset(offset).Limit(limit).Find(&list)
	if result.Error != nil {
		return model.Todo{}, nil, errors.New("unable to fetch records")
	}
	return list, metadata, nil
}
func (r *repo) FetchTodoByID(id uint) (model.Todo, error) {
	return model.Todo{}, nil
}
func (r *repo) CreateTodo(todo model.Todo) error {
	exTodo := model.Todo{}
	result := r.DB.Model("todos").Where("lower(title)=?", todo.Title).First(&exTodo)
	if result.Error == nil && time.Now().After(exTodo.Due) {
		return errors.New("todo exists")
	}
	result = r.DB.Create(&model.Todo{
		Title:       todo.Title,
		Description: todo.Description,
		Due:         todo.Due,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *repo) UpdateTodo(id uint, update model.Todo) error {
	return nil
}
func (r *repo) DeleteTodo(id uint) error {
	return nil
}
