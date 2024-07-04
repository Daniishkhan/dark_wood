package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(`Welcome to Dark Wood 🌳🌲
You are a brave adventurer in the Dark Woods. You must find the Onepiece to unlock its power.
You can move around the forest using the commands: north, south, east, west.
You can also look around using the command: look.
You can also quit the game using the command: quit. Good luck!
----Starting Game-----`)
	fmt.Println("What is your name adventurer?")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Nice to meet you,", name)

	for {
		fmt.Println("Enter Command:")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSuffix(command, "\n")
		fmt.Println("You entered:", command)

		if command == "quit" {
			fmt.Println("Thank you for playing Dark Wood!")
			break
		}
		processCommand(command, name)
	}

}

func processCommand(command string, name string) {
	fmt.Println("Processing command:", command)
	switch command {
	case "look":
		fmt.Printf("You are in the dark woods, %s.\n", name)
	case "north":
		fmt.Println("You moved north.")
	case "south":
		fmt.Println("You moved south.")
	case "east":
		fmt.Println("You moved east.")
	case "west":
		fmt.Println("You moved west.")
	default:
		fmt.Println("I don't know what to do")
	}
}
