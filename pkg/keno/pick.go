package keno

import (
	"log/slog"
)

// A Card is an individuals card
// It has their "picks" (choice of nums)
// As well as the price per game, # of picks per game
// The starting game ID, and the number of games they've paid for
type Card struct {
	CardId    int   `json:"card_id"`
	NumPicks  int   `json:"picks_per_game"`
	Picks     []int `json:"picks"`
	GamePrice int   `json:"price_per_game"`
	NumGames  int   `json:"number_games"`
	StartId   int   `json:"start_game_num"`
}

// Test all values in provided picks to ensure they don't put invalid nums
func ValidatePicks(picks []int) bool {
	// TODO: Validate length of picks here too
	for _, pick := range picks {
		if pick > 80 || pick < 0 {
			return false
		}
	}

	return true
}

// Take a game and figure out how much we won on it
func (c *Card) CalcWin(g *Game) int {
	if g.Id < c.StartId || g.Id > (c.StartId+c.NumGames) {
		slog.Error("Game not valid for this card.", "game", g.Id)
		return 0
	}

	var numCorrect int = 0
	// Since the picks map to indices in the Games board, we can just check each
	// pick against the game and test if its 1, O(n) moment
	// since picks are 1 indexed we minus one
	for _, pick := range c.Picks {
		if g.Board[pick-1] == 1 {
			numCorrect++
		}
	}

	return numCorrect
}
