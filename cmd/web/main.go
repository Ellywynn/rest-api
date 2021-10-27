package main

import (
	"log"
	"net/http"

	"github.com/ellywynn/rest-api/cmd/web/routes"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.IndexPage)
	mux.HandleFunc("/course", routes.GetAllCourses)

	const PORT = ":4000"

	log.Println("Starting server on localhost" + PORT)
	err := http.ListenAndServe(PORT, mux)
	log.Fatal(err)
}
