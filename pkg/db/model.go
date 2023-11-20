package db

import "github.com/polyrain/go-keno/pkg/keno"

// This interface allows accessing the PG database
// And pulling games and cards into useable structs in our program
type KenoData interface {
	GetGame(gameId int) keno.Game
	GetCard(cardId int) keno.Card
	GetGamesBetween(startId int, endId int) []keno.Game
}
