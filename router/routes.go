package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoApi/handler"
)

func initializeRouter(r *gin.Engine) {

	v1 := r.Group("/api/v1/")
	{
		v1.GET("/book", handler.GetBooksHandler)

		v1.POST("/book", handler.PostBooksHandler)

		v1.PUT("/book", handler.PutBooksHandler)

		v1.DELETE("/book", handler.DeleteBooksHandler)

	}
}
