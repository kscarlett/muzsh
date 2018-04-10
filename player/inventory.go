package player

import (
	"github.com/kscarlett/muzsh/game"
)

// PlayerInventory represents a player's inventory. It's currently just a collection of Items
type PlayerInventory struct {
	Items []game.Item
}

// NewPlayerInventory returns a pointer to a new empty PlayerInventory
func newPlayerInventory() *PlayerInventory {
	return &PlayerInventory{Items: []game.Item{}}
}
