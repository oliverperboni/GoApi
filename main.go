package main

import (
	"github.com/oliverperboni/GoApi/config"
	"github.com/oliverperboni/GoApi/handler"
	"github.com/oliverperboni/GoApi/repository"
	"github.com/oliverperboni/GoApi/router"
	"github.com/oliverperboni/GoApi/services"
)

func main() {

	config.ConnectDatabase()

	bookRepository := repository.CreateBookRepository(config.GetDB()	)
	bookService := services.CreateBookService(bookRepository) // Adjust as per your service structure
	bookHandler := handler.CreateBookHandler(bookService)     // Assuming a constructor function NewBookHandler

	router.Initialize(&bookHandler)

}
