package usecase

import (
	"context"
	"github.com/mta-hosting-optimizer/internal/entity"
)

type Usecase struct {
	ipConfigRepo    IpConfigResource
	thresholdNumber int
}

// IpConfigResource Repo Interface
type IpConfigResource interface {
	GetIPConfigData(ctx context.Context) (entity.MockServiceResponse, error)
}

// New initialize Usecase object
func New(ipConfigRepo IpConfigResource, thresholdNumber int) *Usecase {
	return &Usecase{
		ipConfigRepo:    ipConfigRepo,
		thresholdNumber: thresholdNumber,
	}
}
