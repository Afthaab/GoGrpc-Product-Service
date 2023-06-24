package repository

import (
	"context"

	"github.com/product/pkg/domain"
	interfaces "github.com/product/pkg/repository/interface"
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
	sizeData.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return sizeData, nil
}

func NewProductRepo(db *mongo.Database) interfaces.ProductRepository {
	return &userDataBase{
		DB: db,
	}
}
