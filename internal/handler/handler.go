package handler

import (
	"fmt"
	"github.com/CatCanCreate/gomeetup/internal/api"
	"net/http"
)

type Handler struct {
	jokeClient api.Client
}

func NewHandler(apiClient api.Client) *Handler {
	return &Handler{
		jokeClient: apiClient,
	}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	joke, err := h.jokeClient.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, joke.Joke)
}
