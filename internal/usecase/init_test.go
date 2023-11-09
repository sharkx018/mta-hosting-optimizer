package usecase

import (
	"github.com/mta-hosting-optimizer/internal/mock_gen"
	"testing"
)

func TestNewUsecase(t *testing.T) {
	// Create a mock IpConfigResource for testing
	mockRepo := &mock_gen.MockIpConfigResource{}

	// Call the New function to create a Usecase instance
	usecase := New(mockRepo, 5)

	// Assert that the Usecase instance is not nil
	if usecase == nil {
		t.Error("Expected non-nil Usecase instance, got nil")
	}

	// Assert that the fields are initialized correctly
	if usecase.ipConfigRepo != mockRepo {
		t.Errorf("Expected ipConfigRepo to be %v, got %v", mockRepo, usecase.ipConfigRepo)
	}

	if usecase.thresholdNumber != 5 {
		t.Errorf("Expected thresholdNumber to be 5, got %v", usecase.thresholdNumber)
	}
}
