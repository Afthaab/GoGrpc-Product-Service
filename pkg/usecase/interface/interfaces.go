package interfaces

import (
	"context"

	domain "github.com/product/pkg/domain"
)

type ProductUseCase interface {
	AddSizeBasedPrices(ctx context.Context, sizseData domain.Size) (domain.Size, error)
	ViewSizeBasedPrice(ctx context.Context) ([]domain.Size, error)
	EditSizeBasedPrice(ctx context.Context, sizeData domain.Size) error
	DeleteSizeBasedPrize(ctx context.Context, sizeData domain.Size) error

	AddCategory(ctx context.Context, CategoryData domain.Category) (domain.Category, error)
	EditCategory(ctx context.Context, categoryData domain.Category) error
	ViewCategories(ctx context.Context) ([]domain.Category, error)
	ViewCategoryById(ctx context.Context, categoryid string) (domain.Category, error)
	DeleteCategory(ctx context.Context, categoryData domain.Category) error

	AddProduct(ctx context.Context, productData domain.Products) (domain.Products, error)
	ViewProducts(ctx context.Context) ([]domain.Products, error)
	ViewProductById(ctx context.Context, productData domain.Products) (domain.Products, error)

	CalculatePrice(ctx context.Context, sizeData []string, basePrice float64) ([]domain.Size, error)

	AddFoodType(ctx context.Context, typeData domain.Foodtype) (domain.Foodtype, error)
	EditFoodType(ctx context.Context, typeData domain.Foodtype) error
	ViewFoodType(ctx context.Context) ([]domain.Foodtype, error)
	ViewTypeById(ctx context.Context, typeid string) (domain.Foodtype, error)
	DeleteFoodType(ctx context.Context, typeData domain.Foodtype) error
}
