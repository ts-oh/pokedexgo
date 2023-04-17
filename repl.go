package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startInput() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("")
	fmt.Println("Hello! 😊 type your commands below to start")
	fmt.Println("Type e.g. 'help' for the help menu")
	fmt.Println("")

	for {
		fmt.Print(":")
		input.Scan()
		text := input.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			fmt.Println(":")
			continue
		}
		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		command.callback()
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
