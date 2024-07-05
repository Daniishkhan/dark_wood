package main

type Game struct {
	StartingLocation Location
	Player           Player
}

type GameMap struct {
	Name        string
	Description string
	Locations   []Location
}

type Player struct {
	Name           string
	PlayerLocation *Location
	Inventory      []Item
}

type Location struct {
	Name        string
	Description string
}

type Room struct {
	Name        string
	Description string
	Exits       map[string]Location
}

type Item struct {
	Name        string
	Description string
}

func NewGame() *Game {
	startingLocation := &Location{
		Name: "Start",
		Description: `Welcome to Dark Wood 🌳🌲
		You are a brave adventurer in the Dark Woods. You must find the Onepiece to unlock its power.
		You can move around the forest using the commands: north, south, east, west.
		You can also look around using the command: look.
		You can also quit the game using the command: quit. Good luck!
		----Starting Game-----`,
	}

	player := Player{
		Name:           "Danish",
		PlayerLocation: startingLocation,
	}

	return &Game{
		StartingLocation: *startingLocation,
		Player:           player,
	}

}
