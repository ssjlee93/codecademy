package main

import (
	"fmt"
	"log"
	"recommendation-software/pkg/pokeapi"
	"sync"
)

type Cache struct {
	mu      sync.Mutex
	visited map[string]bool
}

func (c *Cache) checkAndSet(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.visited[url] {
		return true
	}
	c.visited[url] = true
	return false
}

var (
	cache = Cache{visited: make(map[string]bool)}
	api   *pokeapi.Pokeapi
	wg    sync.WaitGroup
)

func crawl(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	pokemon, err := api.GetPokemon(url)
	if err != nil {
		log.Printf("Error crawling %s: %v", url, err)
		return
	}

	games := make([]string, 0, len(pokemon.GameIndices))
	for _, gi := range pokemon.GameIndices {
		games = append(games, gi.Version.Name)
	}

	fmt.Printf("ID: %d | Name: %s | Games: %v\n", pokemon.ID, pokemon.Name, games)
}

func main() {
	log.Println("Starting crawler...")
	api = pokeapi.New()

	list, err := api.GetPokemonList(20, 0)
	if err != nil {
		log.Fatalf("Failed to get pokemon list: %v", err)
	}

	for _, res := range list.Results {
		if cache.checkAndSet(res.URL) {
			continue
		}
		wg.Add(1)
		go crawl(res.URL, &wg)
	}

	wg.Wait()
	log.Println("Crawling finished.")
}
