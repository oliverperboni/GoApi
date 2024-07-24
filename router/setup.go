package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoApi/handler"
)

func setupBookRoutes(v *gin.RouterGroup, h *handler.BookHadler) {
	v.GET("/book", h.GetBook)

	v.GET("/book/id/:id", h.GetBookById)

	v.GET("/book/name/:name", h.GetBookByName)

	v.POST("/book", h.PostBook)

	v.PUT("/book", h.PutBook)

	v.DELETE("/book", h.DeleteBook)
}

func setupListRoutes(v *gin.RouterGroup, h *handler.ListHandler) {

	//* recebe uma list
	v.POST("/lists", h.PostList)
	//* recebe um bookID e listID
	v.POST("/lists/AddBook", h.PostBookList)

	//* recebe uma list
	v.DELETE("/lists", h.DeleteList)

	//* recebe um bookID e listID
	v.DELETE("/list/:id/book/:id", h.DeleteBookList)

	//* recebe um bookID e listID
	v.GET("/list/book", h.GetBookList)

	//* recebe um userID e um listID
	v.GET("/list/user", h.GetAllUserList)

	//* recebe um userID e um listID
	v.GET("/list/allBooks", h.GetAllBookList)
}
