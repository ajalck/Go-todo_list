package handler

import (
	"net/http"
	"strconv"

	repo "github.com/ajalck/todo_list/pkg/app/repository"
	"github.com/ajalck/todo_list/pkg/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo repo.Repo
}

func NewHandler(repo repo.Repo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) FetchTodo(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	list, metadata, err := h.repo.FetchTodo(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result := struct {
		List     []model.Todo
		MetaData interface{}
	}{
		List:     list,
		MetaData: metadata,
	}
	c.JSON(200, result)
}
func (h *Handler) FetchTodoByID(c *gin.Context) {
	todoId, _ := strconv.Atoi(c.Query("todoId"))
	todo, err := h.repo.FetchTodoByID(uint(todoId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, todo)
}
func (h *Handler) CreateTodo(c *gin.Context) {
	body := model.Todo{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid inputs": err.Error()})
		return
	}
	if err := h.repo.CreateTodo(body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "Successfully created new todo")
}
func (h *Handler) UpdateTodo(c *gin.Context) {
	todoId, _ := strconv.Atoi(c.Query("todoId"))
	body := model.Todo{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inputs"})
		return
	}
	err := h.repo.UpdateTodo(uint(todoId), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "Successfully updated todo")
}
func (h *Handler) DeleteTodo(c *gin.Context) {
	todoId, _ := strconv.Atoi(c.Query("todoId"))
	err := h.repo.DeleteTodo(uint(todoId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "Successfully deleted todo")
}
