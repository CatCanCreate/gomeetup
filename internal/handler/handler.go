package handler

import (
	"fmt"
	"github.com/CatCanCreate/gomeetup/internal/api"
	"net/http"
)

type Handler struct {
	JokeClient api.Client
}

func NewHandler(apiClient api.Client) *Handler {
	return &Handler{
		JokeClient: apiClient,
	}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	joke, err := h.JokeClient.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, joke.Joke)
}
