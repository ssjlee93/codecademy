package pokeapi

import (
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

type Pokeapi struct {
	Client *http.Client
}

func New() *Pokeapi {
	return &Pokeapi{
		Client: &http.Client{},
	}
}
