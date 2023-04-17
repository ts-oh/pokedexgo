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

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
			description: "exit the program",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "list location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "list previous location areas",
			callback:    callbackMapb,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
