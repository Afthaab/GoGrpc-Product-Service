package api

import (
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/product/pkg/api/handler"
	"github.com/product/pkg/pb"
	"google.golang.org/grpc"
)

type ServerHttp struct {
	Engine *gin.Engine
}

func NewGRPCServer(productHandler *handler.ProductHandler, grpcPort string) {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalln("Failed to listen to GRPC Port", err)
	}

	//creating a new GRPC Server
	grpcServer := grpc.NewServer()

	pb.RegisterProductManagementServer(grpcServer, productHandler)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("Could not servre the grpc Server", err)
	}

}

func NewServerHttp(productHandler *handler.ProductHandler) *ServerHttp {
	engine := gin.New()

	engine.Use(gin.Logger())

	go NewGRPCServer(productHandler, "8891")

	return &ServerHttp{
		Engine: engine,
	}
}

func (s *ServerHttp) Start() {
	s.Engine.Run(":7779")
}
