package api

// Client interacts with 3-rd party joke API
type Client interface {
	GetJoke() (*JokeResponse, error)
}
