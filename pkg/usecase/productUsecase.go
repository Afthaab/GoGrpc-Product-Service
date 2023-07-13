package usecase

import (
	"context"
	"errors"
	"strconv"

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

func (u *productUseCase) AddCategory(ctx context.Context, categoryData domain.Category) (domain.Category, error) {
	categoryData, err := u.Repo.AddCategory(ctx, categoryData)
	if err != nil {
		return domain.Category{}, err
	}
	return domain.Category{}, nil
}

func (u *productUseCase) AddProduct(ctx context.Context, productData domain.Products) (domain.Products, error) {
	productData, err := u.Repo.AddProduct(ctx, productData)
	if err != nil {
		return productData, err
	}
	return productData, nil
}

func (u *productUseCase) ViewSizeBasedPrice(ctx context.Context) ([]domain.Size, error) {
	sizes, err := u.Repo.ViewSizes(ctx)
	if err != nil {
		return sizes, err
	}
	return sizes, nil
}

func (u *productUseCase) ViewCategories(ctx context.Context) ([]domain.Category, error) {
	categories, err := u.Repo.ViewCategories(ctx)
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func (u *productUseCase) ViewProducts(ctx context.Context) ([]domain.Products, error) {
	products, err := u.Repo.ViewProducts(ctx)
	if err != nil {
		return products, err
	}
	return products, nil
}
func (u *productUseCase) AddFoodType(ctx context.Context, typeData domain.Foodtype) (domain.Foodtype, error) {
	typeData, err := u.Repo.AddFoodType(ctx, typeData)
	if err != nil {
		return typeData, err
	}
	return typeData, nil
}

func (u *productUseCase) ViewFoodType(ctx context.Context) ([]domain.Foodtype, error) {
	typeData, err := u.Repo.ViewFoodType(ctx)
	if err != nil {
		return typeData, err
	}
	return typeData, nil
}

func (u *productUseCase) ViewProductById(ctx context.Context, productData domain.Products) (domain.Products, error) {
	productData, err := u.Repo.ViewProductById(ctx, productData)
	if err != nil {
		return productData, err
	}
	return productData, nil
}

func (u *productUseCase) ViewTypeById(ctx context.Context, typeid string) (domain.Foodtype, error) {
	typeData, err := u.Repo.ViewTypeById(ctx, typeid)
	if err != nil {
		return typeData, err
	}
	return typeData, nil
}

func (u *productUseCase) ViewCategoryById(ctx context.Context, catId string) (domain.Category, error) {
	categoryData, err := u.Repo.ViewCategoryById(ctx, catId)
	if err != nil {
		return categoryData, err
	}
	return categoryData, nil
}

func (u *productUseCase) CalculatePrice(ctx context.Context, sizeData []string, basePrice float64) ([]domain.Size, error) {
	sliceSizeData := []domain.Size{}
	for _, id := range sizeData {
		sizeDatas := domain.Size{}
		sizeDatas, err := u.Repo.ViewSizeById(ctx, id)
		if err != nil {
			return sliceSizeData, err
		}
		// setting the price by comparing the percentage of the size
		price, _ := strconv.ParseFloat(sizeDatas.Price, 64)
		newPrice := (price * basePrice) / 100
		sizeDatas.Price = strconv.FormatFloat(newPrice, 'f', -1, 64)

		sliceSizeData = append(sliceSizeData, sizeDatas)
	}
	return sliceSizeData, nil
}

func (u *productUseCase) EditCategory(ctx context.Context, categoryData domain.Category) error {
	err := u.Repo.EditCategory(ctx, categoryData)
	if err != nil {
		return errors.New("Could not update the details")
	}
	return nil
}

func (u *productUseCase) EditFoodType(ctx context.Context, typeData domain.Foodtype) error {
	err := u.Repo.EditFoodType(ctx, typeData)
	if err != nil {
		return errors.New("Could not update the food type")
	}
	return nil
}

func (u *productUseCase) EditSizeBasedPrice(ctx context.Context, sizeData domain.Size) error {
	err := u.Repo.EditSizeBasePrice(ctx, sizeData)
	if err != nil {
		return errors.New("Could not edit the size")
	}
	return nil
}
func (u *productUseCase) DeleteSizeBasedPrize(ctx context.Context, sizeData domain.Size) error {
	err := u.Repo.DeleteSizeBasedPrize(ctx, sizeData)
	if err != nil {
		return err
	}
	return nil
}

func (u *productUseCase) DeleteCategory(ctx context.Context, categoryData domain.Category) error {
	err := u.Repo.DeleteCategory(ctx, categoryData)
	if err != nil {
		return err
	}
	return nil
}

func (u *productUseCase) DeleteFoodType(ctx context.Context, typeData domain.Foodtype) error {
	err := u.Repo.DeleteFoodType(ctx, typeData)
	if err != nil {
		return err
	}
	return nil
}
func NewProductUseCase(repo interfaces.ProductRepository) service.ProductUseCase {
	return &productUseCase{
		Repo: repo,
	}
}
