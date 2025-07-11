package main

import (
	"fmt"
	"os"
	"pokedexcli/internal/Requests"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	next     string
	previous string
}

var cfg = config{
	next:     "https://pokeapi.co/api/v2/location-area/",
	previous: "",
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Display available commands",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Get a list of locations",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Get the previous list of locations",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "Get a list of pokemon in a location",
		callback:    commandExplore,
	},
}

func commandExit(c *config, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}

func commandHelp(c *config, arg string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}

func commandMap(c *config, arg string) error {
	err := Requests.MakeRequest(c.next)
	if err != nil {
		return err
	}

	if Requests.JsonMapData.Next == nil {
		c.next = ""
	} else {
		c.next = *Requests.JsonMapData.Next
	}
	if Requests.JsonMapData.Previous == nil {
		c.previous = ""
	} else {
		c.previous = *Requests.JsonMapData.Previous
	}

	printMap()
	return nil
}

func commandMapb(c *config, arg string) error {
	if c.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	err := Requests.MakeRequest(c.previous)
	if err != nil {
		return err
	}

	if Requests.JsonMapData.Next == nil {
		c.next = ""
	} else {
		c.next = *Requests.JsonMapData.Next
	}
	if Requests.JsonMapData.Previous == nil {
		c.previous = ""
	} else {
		c.previous = *Requests.JsonMapData.Previous
	}

	printMap()
	return nil
}

func printMap() {
	for _, item := range Requests.JsonMapData.Results {
		fmt.Println(item.Name)
	}
}

func commandExplore(c *config, arg string) error {
	if err := Requests.ExploreRequest(arg); err != nil {
		return err
	}

	printPokemon()
	return nil
}

func printPokemon() {
	fmt.Println("Found Pokemon:")
	for _, pokemon := range Requests.JsonExploreData.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
}
