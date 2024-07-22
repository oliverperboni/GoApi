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

	listRepository := repository.CreateListRepository(config.GetDB())
	listService := services.CreateListService(listRepository)
	listaHanlder := handler.CreateNewListHandler(listService)

	bookRepository := repository.CreateBookRepository(config.GetDB())
	bookService := services.CreateBookService(bookRepository)
	bookHandler := handler.CreateBookHandler(bookService)

	router.Initialize(&bookHandler, &listaHanlder)

}
