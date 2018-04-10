package game

import (
	"fmt"

	"github.com/kscarlett/muzsh/colours"
)

// Room defines a room in the game, along with its contents.
type Room struct {
	Name       string
	Visual     string
	Contents   []Item
	Neighbours RoomLinks
}

// RoomLinks define the connections between different rooms.
type RoomLinks struct {
	North *Room
	East  *Room
	South *Room
	West  *Room
}

// InspectRoom prints the properties of the room (visual text, items, links)
func (r *Room) InspectRoom() {
	fmt.Printf("You look around %s.\n", colours.Location(r.Name))

	fmt.Printf("%s\n", r.Visual)

	if len(r.Contents) > 0 {
		r.printRoomItems()
	} else {
		fmt.Printf("There is nothing you can pick up here.\n")
	}

	r.printRoomLinks()
}

func (r *Room) printRoomItems() {
	fmt.Printf("You also see ")

	for i, item := range r.Contents {
		fmt.Printf("%s %s", colours.Item(item.NameArticle), colours.Item(item.Name))

		if i < len(r.Contents) {
			fmt.Printf(", ")
		}
	}

	fmt.Printf(".\n")
}

func (r *Room) printRoomLinks() {
	printAnd := false

	fmt.Printf("From here you can go ")

	if r.Neighbours.North != nil {
		fmt.Printf("%s to %s", colours.Variable("North"), colours.Location(r.Neighbours.North.Name))
		printAnd = true
	}

	if r.Neighbours.East != nil {
		if printAnd {
			fmt.Printf(", ")
			printAnd = false
		}

		fmt.Printf("%s to %s", colours.Variable("East"), colours.Location(r.Neighbours.East.Name))
		printAnd = true
	}

	if r.Neighbours.South != nil {
		if printAnd {
			fmt.Printf(", ")
			printAnd = false
		}

		fmt.Printf("%s to %s", colours.Variable("South"), colours.Location(r.Neighbours.South.Name))
		printAnd = true
	}

	if r.Neighbours.West != nil {
		if printAnd {
			fmt.Printf(", ")
			printAnd = false
		}

		fmt.Printf("%s to %s", colours.Variable("West"), colours.Location(r.Neighbours.West.Name))
		printAnd = true
	}

	fmt.Printf(".\n")
}
