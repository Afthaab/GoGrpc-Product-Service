package repository

import (
	"context"

	"github.com/product/pkg/domain"
	interfaces "github.com/product/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userDataBase struct {
	DB *mongo.Database
}

func (r *userDataBase) AddSizeAndPriceRnage(ctx context.Context, sizeData domain.Size) (domain.Size, error) {
	collection := r.DB.Collection("size")
	result, err := collection.InsertOne(ctx, sizeData)
	if err != nil {
		return sizeData, err
	}
	sizeData.Id = result.InsertedID.(primitive.ObjectID).Hex()
	return sizeData, nil
}

func (r *userDataBase) AddCategory(ctx context.Context, categorydata domain.Category) (domain.Category, error) {
	collection := r.DB.Collection("category")
	result, err := collection.InsertOne(ctx, categorydata)
	if err != nil {
		return categorydata, err
	}
	categorydata.Id = result.InsertedID.(primitive.ObjectID).Hex()
	return categorydata, nil
}

func (r *userDataBase) AddProduct(ctx context.Context, productdata domain.Products) (domain.Products, error) {
	collection := r.DB.Collection("products")
	result, err := collection.InsertOne(ctx, productdata)
	if err != nil {
		return productdata, err
	}
	productdata.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return productdata, nil
}

func (r *userDataBase) AddFoodType(ctx context.Context, typeData domain.Foodtype) (domain.Foodtype, error) {
	collection := r.DB.Collection("foodtype")
	result, err := collection.InsertOne(ctx, typeData)
	if err != nil {
		return typeData, err
	}
	typeData.Id = result.InsertedID.(primitive.ObjectID).Hex()
	return typeData, nil
}

func (r *userDataBase) ViewFoodType(ctx context.Context) ([]domain.Foodtype, error) {
	collection := r.DB.Collection("foodtype")

	var typeDatas []domain.Foodtype

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return typeDatas, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var typeData domain.Foodtype
		if err := cursor.Decode(&typeData); err != nil {
			return typeDatas, err
		}
		typeDatas = append(typeDatas, typeData)
	}

	// Check for any errors during cursor iteration
	if err := cursor.Err(); err != nil {
		return typeDatas, err
	}

	// Return the retrieved documents
	return typeDatas, nil

}

func (r *userDataBase) ViewSizes(ctx context.Context) ([]domain.Size, error) {
	collection := r.DB.Collection("size")

	// Define a slice to store the retrieved documents
	var sizes []domain.Size

	// Query the collection and get a cursor to iterate over the results
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return sizes, err
	}

	// Iterate over the cursor to fetch documents
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var size domain.Size
		if err := cursor.Decode(&size); err != nil {
			return sizes, err
		}
		sizes = append(sizes, size)
	}

	// Check for any errors during cursor iteration
	if err := cursor.Err(); err != nil {
		return sizes, err
	}

	// Return the retrieved documents
	return sizes, nil
}
func (r *userDataBase) ViewCategories(ctx context.Context) ([]domain.Category, error) {
	collection := r.DB.Collection("category")

	var categories []domain.Category

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return categories, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var category domain.Category
		if err := cursor.Decode(&category); err != nil {
			return categories, err
		}

		categories = append(categories, category)

	}

	if err := cursor.Err(); err != nil {
		return categories, err
	}

	return categories, nil
}

func (r *userDataBase) ViewProducts(ctx context.Context) ([]domain.Products, error) {
	collection := r.DB.Collection("products")

	var products []domain.Products

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return products, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product domain.Products
		if err := cursor.Decode(&product); err != nil {
			return products, err
		}
		products = append(products, product)
	}
	if err := cursor.Err(); err != nil {
		return products, err
	}

	return products, nil
}

func (r *userDataBase) ViewProductById(ctx context.Context, productData domain.Products) (domain.Products, error) {
	collection := r.DB.Collection("products")

	// string to primitive.ObjectID
	pid, _ := primitive.ObjectIDFromHex(productData.ID)

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": pid}

	err := collection.FindOne(ctx, filter).Decode(&productData)
	if err != nil {
		return productData, err
	}
	return productData, nil
}

func (r *userDataBase) ViewSizeById(ctx context.Context, id string) (domain.Size, error) {
	collection := r.DB.Collection("size")

	sizeData := domain.Size{}

	// string to primitive.ObjectID
	sid, _ := primitive.ObjectIDFromHex(id)

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": sid}

	err := collection.FindOne(ctx, filter).Decode(&sizeData)
	if err != nil {
		return sizeData, err
	}
	return sizeData, nil

}

func (r *userDataBase) ViewCategoryById(ctx context.Context, categoryid string) (domain.Category, error) {
	collection := r.DB.Collection("category")

	// string to primitive.ObjectID
	cid, _ := primitive.ObjectIDFromHex(categoryid)

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": cid}

	categoryData := domain.Category{}
	err := collection.FindOne(ctx, filter).Decode(&categoryData)
	if err != nil {
		return categoryData, err
	}
	return categoryData, nil
}

func (r *userDataBase) ViewTypeById(ctx context.Context, typeid string) (domain.Foodtype, error) {
	collection := r.DB.Collection("foodtype")

	// string to primitive.ObjectID
	tid, _ := primitive.ObjectIDFromHex(typeid)

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": tid}

	typeData := domain.Foodtype{}
	err := collection.FindOne(ctx, filter).Decode(&typeData)
	if err != nil {
		return typeData, err
	}
	return typeData, nil

}

func NewProductRepo(db *mongo.Database) interfaces.ProductRepository {
	return &userDataBase{
		DB: db,
	}
}
