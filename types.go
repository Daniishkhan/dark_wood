package main

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

	forest := &Location{
		Name:        "Forest",
		Description: "You are in a dense, dark forest. The trees loom ominously above you.",
		Exits:       make(map[string]*Location),
	}

	// Connect locations
	start.Exits["north"] = forest
	forest.Exits["south"] = start

	// Create game
	game := &Game{
		CurrentLocation: start,
		Player: Player{
			Name:      "", // We'll set this in main.go
			Inventory: []Item{},
		},
		Locations: map[string]*Location{
			"Start":  start,
			"Forest": forest,
		},
	}

	return game
}
