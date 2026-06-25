package repository

import (
	"github.com/asta-kunn/go-api-template/internal/core/domain"
)

type memoryRepo struct {
	data map[string]*domain.OnboardingRequest
}

func NewMemoryOnboardingRepository() domain.OnboardingRepository {
	return &memoryRepo{
		data: make(map[string]*domain.OnboardingRequest),
	}
}

func (m *memoryRepo) Save(req *domain.OnboardingRequest) error {
	m.data[req.ID] = req
	return nil
}

func (m *memoryRepo) FindByID(id string) (*domain.OnboardingRequest, error) {
	return m.data[id], nil
}