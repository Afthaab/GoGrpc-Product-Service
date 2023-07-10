package handler

import (
	"bytes"
	"context"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	domain "github.com/product/pkg/domain"
	"github.com/product/pkg/pb"
	service "github.com/product/pkg/usecase/interface"
	utility "github.com/product/pkg/utils"
)

type ProductHandler struct {
	productUsecase service.ProductUseCase
	pb.ProductManagementServer
}

func (h *ProductHandler) AddSizeBazedPrize(ctx context.Context, req *pb.AddSizeBazedPrizeRequest) (*pb.AddSizeBazedPrizeResponse, error) {
	sizeData := domain.Size{
		Name:  req.Name,
		Price: req.Pricerange,
	}
	sizeData, err := h.productUsecase.AddSizeBasedPrices(ctx, sizeData)
	if err != nil {
		return &pb.AddSizeBazedPrizeResponse{
			Status: http.StatusBadRequest,
			Error:  "Could not add the new size",
		}, err
	} else {
		return &pb.AddSizeBazedPrizeResponse{
			Status: http.StatusOK,
			Error:  "nil",
			Sizeid: sizeData.Id,
		}, nil
	}
}

func (h *ProductHandler) AddImage(ctx context.Context, req *pb.AddImageRequest) (*pb.AddImageResponse, error) {

	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
		// Add any necessary AWS credentials or session options
	})

	if err != nil {
		return &pb.AddImageResponse{
			Status: http.StatusInternalServerError,
			Error:  "Could not create a session",
		}, err
	}

	// Create an S3 service client
	s3Client := s3.New(sess)

	// Generate a unique filename for the image
	filename := utility.GenerateUniqueFilename()

	// Upload the image data to S3 bucket
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("images-for-deliveryapp"),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(req.ImageData),
	})
	if err != nil {
		return &pb.AddImageResponse{
			Status: http.StatusInternalServerError,
			Error:  "Could not upload file to S3 Bucket",
		}, err
	}

	// Get the URL of the uploaded image
	imageURL := s3Client.Endpoint + "/" + "images-for-deliveryapp" + "/" + filename

	return &pb.AddImageResponse{
		Status:   http.StatusOK,
		Error:    "nil",
		Imageurl: imageURL,
	}, nil

}

func (h *ProductHandler) AddCategories(ctx context.Context, req *pb.AddCategoriesRequest) (*pb.AddCategoriesResponse, error) {

	categoryData := domain.Category{
		Categoryname: req.CateggoryName,
		Imageurl:     req.Imageurl,
	}

	categoryData, err := h.productUsecase.AddCategory(ctx, categoryData)
	if err != nil {
		return &pb.AddCategoriesResponse{
			Status: http.StatusBadRequest,
			Error:  "Could not add the Category",
		}, err
	} else {
		return &pb.AddCategoriesResponse{
			Status:     http.StatusAccepted,
			Error:      "nil",
			Categoryid: categoryData.Id,
		}, nil
	}
}

func (h *ProductHandler) AddProduct(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	productData := domain.Products{
		Productname:  req.Name,
		Calories:     req.Calories,
		Availibility: req.Availibilty,
		Categoryid:   req.Categoryid,
		Typeid:       req.Typeid,
		Baseprice:    float64(req.Baseprice),
		Sizeid:       req.Sizeid,
		Imageurls:    req.Imagelink,
	}
	productData, err := h.productUsecase.AddProduct(ctx, productData)
	if err != nil {
		return &pb.AddProductResponse{
			Status:    http.StatusBadRequest,
			Error:     "Could not add the product",
			Productid: productData.ID,
		}, err
	} else {
		return &pb.AddProductResponse{
			Status:    http.StatusOK,
			Error:     "nil",
			Productid: productData.ID,
		}, nil
	}

}

func (h *ProductHandler) ViewSizeBasedPrize(ctx context.Context, req *pb.ViewSizeBasedPriceRequest) (*pb.ViewSizeBasedPriceRespose, error) {
	var sizes []domain.Size
	sizes, err := h.productUsecase.ViewSizeBasedPrice(ctx)
	if err != nil {
		return &pb.ViewSizeBasedPriceRespose{
			Status: http.StatusBadRequest,
			Error:  "Could not view the size based prices",
		}, err
	} else {

		var pbSizes []*pb.AddSizeBazedPrizeRequest
		for _, size := range sizes {
			pbSize := &pb.AddSizeBazedPrizeRequest{
				Sizeid:     size.Id,
				Name:       size.Name,
				Pricerange: size.Price,
			}
			pbSizes = append(pbSizes, pbSize)
		}

		return &pb.ViewSizeBasedPriceRespose{
			Status: http.StatusOK,
			Error:  "nil",
			Sizes:  pbSizes,
		}, nil
	}
}

func (h *ProductHandler) ViewCategories(ctx context.Context, req *pb.ViewCategoriesRequest) (*pb.ViewCategoriesResponse, error) {
	categories, err := h.productUsecase.ViewCategories(ctx)
	if err != nil {
		return &pb.ViewCategoriesResponse{
			Status: http.StatusBadRequest,
			Error:  "Could not view the Categories",
		}, err
	} else {
		var pbCategories []*pb.AddCategoriesRequest
		for _, category := range categories {
			pbCategory := &pb.AddCategoriesRequest{
				Id:            category.Id,
				CateggoryName: category.Categoryname,
				Imageurl:      category.Imageurl,
			}
			pbCategories = append(pbCategories, pbCategory)
		}

		return &pb.ViewCategoriesResponse{
			Status:     http.StatusOK,
			Error:      "nil",
			Categories: pbCategories,
		}, nil
	}
}

func (h *ProductHandler) ViewProduct(ctx context.Context, req *pb.ViewProductRequest) (*pb.ViewProductResponse, error) {
	products, err := h.productUsecase.ViewProducts(ctx)
	if err != nil {
		return &pb.ViewProductResponse{
			Status: http.StatusNotFound,
			Error:  "Could not view the products",
		}, err
	} else {
		var pbProducts []*pb.AddProductRequest
		for _, product := range products {
			pbProduct := &pb.AddProductRequest{
				Id:          product.ID,
				Name:        product.Productname,
				Calories:    product.Calories,
				Availibilty: product.Availibility,
				Categoryid:  product.Categoryid,
				Typeid:      product.Typeid,
				Baseprice:   float32(product.Baseprice),
				Sizeid:      product.Sizeid,
				Imagelink:   product.Imageurls,
			}
			pbProducts = append(pbProducts, pbProduct)
		}

		return &pb.ViewProductResponse{
			Status:   http.StatusOK,
			Error:    "nil",
			Products: pbProducts,
		}, nil
	}
}
func (h *ProductHandler) ViewProductById(ctx context.Context, req *pb.ViewProductByIdRequest) (*pb.ViewProductByIdResponse, error) {
	productData := domain.Products{
		ID: req.Pid,
	}
	productData, err := h.productUsecase.ViewProductById(ctx, productData)
	if err != nil {
		return &pb.ViewProductByIdResponse{
			Status: http.StatusNotFound,
			Error:  "Could not View the product by id",
		}, err
	}
	categoryData, err := h.productUsecase.ViewCategoryById(ctx, productData.Categoryid)
	if err != nil {
		return &pb.ViewProductByIdResponse{
			Status: http.StatusNotFound,
			Error:  "Could not get the category data",
		}, err
	}

	typeData, err := h.productUsecase.ViewTypeById(ctx, productData.Typeid)
	if err != nil {
		return &pb.ViewProductByIdResponse{
			Status: http.StatusNotFound,
			Error:  "Could not get the type data",
		}, err
	}

	sizeData, err := h.productUsecase.CalculatePrice(ctx, productData.Sizeid, productData.Baseprice)
	if err != nil {
		return &pb.ViewProductByIdResponse{
			Status: http.StatusBadRequest,
			Error:  "Could not caluculate the price",
		}, err
	}

	var sizes []*pb.AddSizeBazedPrizeRequest
	for _, sizeDatas := range sizeData {
		size := &pb.AddSizeBazedPrizeRequest{
			Sizeid:     sizeDatas.Id,
			Name:       sizeDatas.Name,
			Pricerange: sizeDatas.Price,
		}
		sizes = append(sizes, size)
	}

	return &pb.ViewProductByIdResponse{
		Status:       http.StatusOK,
		Error:        "nil",
		Pid:          productData.ID,
		Name:         productData.Productname,
		Calories:     productData.Calories,
		Categoryname: categoryData.Categoryname,
		Typename:     typeData.Foodtype,
		Sizeandprice: sizes,
		Availibilty:  productData.Availibility,
		Imagelink:    productData.Imageurls,
	}, nil

}

func (h *ProductHandler) ViewFoodType(ctx context.Context, req *pb.ViewFoodtypeRequest) (*pb.ViewFoodTypeResponse, error) {
	typeData, err := h.productUsecase.ViewFoodType(ctx)
	if err != nil {
		return &pb.ViewFoodTypeResponse{
			Status: http.StatusBadRequest,
			Error:  "Could not view the food type",
		}, err
	} else {
		var pbFoodtypes []*pb.AddFoodTypeRequest
		for _, foodtype := range typeData {
			pbFoodtype := &pb.AddFoodTypeRequest{
				Typeid:   foodtype.Id,
				Name:     foodtype.Foodtype,
				Imageurl: foodtype.Imageurl,
			}
			pbFoodtypes = append(pbFoodtypes, pbFoodtype)
		}
		return &pb.ViewFoodTypeResponse{
			Status:    http.StatusOK,
			Error:     "nil",
			Foodtypes: pbFoodtypes,
		}, nil
	}
}

func (h *ProductHandler) AddFoodType(ctx context.Context, req *pb.AddFoodTypeRequest) (*pb.AddFoodTypeResponse, error) {
	typeData := domain.Foodtype{
		Foodtype: req.Name,
		Imageurl: req.Imageurl,
	}
	typeData, err := h.productUsecase.AddFoodType(ctx, typeData)
	if err != nil {
		return &pb.AddFoodTypeResponse{
			Status: http.StatusBadRequest,
			Error:  "Could not add the Foodtype",
		}, err
	} else {
		return &pb.AddFoodTypeResponse{
			Status: http.StatusOK,
			Error:  "nil",
			Typeid: typeData.Id,
		}, nil
	}
}

func NewProductHandler(service service.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUsecase: service,
	}
}
