package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoTomekeeper/handler"
)

func initializeRouter(r *gin.Engine, h *handler.BookHadler, l *handler.ListHandler, rev *handler.ReviewHandler, u *handler.UserHandler) {

	v1 := r.Group("/api/v1/")
	{
		setupBookRoutes(v1, h)

		setupListRoutes(v1, l)

		setupReviewRoutes(v1, rev)

		setupUserRoutes(v1, u)
	}
}
