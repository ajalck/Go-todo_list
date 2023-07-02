package di

import (
	"github.com/ajalck/todo_list/pkg/app/handler"
	"github.com/ajalck/todo_list/pkg/app/repository"
	"gorm.io/gorm"
)

//	func InitializeTodo(wg *sync.WaitGroup) handler.Handler {
//		defer wg.Done()
//		log.Println("fun called................")
//		wire.Build(
//			handler.NewHandler,
//			repository.NewRepo,
//			configDB.ConnectDB,
//		)
//		return handler.Handler{}
//	}

func InitializeTodo(db *gorm.DB, hch chan *handler.Handler) {
	var (
		repo    repository.Repo  = repository.NewRepo(db)
		handler *handler.Handler = handler.NewHandler(repo)
	)
	hch <- handler
}
