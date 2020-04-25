package handler

import (
	"github.com/CatCanCreate/gomeetup/internal/api"
	"github.com/CatCanCreate/gomeetup/internal/api/apiMocks"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Hello(t *testing.T) {
	tests := []struct {
		name     string
		joke     *api.JokeResponse
		err      error
		codeWant int
		bodyWant string
	}{
		{
			name:     "simple test",
			joke:     &api.JokeResponse{Joke: "test joke"},
			err:      nil,
			codeWant: 200,
			bodyWant: "test joke",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiMock := &apiMocks.Client{}
			apiMock.On("GetJoke").Return(tt.joke, tt.err)

			h := NewHandler(apiMock)

			req, _ := http.NewRequest("GET", "/hello", nil)
			rr := httptest.NewRecorder()

			h.Hello(rr, req)

			gotRaw, _ := ioutil.ReadAll(rr.Body)
			got := string(gotRaw)

			if got != tt.bodyWant {
				t.Errorf("wrong response body %s want %s", got, tt.bodyWant)
			}

			if status := rr.Result().StatusCode; status != tt.codeWant {
				t.Errorf("wrong response code %d want %d", status, tt.codeWant)
			}

		})
	}
}
