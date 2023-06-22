package repository

import (
	interfaces "github.com/product/pkg/repository/interface"
	"go.mongodb.org/mongo-driver/mongo"
)

type userDataBase struct {
	DB *mongo.Client
}

func NewProductRepo(db *mongo.Client) interfaces.ProductRepository {
	return &userDataBase{
		DB: db,
	}
}
