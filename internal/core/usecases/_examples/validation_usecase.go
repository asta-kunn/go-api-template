package usecases
import (
	"github.com/asta-kunn/go-api-template/internal/core/domain"
)

type OnboardingUsecase struct {
	repo domain.OnboardingRepository
}

// Inject the repository via constructor
func NewOnboardingUsecase(repo domain.OnboardingRepository) *OnboardingUsecase {
	return &OnboardingUsecase{repo: repo}
}

// Centralized validation logic
func (uc *OnboardingUsecase) ProcessOnboarding(req *domain.OnboardingRequest) error {
	if req.DocumentStatus != "VERIFIED" {
		return domain.ErrInvalidDocument
	}
	
	// Dynamic logic or further validation can be centralized here
	
	return uc.repo.Save(req)
}