package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("🚑 Welcome to the pokedex help menu")
	fmt.Println("⌨️ Available commands")
	fmt.Println("")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}

	return nil
}
