package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoApi/handler"
)

func initializeRouter(r *gin.Engine) {

	v1 := r.Group("/api/v1/")
	{
		v1.GET("/ping", handler.GetOpeningHandler)

		v1.POST("/ping", handler.PostOpeningHandler)

		v1.PUT("/ping", handler.PutOpeningHandler)

		v1.DELETE("/ping", handler.DeleteOpeningHandler)

	}
}
