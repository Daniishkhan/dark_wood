# Key Go Concepts for CLI Game Development

## 1. Basic Syntax and Control Structures

- Variables and constants declaration
- Function declaration and usage
- Conditional statements (if, else, switch)
- Loops (for, range)

Example:
```go
func processCommand(command string) {
    switch command {
    case "look":
        fmt.Println("You see a dark forest.")
    default:
        fmt.Println("Unknown command.")
    }
}
```

## 2. Packages and Imports

- Understanding the `main` package
- Importing standard library packages
- Creating and using custom packages

Example:
```go
package main

import (
    "fmt"
    "strings"
    "github.com/yourusername/darkwoods/game"
)
```

## 3. Structs and Methods

- Defining custom types with structs
- Adding methods to structs

Example:
```go
type Player struct {
    Name     string
    Location string
    Items    []string
}

func (p *Player) Move(direction string) {
    // Implementation
}
```

## 4. Interfaces

- Defining and implementing interfaces
- Using interfaces for abstraction

Example:
```go
type GameObject interface {
    Describe() string
    Interact(player *Player) string
}
```

## 5. Error Handling

- Creating and returning errors
- Using the `error` interface
- Handling errors with `if err != nil`

Example:
```go
func (p *Player) PickUp(item string) error {
    if !itemExists(item) {
        return fmt.Errorf("item %s does not exist", item)
    }
    // Implementation
    return nil
}
```

## 6. Slices and Maps

- Using slices for dynamic arrays
- Using maps for key-value storage

Example:
```go
type Game struct {
    Locations map[string]*Location
    Items     []string
}
```

## 7. Goroutines and Channels (for advanced features)

- Basic concurrency with goroutines
- Communication between goroutines using channels

Example:
```go
func (g *Game) StartBackgroundTasks() {
    go g.updateGameState()
}
```

## 8. File I/O

- Reading from and writing to files (for save/load functionality)

Example:
```go
import "io/ioutil"

func SaveGame(g *Game) error {
    data, err := json.Marshal(g)
    if err != nil {
        return err
    }
    return ioutil.WriteFile("savegame.json", data, 0644)
}
```

## 9. JSON Marshaling/Unmarshaling

- Converting structs to JSON and vice versa (for save/load functionality)

Example:
```go
import "encoding/json"

func LoadGame(filename string) (*Game, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    var game Game
    err = json.Unmarshal(data, &game)
    if err != nil {
        return nil, err
    }
    return &game, nil
}
```

## 10. Testing

- Writing and running unit tests
- Using the `testing` package

Example:
```go
func TestPlayerMove(t *testing.T) {
    p := &Player{Location: "Forest"}
    p.Move("north")
    if p.Location != "Cave" {
        t.Errorf("Expected location to be Cave, got %s", p.Location)
    }
}
```

These concepts form the foundation of Go programming and will be crucial in developing your CLI game. As you build your game, you'll naturally engage with these concepts, deepening your understanding of Go.