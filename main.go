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

	fmt.Printf(`Welcome to Dark Wood, %s 🌳🌲
You must find the Onepiece to unlock its power.
You can move around the forest using the commands: north, south, east, west.
You can also look around using the command: look.
You can also quit the game using the command: quit.
----Starting Game----- 
`, game.Player.Name)
	fmt.Printf(game.CurrentLocation.Description)

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
		fmt.Printf("You are at %s. %s\n", game.CurrentLocation.Name, game.CurrentLocation.Description)
	case "north", "south", "east", "west":
		if newLoc, exists := game.CurrentLocation.Exits[command]; exists {
			game.CurrentLocation = newLoc
			fmt.Printf("You move %s.\n", command)
			fmt.Println(game.CurrentLocation.Description)
		} else {
			fmt.Println("You can't go that way.")
		}
	default:
		fmt.Println("I don't understand your command, please try again.")
	}
}
