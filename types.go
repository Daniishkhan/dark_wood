package main

import (
	"fmt"
	"strings"
)

type Game struct {
	CurrentLocation *Location
	Player          Player
	Locations       map[string]*Location
}

type Player struct {
	Name      string
	Inventory []Item
}

type Location struct {
	Name        string
	Description string
	Exits       map[string]*Location
	Items       []Item
}

type Item struct {
	Name        string
	Description string
}

func NewGame() *Game {
	// Create locations
	start := &Location{
		Name:        "Start",
		Description: `You are at the gate of the Darkwoods.`,
		Exits:       make(map[string]*Location),
	}

	deepForest := &Location{
		Name:        "Forest",
		Description: "You are in a dense, dark forest. The trees loom ominously above you.",
		Exits:       make(map[string]*Location),
		Items: []Item{
			{Name: "Rock", Description: "A large rock that can be used to break open a door."},
		},
	}

	clearingWithWell := &Location{
		Name:        "Clearing with Well",
		Description: "You emerge into a small clearing. In the center stands an old stone well, its stones covered in moss.",
		Exits:       make(map[string]*Location),
		Items: []Item{
			{Name: "Water", Description: "A clear stream of water flows from the well."},
		},
	}

	abandonedCabin := &Location{
		Name:        "Abandoned Cabin",
		Description: "You emerge into a small clearing. In the center stands an old stone cabin, its walls covered in old mud.",
		Exits:       make(map[string]*Location),
		Items: []Item{
			{Name: "Rock", Description: "A large rock that can be used to break open a door."},
		},
	}

	// Connect locations
	start.Exits["north"] = deepForest
	deepForest.Exits["south"] = start
	deepForest.Exits["east"] = clearingWithWell
	deepForest.Exits["west"] = abandonedCabin
	clearingWithWell.Exits["west"] = deepForest
	abandonedCabin.Exits["east"] = deepForest

	// Create game
	game := &Game{
		CurrentLocation: start,
		Player: Player{
			Name:      "", // We'll set this in main.go
			Inventory: []Item{},
		},
		Locations: map[string]*Location{
			"Start":               start,
			"Deep Forest":         deepForest,
			"Clearning With Well": clearingWithWell,
			"Abandoned Cabin":     abandonedCabin,
		},
	}

	return game
}

func (l *Location) AddItem(item Item) {
	l.Items = append(l.Items, item)
}

func (l *Location) RemoveItem(item Item) {
	for i, v := range l.Items {
		if v.Name == item.Name {
			l.Items = append(l.Items[:i], l.Items[i+1:]...)
			return
		}
	}
}

// pickup item
func (g *Game) PickUpItem(itemName string) {
	for i, item := range g.CurrentLocation.Items {
		if item.Name == itemName {
			g.Player.Inventory = append(g.Player.Inventory, item)
			g.CurrentLocation.Items = append(g.CurrentLocation.Items[:i], g.CurrentLocation.Items[i+1:]...)
			fmt.Printf("You picked up the %s.\n", itemName)
			return
		}
	}
	fmt.Printf("There is no %s here.\n", itemName)
}

func (g *Game) DropItem(itemName string) {
	for i, item := range g.Player.Inventory {
		if strings.EqualFold(item.Name, itemName) {
			g.CurrentLocation.Items = append(g.CurrentLocation.Items, item)
			g.Player.Inventory = append(g.Player.Inventory[:i], g.Player.Inventory[i+1:]...)
			fmt.Printf("You dropped the %s.\n", item.Name)
			return
		}
	}
	fmt.Printf("You don't have a %s.\n", itemName)
}

func (g *Game) ExamineItem(itemName string) {
	// Check inventory first
	for _, item := range g.Player.Inventory {
		if strings.EqualFold(item.Name, itemName) {
			fmt.Println(item.Description)
			return
		}
	}
	// Then check current location
	for _, item := range g.CurrentLocation.Items {
		if strings.EqualFold(item.Name, itemName) {
			fmt.Println(item.Description)
			return
		}
	}
	fmt.Printf("You don't see a %s here.\n", itemName)
}
