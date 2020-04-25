package main

import (
	"github.com/CatCanCreate/gomeetup/internal/handler"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	h := handler.NewHandler()

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Println("server start on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	log.Fatalln(err)

}
