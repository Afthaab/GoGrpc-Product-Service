package interfaces

import (
	"context"

	"github.com/product/pkg/domain"
)

type ProductRepository interface {
	AddSizeAndPriceRnage(ctx context.Context, sizeData domain.Size) (domain.Size, error)
}
