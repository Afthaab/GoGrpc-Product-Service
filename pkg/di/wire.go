//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	server "github.com/product/pkg/api"
	handler "github.com/product/pkg/api/handler"
	config "github.com/product/pkg/config"
	db "github.com/product/pkg/db"
	repository "github.com/product/pkg/repository"
	usecase "github.com/product/pkg/usecase"
)

func InitializeApi(cfg config.Config) (*server.ServerHttp, error) {
	wire.Build(
		db.ConnectToDatabase,
		repository.NewProductRepo,
		usecase.NewProductUseCase,
		handler.NewProductHandler,
		server.NewServerHttp)
	return &server.ServerHttp{}, nil

	//go run github.com/google/wire/cmd/wire
}
