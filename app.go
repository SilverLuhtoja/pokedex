package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func initApp() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		givenCommand := input[0]
		if command, ok := getCommands()[givenCommand]; ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("No command available for > %s <. Type help for available commands.\n\n", input)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback: func() error {
				fmt.Println()
				fmt.Println("help: Displays a help message")
				fmt.Println("exit: Exit the Pokedex")
				return nil
			},
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback: func() error {
				os.Exit(1)
				return nil
			},
		},
	}
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}
