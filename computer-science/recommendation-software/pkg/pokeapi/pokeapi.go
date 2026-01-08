package pokeapi

import (
	pokego "github.com/JoshGuarino/PokeGo/pkg"
	"log"
)

type Pokeapi struct {
	Client pokego.PokeGo
}

func New() *Pokeapi {
	client := pokego.NewClient()
	return &Pokeapi{Client: client}
}

func (p *Pokeapi) GetTypes(limit, offset int) []string {
	// Main client example returning first page of 20 results
	typeList, err := p.Client.Pokemon.GetTypeList(limit, offset)
	if err != nil {
		log.Printf("Error getting type list : %v\n", err)
	}

	// unwrap the data
	result := make([]string, 0, len(typeList.Results))

	for _, r := range typeList.Results {
		if r.Name != "" {
			result = append(result, r.Name)
		}
	}

	return result
}
