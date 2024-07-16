package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoApi/handler"
)

func initializeRouter(r *gin.Engine, h *handler.BookHadler) {

	v1 := r.Group("/api/v1/")
	{
		v1.GET("/book")

		v1.POST("/book", h.PostBook)

		v1.PUT("/book")

		v1.DELETE("/book")

	}
}
