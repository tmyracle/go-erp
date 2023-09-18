package main

import (
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tmyracle/go-erp/api/handlers"
	"github.com/tmyracle/go-erp/database"
	"github.com/tmyracle/go-erp/models"
)

func setupRouter() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.InitDatabase()
	if err := database.DB.AutoMigrate(&models.Account{}); err != nil {
    log.Fatalf("Failed to auto-migrate: %v", err)
	}


	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	// Get accounts
	r.GET("/api/accounts", handlers.GetAccounts)
	r.POST("/api/accounts", handlers.CreateAccount)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
