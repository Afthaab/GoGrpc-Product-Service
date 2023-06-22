package usecase

import (
	interfaces "github.com/product/pkg/repository/interface"
	service "github.com/product/pkg/usecase/interface"
)

type productUseCase struct {
	Repo interfaces.ProductRepository
}

func NewProductUseCase(repo interfaces.ProductRepository) service.ProductUseCase {
	return &productUseCase{
		Repo: repo,
	}
}
