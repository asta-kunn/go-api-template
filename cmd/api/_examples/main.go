package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/asta-kunn/go-api-template/internal/adapters/handlers"
	"github.com/asta-kunn/go-api-template/internal/adapters/repository"
	"github.com/asta-kunn/go-api-template/internal/core/usecases"
)

func main() {
	// 1. Initialize Repositories
	repo := repository.NewMemoryOnboardingRepository()

	// 2. Initialize Usecases (injecting repo)
	usecase := usecases.NewOnboardingUsecase(repo)

	// 3. Initialize HTTP Handlers (injecting usecase)
	handler := handlers.NewOnboardingHandler(usecase)

	// 4. Setup Router
	router := gin.Default()
	
	// Optional: Add global middleware for logging, CORS, or Prometheus metrics here
	
	api := router.Group("/api/v1")
	{
		api.POST("/onboarding", handler.Submit)
	}

	log.Println("Server starting on port 8080...")
	router.Run(":8080")
}