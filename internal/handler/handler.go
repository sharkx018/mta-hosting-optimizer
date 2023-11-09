package handler

import (
	"context"
	"github.com/mta-hosting-optimizer/internal/entity"
	"github.com/mta-hosting-optimizer/internal/helper"
	"net/http"
)

func (h *Handler) GetInefficientMTAsHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	// Call the use case method
	response, err := h.inefficientUc.GetInactiveServers(ctx)
	if err != nil {
		rr := &entity.InefficentMTAResponseParams{
			Data: nil,
			Error: &entity.CommonErrorResponse{
				Message: err.Error(),
			},
		}
		helper.WriteCustomResp(w, 500, rr)
	}

	response.Error = nil
	helper.WriteCustomResp(w, http.StatusOK, response)

}
