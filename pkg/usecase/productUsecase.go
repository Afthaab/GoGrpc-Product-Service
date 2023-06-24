package usecase

import (
	"context"

	"github.com/product/pkg/domain"
	interfaces "github.com/product/pkg/repository/interface"
	service "github.com/product/pkg/usecase/interface"
)

type productUseCase struct {
	Repo interfaces.ProductRepository
}

func (u *productUseCase) AddSizeBasedPrices(ctx context.Context, sizseData domain.Size) (domain.Size, error) {
	sizseData, err := u.Repo.AddSizeAndPriceRnage(ctx, sizseData)
	if err != nil {
		return sizseData, err
	}
	return sizseData, nil
}

func NewProductUseCase(repo interfaces.ProductRepository) service.ProductUseCase {
	return &productUseCase{
		Repo: repo,
	}
}
