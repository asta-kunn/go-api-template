package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/asta-kunn/go-api-template/internal/core/domain"
	"github.com/asta-kunn/go-api-template/internal/core/usecases"
)

type OnboardingHandler struct {
	usecase *usecases.OnboardingUsecase
}

func NewOnboardingHandler(uc *usecases.OnboardingUsecase) *OnboardingHandler {
	return &OnboardingHandler{usecase: uc}
}

func (h *OnboardingHandler) Submit(c *gin.Context) {
	var req domain.OnboardingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := h.usecase.ProcessOnboarding(&req)
	if err != nil {
		if err == domain.ErrInvalidDocument {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Onboarding successful", "data": req})
}