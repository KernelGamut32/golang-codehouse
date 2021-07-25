package main

import (
	"github.com/KernelGamut32/golang-codehouse/exercises/gin/backend/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run()
}
