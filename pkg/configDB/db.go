package configDB

import (
	"log"

	"github.com/ajalck/todo_list/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(ch chan *gorm.DB) {
	db, err := gorm.Open(postgres.Open("postgres://ajalck:ack12345@localhost:5432/todo_list"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB")
	}

	if err := db.AutoMigrate(&model.Todo{}); err != nil {
		log.Fatal("Failed to sync `ToDo` model")
	}
	log.Println("Connected to DB successfully")
	ch <- db
}
