package pokeapi

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strings"
)

var dataLink = Links{
	NextUrl: "https://pokeapi.co/api/v2/location-area",
	PrevUrl: "",
}

var totalCaughtPokemons = make(map[string]Pokemon)

func PokedexCommand(_ string, _ string) {
	fmt.Println("Your Pokedex:")
	for item := range totalCaughtPokemons {
		fmt.Println(item)
	}
}

func InspectCommand(_ string, text string) {
	words := strings.Fields(text)
	if len(words) > 1 {
		pokemonToInspect := words[1]
		pokemon, exists := totalCaughtPokemons[pokemonToInspect]
		if exists {
			fmt.Println(pokemon)
		} else {
			fmt.Println("you have not caught that pokemon")
		}
	}
}

func CatchCommand(_ string, text string) {
	words := strings.Fields(text)
	if len(words) > 1 {
		pokemonToCatch := words[1]
		fmt.Printf("Throwing a Pokeball at %v...\n", pokemonToCatch)
		getDataByPokemonName(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonToCatch))
	}
}

func ExploreCommand(_ string, text string) {
	words := strings.Fields(text)
	if len(words) > 1 {
		fmt.Printf("Exploring %s...\n", words[1])
		urlToSend := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", words[1])
		getDataByLocationName(urlToSend)
	}
}

func CommandMap(_ string, _ string) {
	result := getData(dataLink.NextUrl)
	for _, location := range result.Results {
		fmt.Println(location.Name)
	}
}

func CommandMapb(_ string, _ string) {
	if dataLink.PrevUrl != "" {
		result := getData(dataLink.PrevUrl)
		for _, location := range result.Results {
			fmt.Println(location.Name)
		}
	}
}

func getData(urlToFetch string) Response {
	res, resErr := http.Get(urlToFetch)
	if resErr != nil {
		log.Fatal("Error while fetching data:", resErr)
	}

	defer res.Body.Close()
	var responseToReturn Response

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&responseToReturn); err != nil {
		log.Fatal("Error while decoding:", err)
	}

	dataLink.NextUrl = responseToReturn.Next
	dataLink.PrevUrl = responseToReturn.Prev

	return responseToReturn
}

func getDataByLocationName(url string) {
	res, resErr := http.Get(url)
	if resErr != nil {
		log.Fatal(resErr)
	}

	var result PokemonResult

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&result); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range result.Pokemon_encounters {
		fmt.Println(encounter.Pokemon.Name)
	}
}

func getDataByPokemonName(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	var pokemon Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	catchChance := rand.IntN(100) + 1
	threshold := 100 - pokemon.Base_experience/10

	if catchChance > threshold {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		totalCaughtPokemons[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
}
