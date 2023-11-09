package handler

import (
	"context"
	"github.com/mta-hosting-optimizer/internal/entity"
)

type InefficentMTAsUsecase interface {
	GetInactiveServers(ctx context.Context) (entity.InefficentMTAResponseParams, error)
}

// Handler defines the handler
type Handler struct {
	inefficientUc InefficentMTAsUsecase
}

// New creates handler
func New(inefficientUc InefficentMTAsUsecase) *Handler {
	return &Handler{
		inefficientUc: inefficientUc,
	}
}
