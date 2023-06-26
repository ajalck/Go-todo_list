package server

import (
	"log"
	"sync"

	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

func InitServer(wg *sync.WaitGroup) {
	defer wg.Done()
	r := gin.Default()

	if err := r.Run(port); err != nil {
		log.Fatal("Failed to initiate server")
	}

}
