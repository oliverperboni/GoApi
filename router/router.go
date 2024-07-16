package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliverperboni/GoApi/handler"
)

func Initialize(h *handler.BookHadler) {

	r := gin.Default()

	initializeRouter(r, h)

	r.Run() // listen and serve on 0.0.0.0:8080

}
