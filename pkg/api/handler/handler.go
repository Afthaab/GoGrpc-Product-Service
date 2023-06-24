package handler

import (
	"context"
	"net/http"

	domain "github.com/product/pkg/domain"
	"github.com/product/pkg/pb"
	service "github.com/product/pkg/usecase/interface"
)

type ProductHandler struct {
	productUsecase service.ProductUseCase
	pb.ProductManagementServer
}

func (h *ProductHandler) AddSizeBazedPrize(ctx context.Context, req *pb.AddSizeBazedPrizeRequest) (*pb.AddSizeBazedPrizeResponse, error) {
	sizeData := domain.Size{
		Name:  req.Name,
		Price: req.Pricerange,
	}
	sizeData, err := h.productUsecase.AddSizeBasedPrices(ctx, sizeData)
	if err != nil {
		return &pb.AddSizeBazedPrizeResponse{
			Status: http.StatusBadRequest,
			Error:  "Could not add the new size",
		}, err
	} else {
		return &pb.AddSizeBazedPrizeResponse{
			Status: http.StatusOK,
			Error:  "nil",
			Sizeid: sizeData.ID,
		}, nil
	}
}

func NewProductHandler(service service.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUsecase: service,
	}
}
