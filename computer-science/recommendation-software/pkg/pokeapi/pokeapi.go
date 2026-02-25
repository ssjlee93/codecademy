package pokeapi

import (
	"encoding/json"
	"fmt"
	"log"
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

type NamedApiResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type TypeList struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedApiResource `json:"results"`
}

func (p *Pokeapi) GetTypes(limit, offset int) []string {
	url := fmt.Sprintf("%s/type?limit=%d&offset=%d", baseURL, limit, offset)
	resp, err := p.Client.Get(url)
	if err != nil {
		log.Printf("Error getting type list : %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error getting type list: status code %d\n", resp.StatusCode)
		return nil
	}

	var typeList TypeList
	if err := json.NewDecoder(resp.Body).Decode(&typeList); err != nil {
		log.Printf("Error decoding type list : %v\n", err)
		return nil
	}

	result := make([]string, 0, len(typeList.Results))
	for _, r := range typeList.Results {
		if r.Name != "" {
			result = append(result, r.Name)
		}
	}

	return result
}

type Pokemon struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	GameIndices []struct {
		GameIndex int              `json:"game_index"`
		Version   NamedApiResource `json:"version"`
	} `json:"game_indices"`
}

func (p *Pokeapi) GetPokemon(url string) (*Pokemon, error) {
	resp, err := p.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting pokemon: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting pokemon: status code %d", resp.StatusCode)
	}

	var pokemon Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return nil, fmt.Errorf("error decoding pokemon: %w", err)
	}

	return &pokemon, nil
}

func (p *Pokeapi) GetPokemonList(limit, offset int) (*TypeList, error) {
	url := fmt.Sprintf("%s/pokemon?limit=%d&offset=%d", baseURL, limit, offset)
	resp, err := p.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error getting pokemon list: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting pokemon list: status code %d", resp.StatusCode)
	}

	var typeList TypeList
	if err := json.NewDecoder(resp.Body).Decode(&typeList); err != nil {
		return nil, fmt.Errorf("error decoding pokemon list: %w", err)
	}

	return &typeList, nil
}
