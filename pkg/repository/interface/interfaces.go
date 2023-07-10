package interfaces

import (
	"context"

	"github.com/product/pkg/domain"
)

type ProductRepository interface {
	AddSizeAndPriceRnage(ctx context.Context, sizeData domain.Size) (domain.Size, error)
	AddCategory(ctx context.Context, categorydata domain.Category) (domain.Category, error)
	AddProduct(ctx context.Context, productdata domain.Products) (domain.Products, error)
	ViewCategories(ctx context.Context) ([]domain.Category, error)
	ViewSizes(ctx context.Context) ([]domain.Size, error)
	ViewProducts(ctx context.Context) ([]domain.Products, error)
	AddFoodType(ctx context.Context, typeData domain.Foodtype) (domain.Foodtype, error)
	ViewFoodType(ctx context.Context) ([]domain.Foodtype, error)
	ViewProductById(ctx context.Context, productData domain.Products) (domain.Products, error)
	ViewSizeById(ctx context.Context, id string) (domain.Size, error)
	ViewCategoryById(ctx context.Context, categoryid string) (domain.Category, error)
	ViewTypeById(ctx context.Context, typeid string) (domain.Foodtype, error)
}
