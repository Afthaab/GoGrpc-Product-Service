package interfaces

import (
	"context"

	domain "github.com/product/pkg/domain"
)

type ProductUseCase interface {
	AddSizeBasedPrices(ctx context.Context, sizseData domain.Size) (domain.Size, error)
}
