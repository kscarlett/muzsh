package player

import (
	"github.com/kscarlett/muzsh/item"
)

// Inventory represents a player's inventory. It's currently just a collection of Items
type Inventory struct {
	Items []item.Item
}

// NewrInventory returns a pointer to a new empty Inventory
func newInventory() *Inventory {
	return &Inventory{Items: []item.Item{}}
}
