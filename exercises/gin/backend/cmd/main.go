package main

import (
	"github.com/KernelGamut32/golang-codehouse/exercises/gin/backend/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	routes.SetupRoutes(r)
	r.Run()
}
