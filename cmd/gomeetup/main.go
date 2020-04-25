package main

import (
	"context"
	"github.com/CatCanCreate/gomeetup/internal/api/jokes"
	"github.com/CatCanCreate/gomeetup/internal/config"
	"github.com/CatCanCreate/gomeetup/internal/handler"
	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func gracefulShutdown(server *http.Server, quit <-chan os.Signal, done chan<- bool) {
	<-quit
	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	close(done)
}

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	apiClient := jokes.NewJokeClientAPI(cfg.JokeURL)

	h := handler.NewHandler(apiClient)

	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	addr := cfg.Host + ":" + cfg.Port

	quit := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go gracefulShutdown(server, quit, done)

	log.Println("server start on http://localhost:8080")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", addr, err)
	}

	<-done
}
