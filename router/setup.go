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

func setupListRoutes(v *gin.RouterGroup, h *handler.BookHadler) {
	v.POST("/lists")

	v.DELETE("/lists")

	v.DELETE("/list/:id/book/:id")

	v.GET("/list")

	v.GET("/searchList/:id/book/:id")
}
