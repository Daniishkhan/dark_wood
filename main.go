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
You can pick up items using the command: pickup <item name>.
You can drop items using the command: drop <item name>.
You can examine your inventory using the command: inventory.
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
	words := strings.Fields(command)
	switch words[0] {
	case "look":
		fmt.Printf("You are at %s. %s\n", game.CurrentLocation.Name, game.CurrentLocation.Description)
		if len(game.CurrentLocation.Items) > 0 {
			fmt.Println("You see the following items:")
			for _, item := range game.CurrentLocation.Items {
				fmt.Printf("- %s: %s\n", item.Name, item.Description)
			}
		}
	case "north", "south", "east", "west":
		if newLoc, exists := game.CurrentLocation.Exits[command]; exists {
			game.CurrentLocation = newLoc
			fmt.Printf("You move %s.\n", command)
			fmt.Println(game.CurrentLocation.Description)
		} else {
			fmt.Println("You can't go that way.")
		}
	case "inventory":
		if len(game.Player.Inventory) == 0 {
			fmt.Println("You dont have any inventory")
		} else {
			fmt.Println("You have the following items")
			for _, item := range game.Player.Inventory {
				fmt.Printf("- %s: %s\n", item.Name, item.Description)
			}
		}
	case "take":
		if len(words) < 2 {
			fmt.Println("Take what?")
		} else {
			itemName := strings.Join(words[1:], " ")
			game.PickUpItem(itemName)
		}
	case "drop":
		if len(words) < 2 {
			fmt.Println("Drop what?")
		} else {
			itemName := strings.Join(words[1:], " ")
			fmt.Printf("Attempting to drop: %s\n", itemName)
			fmt.Printf("Current inventory: %v\n", game.Player.Inventory)
			game.DropItem(itemName)
		}
	case "examine":
		if len(words) < 2 {
			fmt.Println("Examine what?")
		} else {
			itemName := strings.Join(words[1:], " ")
			game.ExamineItem(itemName)
		}
	default:
		fmt.Println("I don't understand your command, please try again.")
	}
}
