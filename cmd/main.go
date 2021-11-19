package main

import (
	"fmt"
	"log"

	restapi "github.com/ellywynn/rest-api"
	"github.com/ellywynn/rest-api/pkg/handler"
	"github.com/ellywynn/rest-api/pkg/repository"
	"github.com/ellywynn/rest-api/pkg/service"
	"github.com/spf13/viper"
)

const port = "3000"

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("An erroc occurred while initializing config: %s\n", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	s := new(restapi.Server)
	if err := s.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("An error occurred while running the server: %s\n", err)
	}

	fmt.Println("Server running on port " + port)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
