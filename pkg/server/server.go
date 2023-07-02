package server

import (
	"log"

	"github.com/ajalck/todo_list/pkg/app/handler"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

type Serve struct {
	e *gin.Engine
}

func InitServer(todoHandler *handler.Handler, s chan *Serve) {
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/home/:page/:limit", todoHandler.FetchTodo)
	r.GET("/fetchtodo/:todoId", todoHandler.FetchTodoByID)
	r.POST("/createtodo", todoHandler.CreateTodo)
	r.PATCH("/updatetodo/:todoId", todoHandler.UpdateTodo)
	r.DELETE("/deletetodo/:todoId", todoHandler.DeleteTodo)
	s <- &Serve{e: r}
}

func (s *Serve) Start() {
	if err := s.e.Run(port); err != nil {
		log.Fatal("Failed to start server")
	}
}
