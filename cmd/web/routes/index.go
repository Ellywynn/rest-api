package routes

import (
	"html/template"
	"log"
	"net/http"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("./public/pages/index.tmpl")

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}

	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
	}

	w.Write([]byte("Hello there"))
}
