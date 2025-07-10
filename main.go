package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	stringSlice := strings.Fields(text)

	return stringSlice
}

func main() {
	buffer := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if buffer.Scan() {
			inputText := cleanInput(buffer.Text())
			if _, ok := commands[inputText[0]]; !ok {
				fmt.Println("Unknown command. Type 'help' for a list of commands.")
				continue
			}
			err := commands[inputText[0]].callback(&cfg)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		}
	}
}
