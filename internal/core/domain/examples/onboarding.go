package domain

import "errors"

// OnboardingRequest represents the core entity
type OnboardingRequest struct {
	ID             string
	CustomerName   string
	DocumentStatus string
}

var ErrInvalidDocument = errors.New("document validation failed")

// OnboardingRepository defines how we interact with data storage
type OnboardingRepository interface {
	Save(req *OnboardingRequest) error
	FindByID(id string) (*OnboardingRequest, error)
}