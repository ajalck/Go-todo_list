package repository

import (
	"errors"
	"math"
	"reflect"
	"time"

	"github.com/ajalck/todo_list/pkg/model"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{db}
}

func (r *repo) FetchTodo(page, limit int) ([]model.Todo, interface{}, error) {
	var totalRecords int64
	r.DB.Model(&model.Todo{}).Where("due > ?", time.Now()).Count(&totalRecords)
	if totalRecords == 0 {
		return nil, nil, errors.New("no records found")
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
	if int64(reflect.ValueOf(metadata).FieldByName("TotalPages").Float()) < int64(page) {
		return nil, nil, errors.New("record not found")
	}
	list := []model.Todo{}
	result := r.DB.Model(&model.Todo{}).Where("due > ?", time.Now()).Select("id", "title", "description", "due").Offset(offset).Limit(limit).Find(&list)
	if result.Error != nil {
		return nil, nil, errors.New("unable to fetch records")
	}
	return list, metadata, nil
}
func (r *repo) FetchTodoByID(id uint) (model.Todo, error) {
	todo := model.Todo{}
	result := r.DB.Where("id", id).Select("id", "created_at", "updated_at", "title", "description", "due").First(&todo)
	if result.Error != nil {
		return model.Todo{}, result.Error
	}
	return todo, nil
}
func (r *repo) CreateTodo(todo model.Todo) error {
	exTodo := model.Todo{}
	result := r.DB.Table("todos").Where("LOWER(title)=LOWER(?)", todo.Title).First(&exTodo)
	if result.Error == nil && time.Now().Before(exTodo.Due) {
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
	result := r.DB.Table("todos").Where("id", id).Updates(&model.Todo{
		Description: update.Description,
		Due:         update.Due,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *repo) DeleteTodo(id uint) error {
	result := r.DB.Table("todos").Delete(&model.Todo{ID: id})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
