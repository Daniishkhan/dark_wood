# Dark Wood - Go CLI Adventure Game

## Overview
Dark Wood is a text-based adventure game played through the command line interface (CLI). The game immerses players in a mysterious forest filled with dangers, secrets, and treasures. The ultimate objective is to find the powerful weapon called "Onepiece," hidden deep within the dark Wood.

## Story
In a realm shrouded in perpetual twilight, legends speak of a powerful weapon known as the Onepiece, a blade forged by the ancient gods and hidden away in the depths of the Dark Wood. Many have ventured into the Wood in search of this legendary weapon, but none have returned.

## Objective
As a brave adventurer, you have decided to embark on this perilous quest. Armed with only your wits and a simple map, you must navigate through the twisted paths of the forest, confront fearsome creatures, solve ancient puzzles, and uncover hidden secrets. Will you be the one to find the Onepiece and unlock its incredible power?

## Game Features
- Text-Based Exploration: Navigate through the forest using simple text commands.
- Interactive Storytelling: Engage with the environment and characters through descriptive narrative.
- Puzzles and Challenges: Solve various puzzles and overcome challenges to progress.
- Inventory Management: Collect items that aid in your quest and manage your inventory.
- Random Encounters: Face random encounters with creatures and events that will test your survival skills.

## Go-Specific Concepts and Implementation Details

### 1. Project Structure
Organize your project using Go's recommended package structure:

```
dark-Wood/
├── cmd/
│   └── darkWood/
│       └── main.go
├── internal/
│   ├── game/
│   │   ├── game.go
│   │   ├── player.go
│   │   ├── item.go
│   │   ├── enemy.go
│   │   └── map.go
│   ├── ui/
│   │   └── cli.go
│   └── utils/
│       └── random.go
├── pkg/
│   └── storage/
│       └── save.go
├── go.mod
└── go.sum
```

This structure separates concerns and promotes modularity.

### 2. Interfaces and Dependency Injection
Use interfaces to define behavior and allow for easy testing and flexibility:

```go
type GameState interface {
    Update()
    Render()
}

type InputHandler interface {
    HandleInput(input string) error
}

type Game struct {
    state GameState
    input InputHandler
}
```

### 3. Concurrency with Goroutines and Channels
Implement background tasks using goroutines and communicate using channels:

```go
func (g *Game) Run() {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            g.UpdateWorld()
        case input := <-g.inputChan:
            g.HandleInput(input)
        }
    }
}
```

### 4. Error Handling
Use custom error types and the `errors` package for robust error handling:

```go
import "errors"

var ErrInvalidMove = errors.New("invalid move")

func (p *Player) Move(direction string) error {
    // ... implementation
    return ErrInvalidMove
}
```

### 5. Testing
Write unit tests for your game logic using the `testing` package:

```go
func TestPlayerMove(t *testing.T) {
    player := NewPlayer()
    err := player.Move("north")
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
    // Add more assertions
}
```

### 6. Serialization
Use `encoding/json` for saving and loading game state:

```go
import "encoding/json"

func (g *Game) Save() error {
    data, err := json.Marshal(g)
    if err != nil {
        return err
    }
    // Save data to file
}
```

### 7. Command-Line Flags
Utilize the `flag` package for command-line options:

```go
import "flag"

func main() {
    newGame := flag.Bool("new", false, "Start a new game")
    loadGame := flag.String("load", "", "Load a saved game file")
    flag.Parse()

    // Use flags to determine game start behavior
}
```

### 8. Logging
Implement logging for debugging and monitoring:

```go
import "log"

func init() {
    log.SetPrefix("DARK Wood: ")
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// Usage
log.Printf("Player entered room %s", roomID)
```

## Build and Run
To build and run the game:

1. Ensure you have Go installed (version 1.16+)
2. Clone the repository
3. Navigate to the project directory
4. Run `go build ./cmd/darkWood`
5. Execute the built binary: `./darkWood`

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.