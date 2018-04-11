package player

import (
	"github.com/kscarlett/muzsh/game"
)

// Inventory represents a player's inventory. It's currently just a collection of Items
type Inventory struct {
	Items []game.Item
}

// NewrInventory returns a pointer to a new empty Inventory
func newInventory() *Inventory {
	return &Inventory{Items: []game.Item{}}
}
