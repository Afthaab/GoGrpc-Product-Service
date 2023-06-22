package db

import (
	"context"
	"fmt"

	"github.com/product/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDatabase(cfg config.Config) (*mongo.Client, error) {
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
	databseConn.CreateCollection(ctx, "products")
	databseConn.CreateCollection(ctx, "mans")
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client, err
}
