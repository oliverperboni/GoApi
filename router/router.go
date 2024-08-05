package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoTomekeeper/handler"
)

func Initialize(h *handler.BookHadler, l *handler.ListHandler, rev *handler.ReviewHandler, u *handler.UserHandler) {

	r := gin.Default()

	initializeRouter(r, h, l, rev, u)

	r.Run() // listen and serve on 0.0.0.0:8080

}
