package handler

import (
	"github.com/product/pkg/pb"
	service "github.com/product/pkg/usecase/interface"
)

type ProductHandler struct {
	productUsecase service.ProductUseCase
	pb.ProductManagementServer
}

func NewProductHandler(service service.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUsecase: service,
	}
}
