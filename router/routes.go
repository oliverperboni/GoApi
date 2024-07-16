package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoApi/handler"
)

func initializeRouter(r *gin.Engine, h *handler.BookHadler) {

	v1 := r.Group("/api/v1/")
	{
		v1.GET("/book", h.GetBook)

		v1.GET("/book//id/:id", h.GetBookById)

		v1.GET("/book/name/:name", h.GetBookById)

		v1.POST("/book", h.PostBook)

		v1.PUT("/book", h.PutBook)

		v1.DELETE("/book", h.DeleteBook)

	}
}
