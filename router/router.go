package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoApi/handler"
)

func Initialize(h *handler.BookHadler, l *handler.ListHandler) {

	r := gin.Default()

	initializeRouter(r, h, l)

	r.Run() // listen and serve on 0.0.0.0:8080

}
