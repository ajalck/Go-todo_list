package main

import (
	di "github.com/ajalck/todo_list/pkg/DI"
	"github.com/ajalck/todo_list/pkg/app/handler"
	"github.com/ajalck/todo_list/pkg/configDB"
	"github.com/ajalck/todo_list/pkg/server"
	"gorm.io/gorm"
)

func main() {
	DBch := make(chan *gorm.DB)
	handlerCH := make(chan *handler.Handler)
	serverCH := make(chan *server.Serve)
	go configDB.ConnectDB(DBch)
	go di.InitializeTodo(<-DBch, handlerCH)
	go server.InitServer(<-handlerCH, serverCH)
	serve := <-serverCH
	close(serverCH)
	serve.Start()
}
