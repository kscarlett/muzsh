package player

import (
	"github.com/kscarlett/muzsh/game"
)

// Player represents the player of the game. It holds the player's name, their inventory and their health, hunger and thirst values.
type Player struct {
	Name        string
	Health      int
	Fear        int
	Money       int
	Hunger      int
	Thirst      int
	Inventory   *Inventory
	CurrentRoom *game.Room
}

// New returns a new Player with an empty inventory and full health, hunger and thirst.
func New() *Player {
	return &Player{
		Health:    100,
		Fear:      10,
		Money:     0,
		Hunger:    100,
		Thirst:    100,
		Inventory: newInventory()}
}
