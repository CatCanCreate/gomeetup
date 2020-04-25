package jokes

import (
	"encoding/json"
	"fmt"
	"github.com/CatCanCreate/gomeetup/internal/api"
	"net/http"
)

const getJokePath = "/api?format=json"

type JokeClientAPI struct {
	url string
}

func NewJokeClientAPI(baseURL string) *JokeClientAPI {
	return &JokeClientAPI{url: baseURL}
}

func (jc *JokeClientAPI) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + getJokePath

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request error: %s", http.StatusText(resp.StatusCode))
	}

	var data api.JokeResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil

}
