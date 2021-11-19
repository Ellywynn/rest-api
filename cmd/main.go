package main

import (
	"fmt"
	"log"

	restapi "github.com/ellywynn/rest-api"
	"github.com/ellywynn/rest-api/pkg/handler"
)

const port = "3000"

func main() {
	handlers := new(handler.Handler)
	s := new(restapi.Server)
	if err := s.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("An error occurred while running the server: %s\n", err)
	}

	fmt.Println("Server running on port " + port)
}
