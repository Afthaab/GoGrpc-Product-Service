package db

import (
	"context"
	"fmt"

	"github.com/product/pkg/config"
	utility "github.com/product/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDatabase(cfg config.Config) (*mongo.Database, error) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	url := fmt.Sprintf("mongodb+srv://mongodb:%s@cluster1.kfqqheq.mongodb.net/?retryWrites=true&w=majority", cfg.DBPassword)
	opts := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	databseConn := client.Database("microsvc")

	collectionNames, err := databseConn.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("failed to list collection names: %v", err)
	}

	if !utility.Contains(collectionNames, "category") {
		err = databseConn.CreateCollection(ctx, "category")
		if err != nil {
			return nil, fmt.Errorf("failed to create 'category' collection: %v", err)
		}
	}

	if !utility.Contains(collectionNames, "size") {
		err = databseConn.CreateCollection(ctx, "size")
		if err != nil {
			return nil, fmt.Errorf("failed to create 'size' collection: %v", err)
		}
	}
	if !utility.Contains(collectionNames, "products") {
		err = databseConn.CreateCollection(ctx, "products")
		if err != nil {
			return nil, fmt.Errorf("failed to create 'products' collection: %v", err)
		}
	}
	if !utility.Contains(collectionNames, "foodtype") {
		err = databseConn.CreateCollection(ctx, "foodtype")
		if err != nil {
			return nil, fmt.Errorf("failed to create 'foodtype' collection: %v", err)
		}
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return databseConn, err
}
