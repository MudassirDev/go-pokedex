package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MudassirDev/go-pokedex/packages/pokeapi"
)

type command struct {
	name     string
	message  string
	function func(string, string)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	allCommands := map[string]command{
		"exit": {
			name:     "exit",
			message:  "Closing the Pokedex... Goodbye!",
			function: commandExit,
		},
		"help": {
			name:     "help",
			message:  "Welcome to the Pokedex!\nUsage:\n\n\nhelp: Displays a help message\nexit: Exit the Pokedex",
			function: commandHelp,
		},
		"map": {
			name:     "map",
			message:  "",
			function: pokeapi.CommandMap,
		},
		"mapb": {
			name:     "mapb",
			message:  "",
			function: pokeapi.CommandMapb,
		},
		"explore": {
			name:     "explore",
			message:  "",
			function: pokeapi.ExploreCommand,
		},
		"catch": {
			name:     "catch",
			message:  "",
			function: pokeapi.CatchCommand,
		},
		"inspect": {
			name:     "inspect",
			message:  "",
			function: pokeapi.InspectCommand,
		},
		"pokedex": {
			name:     "pokedex",
			message:  "",
			function: pokeapi.PokedexCommand,
		},
	}

	fmt.Print("Pokedex > ")

	for scanner.Scan() {
		text := scanner.Text()

		for item := range allCommands {
			currentCommand := allCommands[item]
			if strings.Contains(text, item) {
				currentCommand.function(currentCommand.message, text)
			}
		}

		fmt.Print("Pokedex > ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func cleanInput(text string) []string {
	sliceToRetun := strings.Fields(strings.ToLower(text))
	return sliceToRetun
}

func commandExit(message string, _ string) {
	fmt.Println(message)
	os.Exit(0)
}

func commandHelp(message string, _ string) {
	fmt.Println(message)
}
