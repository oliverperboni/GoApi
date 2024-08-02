package main

import (
	"github.com/oliverperboni/GoTomekeeper/config"
	"github.com/oliverperboni/GoTomekeeper/handler"
	"github.com/oliverperboni/GoTomekeeper/repository"
	"github.com/oliverperboni/GoTomekeeper/router"
	"github.com/oliverperboni/GoTomekeeper/services"
)

func main() {

	config.ConnectDatabase()

	listRepository := repository.CreateListRepository(config.GetDB())
	listService := services.CreateListService(listRepository)
	listaHanlder := handler.CreateNewListHandler(listService)

	reviewRepo := repository.CreteNewReviewRepo(config.DB)
	reviewService := services.CreateReviewService(reviewRepo)
	reviewHandler := handler.CreateReviewHandler(reviewService)

	bookRepository := repository.CreateBookRepository(config.GetDB())
	bookService := services.CreateBookService(bookRepository)
	bookHandler := handler.CreateBookHandler(bookService)

	router.Initialize(&bookHandler, &listaHanlder, &reviewHandler)

}
