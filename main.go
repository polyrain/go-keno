package main

import (
	"fmt"

	"github.com/polyrain/go-keno/pkg/keno"
)

// We generate a new game every X seconds and write it to the DB immediately
// Players ask us for cards, and we provide them (and write them to the DB)
// Players can ask us to check a card against the games to figure out their win
func main() {
	fmt.Println("hello world")
	game := keno.NewGame(1)
	//fmt.Println(game)
	//ame.PickNumbers()
	//fmt.Println(game)
	game.RenderGame()
}
