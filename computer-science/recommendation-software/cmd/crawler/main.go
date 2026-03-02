package main

import (
	"fmt"
	"log"
	"recommendation-software/pkg/pokeapi"
	"sync"
)

var (
	cache = make(map[string]bool)
	api   *pokeapi.Pokeapi
	wg    sync.WaitGroup
	mu    sync.Mutex
)

func checkAndSet(url string) bool {
	mu.Lock()
	defer mu.Unlock()
	if cache[url] {
		return true
	}
	cache[url] = true
	return false
}

func crawl(url string) {
	defer wg.Done()
	if checkAndSet(url) {
		return
	}
	res, err := api.GetPokemonURL(url)
	if err != nil {
		log.Printf("Error crawling %s: %v", url, err)
		return
	}
	fmt.Printf("found: %s \n", url)
	nextrun := res.Next
	if nextrun == "" {
		log.Printf("No next run found for %s", url)
		return
	}

	wg.Add(1)
	go crawl(nextrun)

	return
}

func main() {
	log.Println("Starting crawler...")
	api = pokeapi.New()
	wg.Add(1)
	go crawl("https://pokeapi.co/api/v2/pokemon/")
	wg.Wait()
	log.Println("Crawling finished.")
}
