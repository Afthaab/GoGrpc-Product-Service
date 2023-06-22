package main

import (
	"fmt"
	"log"

	"github.com/product/pkg/config"
	"github.com/product/pkg/di"
)

func main() {
	cfg, cfgErr := config.LoadConfig()
	if cfgErr != nil {
		log.Fatalln("Could not load the config file :", cfgErr)
		return
	}
	server, err := di.InitializeApi(cfg)
	if err != nil {
		log.Fatalln("Error in initializing the API", err)
	}
	fmt.Println("Server Started at 7779")
	server.Start()
}
