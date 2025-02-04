package main

import (
	"context"
	handlers "file_rest_api/internal/handler"
	"file_rest_api/internal/repository"
	"file_rest_api/internal/service"
	"file_rest_api/pkg/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	database := db.Connect()
	defer database.Close(context.Background())

	repo := repository.NewRepository(database)
	service := service.NewService(repo)
	h := handlers.NewHandler(service)

	router := gin.Default()
	router.POST("/deposit", h.Deposit)
	router.POST("/transfer", h.Transfer)
	router.GET("/transactions/:user_id", h.GetLastTransactions)

	log.Println("Starting the server on the port 8080")

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalf("Error when starting the server: %v", err)
	}

}
