package routes

import (
	ordersService "github.com/KernelGamut32/golang-codehouse/exercises/order-entry/backend/internal/orders/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/api/orders", ordersService.GetAll)
	r.GET("/api/orders/:id", ordersService.GetOrder)
	r.POST("/api/orders", ordersService.AddOrder)
	r.PUT("/api/orders/:id", ordersService.UpdateOrder)
}
