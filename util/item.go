package util

import (
	"log"

	"github.com/kscarlett/muzsh/item"
)

// RemoveItem removes a given item.Item from the given inventory and returns the resulting inventory
func RemoveItem(collection []item.Item, item *item.Item) []item.Item {
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
