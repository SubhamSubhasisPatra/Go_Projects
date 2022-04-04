package main

import (
	"github.com/gin-gonic/gin"
	routers "go_jwt/routes"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routers.AuthRoutes(router)
	routers.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Hi From API-1"})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Hi From API-2"})
	})

	err := router.Run(":" + port)
	if err != nil {
		return
	}
}
