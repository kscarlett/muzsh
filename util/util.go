package util

import (
	"log"

	"github.com/kscarlett/muzsh/game"
)

// RemoveItem removes a given game.Item from the given inventory and returns the resulting inventory
func RemoveItem(collection []game.Item, item *game.Item) []game.Item {
	for i := range collection {
		if collection[i] == *item {
			copy(collection[i:], collection[i+1:])
			collection = collection[:len(collection)-1]
			return collection
		}
	}

	log.Panic("Something went wrong - could not remove item from collection")
	return collection
}