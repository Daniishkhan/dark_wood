package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What is your name adventurer?")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	game := NewGame()
	game.Player.Name = name // Set the player's name

	fmt.Printf("Nice to meet you, %s\n", game.Player.Name)
	fmt.Println(game.StartingLocation.Description)

	for {
		fmt.Println("\nEnter Command:")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "quit" {
			fmt.Println("Thank you for playing Dark Wood!")
			break
		}

		processCommand(game, command)
	}
}

func processCommand(game *Game, command string) {
	fmt.Println("Processing command:", command)
	switch command {
	case "look":
		fmt.Printf("You are at %s. %s\n", game.Player.PlayerLocation.Name, game.Player.PlayerLocation.Description)
	case "north", "south", "east", "west":
		fmt.Printf("You tried to move %s, but movement is not implemented yet.\n", command)
	default:
		fmt.Println("I don't understand your command, please try again.")
	}
}
