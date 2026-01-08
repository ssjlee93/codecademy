package main

import (
	"fmt"
	"log"
	"recommendation-software/pkg/pokeapi"
)

func main() {
	log.Println("Application started")

	client := pokeapi.New()
	result := client.GetTypes(25, 0)
	for i, r := range result {
		fmt.Printf("%d: %s\n", i, r)
	}
}
