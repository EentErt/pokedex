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
			fmt.Printf("Your command was: %s\n", inputText[0])
		}
	}
}
