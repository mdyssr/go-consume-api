package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// A Response struct to map the Entire Response
type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// // A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

const url = "http://pokeapi.co/api/v2/pokedex/kanto/"

func main() {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("error requesting %s\n", url)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	var responseObject Response
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatal("error parsing json")
	}

	for _, pokemon := range responseObject.Pokemon {
		fmt.Println(pokemon.Species.Name)
	}

}
