package keno

import (
	"fmt"
	"math/rand"
)

// Keno logic, used by main app/server
// Keno rules are simple: Randomly, 40 numbers
// Out of 80 get picked
// Players have pre-chosen the numbers for the game
// They get prizes depending on their choice of #,
// including how many they chose

// Also, we have heads/tails and evens/odds
// As optional extra game modes
// ie we need to have a generic way of checking diff
// rules against the same boardstate

type Game struct {
	Id int `json:"id"`
	// While the Keno board is 2D when viewing,
	// We store it as a flat array for easier use
	Board    []int `json:"board"`
	Finished bool  `json:"completed"`
}

const ROWS int = 8
const COLS int = 10
const TOTAL_NUMS int = 40

func NewGame(id int) *Game {
	// Make a new board
	// We'll check which indices have a 1 in them when checking for wins

	game := &Game{
		Id:       id,
		Board:    make([]int, 0, ROWS*COLS),
		Finished: false,
	}

	i := 0

	for i < ROWS*COLS {
		game.Board = append(game.Board, 0)
		i += 1
	}

	game.PickNumbers()

	return game
}

// Generate a random permutation of the Keno board, taking the last 20 numbers
// This gives us a random set of valid numbers without repetetion
func (g *Game) PickNumbers() []int {
	indices := rand.Perm(ROWS * COLS)[:TOTAL_NUMS]
	for _, i := range indices {
		g.Board[i] = 1
	}

	g.Finished = true

	return g.Board
}

func (g *Game) RenderGame() {
	fmt.Print("|")
	for i := range g.Board {
		// 80 nums total, every 10
		if g.Board[i] == 1 {
			fmt.Printf("%3d", i+1)
		} else {
			fmt.Printf("%3s", "x")
		}
		if i > 0 && (i+1)%10 == 0 {
			fmt.Println("|")
			if (i + 1) < COLS*ROWS {
				fmt.Print("|")
			}
		}
	}

}
