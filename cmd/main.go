package main

import (
	"fmt"
	"os"

	restapi "github.com/ellywynn/rest-api"
	"github.com/ellywynn/rest-api/pkg/handler"
	"github.com/ellywynn/rest-api/pkg/repository"
	"github.com/ellywynn/rest-api/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("An error occurred while initializing config: %s\n", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Can't load .env file: %s\n", err.Error())
	}

	db, err := repository.NewPostgres(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBUser:   viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("An error occurred while connection to database: %s\n", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	port := viper.GetString("port")

	s := new(restapi.Server)
	if err := s.Run(port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("An error occurred while running the server: %s\n", err)
	}

	fmt.Println("Server running on port " + port)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
