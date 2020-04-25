package main

import (
	"github.com/CatCanCreate/gomeetup/internal/config"
	"github.com/CatCanCreate/gomeetup/internal/handler"
	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/http"
)

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	h := handler.NewHandler()

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	addr := cfg.Host + ":" + cfg.Port

	log.Println("server start on http://localhost:8080")
	err = http.ListenAndServe(addr, r)
	log.Fatalln(err)

}
