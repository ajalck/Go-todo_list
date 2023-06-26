package main

import (
	"sync"

	"github.com/ajalck/todo_list/pkg/configDB"
	"github.com/ajalck/todo_list/pkg/server"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go configDB.ConnectDB(wg)
	go server.InitServer(wg)
	wg.Wait()
}
