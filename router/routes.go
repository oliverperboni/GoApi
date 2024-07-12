package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoApi/handler"
)

func initializeRouter(r *gin.Engine, b *handler.BookController) {

	v1 := r.Group("/api/v1/")
	{
		v1.GET("/book", b.GetBooks)

		v1.POST("/book", b.CreateBook)

		v1.PUT("/book", b.UpdateBook)

		v1.DELETE("/book", b.DeleteBook)

	}
}
