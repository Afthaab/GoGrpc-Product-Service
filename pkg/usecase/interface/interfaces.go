package interfaces

import (
	"context"

	domain "github.com/product/pkg/domain"
)

type ProductUseCase interface {
	AddSizeBasedPrices(ctx context.Context, sizseData domain.Size) (domain.Size, error)
	AddCategory(ctx context.Context, CategoryData domain.Category) (domain.Category, error)
	AddProduct(ctx context.Context, productData domain.Products) (domain.Products, error)
	ViewProducts(ctx context.Context) ([]domain.Products, error)
	ViewCategories(ctx context.Context) ([]domain.Category, error)
	ViewSizeBasedPrice(ctx context.Context) ([]domain.Size, error)
	AddFoodType(ctx context.Context, typeData domain.Foodtype) (domain.Foodtype, error)
	ViewFoodType(ctx context.Context) ([]domain.Foodtype, error)
	ViewProductById(ctx context.Context, productData domain.Products) (domain.Products, error)
	CalculatePrice(ctx context.Context, sizeData []string, basePrice float64) ([]domain.Size, error)
	ViewCategoryById(ctx context.Context, categoryid string) (domain.Category, error)
	ViewTypeById(ctx context.Context, typeid string) (domain.Foodtype, error)
}
