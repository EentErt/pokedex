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

var pokedex = make(map[string]Requests.JsonPokemon)

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
	"catch": {
		name:        "catch",
		description: "Catch a pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "Inspect a pokemon in your pokedex",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "View your caught pokemon",
		callback:    commandPokedex,
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

func commandCatch(c *config, pokemon string) error {
	if err := Requests.CatchRequest(pokemon); err != nil {
		return err
	}

	fmt.Printf("%s was caught!\n", pokemon)
	addToPokedex(pokemon)
	return nil
}

func addToPokedex(pokemon string) {
	pokedex[pokemon] = Requests.JsonPokemonData
}

func commandInspect(c *config, pokemon string) error {
	if data, ok := pokedex[pokemon]; ok {
		fmt.Printf("Name: %s\n", data.Name)
		fmt.Printf("Height: %d\n", data.Height)
		fmt.Printf("Weight: %d\n", data.Weight)
		fmt.Println("Stats:")
		for _, stat := range data.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range data.Types {
			fmt.Printf("  -%s\n", t.Type.Name)
		}
		return nil
	}
	return fmt.Errorf("you have not caught that pokemon")
}

func commandPokedex(c *config, arg string) error {
	for _, pokemon := range pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
