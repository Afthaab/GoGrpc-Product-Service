package interfaces

import (
	"context"

	"github.com/product/pkg/domain"
)

type ProductRepository interface {
	AddSizeAndPriceRnage(ctx context.Context, sizeData domain.Size) (domain.Size, error)
	ViewSizes(ctx context.Context) ([]domain.Size, error)
	ViewSizeById(ctx context.Context, id string) (domain.Size, error)
	EditSizeBasePrice(ctx context.Context, sizeData domain.Size) error
	DeleteSizeBasedPrize(ctx context.Context, sizeData domain.Size) error

	AddProduct(ctx context.Context, productdata domain.Products) (domain.Products, error)
	ViewProducts(ctx context.Context) ([]domain.Products, error)
	ViewProductById(ctx context.Context, productData domain.Products) (domain.Products, error)
	DeleteProduct(ctx context.Context, productData domain.Products) error

	AddCategory(ctx context.Context, categorydata domain.Category) (domain.Category, error)
	ViewCategories(ctx context.Context) ([]domain.Category, error)
	EditCategory(ctx context.Context, categoryData domain.Category) error
	ViewCategoryById(ctx context.Context, categoryid string) (domain.Category, error)
	DeleteCategory(ctx context.Context, categoryData domain.Category) error

	AddFoodType(ctx context.Context, typeData domain.Foodtype) (domain.Foodtype, error)
	EditFoodType(ctx context.Context, typeData domain.Foodtype) error
	ViewFoodType(ctx context.Context) ([]domain.Foodtype, error)
	ViewTypeById(ctx context.Context, typeid string) (domain.Foodtype, error)
	DeleteFoodType(ctx context.Context, typeData domain.Foodtype) error
}
