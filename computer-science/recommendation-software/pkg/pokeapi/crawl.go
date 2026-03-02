package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokemonURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokemonURLResponse struct {
	Count    int          `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []PokemonURL `json:"results"`
}

func (p *Pokeapi) GetPokemonURL(url string) (*PokemonURLResponse, error) {
	resp, err := p.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting sources: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting source: status code %d", resp.StatusCode)
	}

	var pokemonURLResponse PokemonURLResponse
	if err := json.NewDecoder(resp.Body).Decode(&pokemonURLResponse); err != nil {
		return nil, fmt.Errorf("error decoding pokemon: %w", err)
	}

	return &pokemonURLResponse, nil
}
