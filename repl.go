package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("")
	fmt.Println("Hello! 😊 type your commands below to start")
	fmt.Println("Type e.g. 'help' for the help menu")

	for {
		fmt.Println("")
		fmt.Print(": ")
		input.Scan()
		text := input.Text()
		cleaned := cleanInput(text)
		fmt.Println("")
		if len(cleaned) == 0 {
			fmt.Println(": ")
			continue
		}
		commandName := cleaned[0]

		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "List location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous location areas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location name}",
			description: "List pokemons found at the specified location",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon name}",
			description: "Catch a specific pokemon",
			callback:    callbackCatch,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
