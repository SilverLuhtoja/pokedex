package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func initApp(cfg *Config) {
	fmt.Println(cfg.Client.BaseURL)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\033[35mPokedex >\033[0m ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		givenCommand := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		if command, ok := getCommands()[givenCommand]; ok {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("No command available for > %s <. Type help for available commands.\n\n", input)
		}
	}
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}
