package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoApi/config"
	"github.com/oliverperboni/GoApi/handler"
	"github.com/oliverperboni/GoApi/repository"
	"github.com/oliverperboni/GoApi/services"
)

func Initialize() {

	r := gin.Default()
	bookRepository := repository.NewBookRepository(config.DB)
	bookService := services.NewBookService(bookRepository)
	bookController := handler.NewBookController(bookService)

	initializeRouter(r, bookController)

	r.Run() // listen and serve on 0.0.0.0:8080

}
