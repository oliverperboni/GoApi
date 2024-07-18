package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoApi/handler"
)

func initializeRouter(r *gin.Engine, h *handler.BookHadler) {

	v1 := r.Group("/api/v1/")
	{
		setupBookRoutes(v1, h)

		setupListRoutes(v1, h)
	}
}
