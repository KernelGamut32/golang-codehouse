package routes

import (
	todosService "github.com/KernelGamut32/golang-codehouse/exercises/gin/backend/internal/todos/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/api/lists", todosService.Lists)
	r.GET("/api/lists/:index", todosService.ListItem)
	r.POST("/api/lists", todosService.AddListItem)
}
