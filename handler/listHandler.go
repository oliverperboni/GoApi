package handler


import "github.com/oliverperboni/GoApi/services"
import "github.com/gin-gonic/gin" 

type ListHandler struct {
	service services.ListService	
}
func CreateNewListHandler(s *services.ListService) ListHandler  {
	return ListHandler{service: *s}
}


func (l *ListHandler) GetList(ctx *gin.Context)  {
	
}

func (l *ListHandler) DeleteList(ctx *gin.Context)  {
	
}

func (l *ListHandler) PostList(ctx *gin.Context)  {
	
}

func (l *ListHandler) PutList(ctx *gin.Context)  {
	
}

func (l *ListHandler) PostBookList(ctx *gin.Context)  {
	
}

func (l *ListHandler) DeleteBookList(ctx *gin.Context)  {
	
}
