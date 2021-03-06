package game

import (
	"fmt"

	"github.com/kscarlett/muzsh/colours"
	"github.com/kscarlett/muzsh/item"
	"github.com/kscarlett/muzsh/util"
)

// Room defines a room in the game, along with its contents.
type Room struct {
	Name       string
	Visual     string
	Contents   []item.Item
	Neighbours RoomLinks
	// Maybe add special room hazards and events
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
	fmt.Fprintf(colours.StdOut, "You %s %s.\n",
		util.RandomString([]string{"look around", "inspect", "scan", "search"}),
		colours.Location(r.Name))

	fmt.Fprintf(colours.StdOut, "%s\n", r.Visual)

	if len(r.Contents) > 0 {
		r.printRoomItems()
	} else {
		fmt.Fprintf(colours.StdOut, "%s\n", util.RandomString([]string{"There is nothing you can pick up here.", "This place has been looted long ago.", "There's nothing that can be of use to you here."}))
	}

	r.printRoomLinks()
}

// PrintRoomVisual prints the visual text of the room
func (r *Room) PrintRoomVisual() {
	fmt.Fprintf(colours.StdOut, "%s\n", r.Visual)
}

func (r *Room) printRoomItems() {
	fmt.Fprintf(colours.StdOut, "You also see ")

	for i, item := range r.Contents {
		if i < 0 {
			fmt.Fprintf(colours.StdOut, ", ")
		}

		fmt.Fprintf(colours.StdOut, "%s %s", colours.Item(item.NameArticle), colours.Item(item.Name))
	}

	fmt.Fprintf(colours.StdOut, ".\n")
}

func (r *Room) printRoomLinks() {
	printAnd := false

	fmt.Fprintf(colours.StdOut, "From here you can go ")

	if r.Neighbours.North != nil {
		fmt.Fprintf(colours.StdOut, "%s to %s",
			colours.Variable("North"),
			colours.Location(r.Neighbours.North.Name))
		printAnd = true
	}

	if r.Neighbours.East != nil {
		if printAnd {
			fmt.Fprintf(colours.StdOut, ", ")
			printAnd = false
		}

		fmt.Fprintf(colours.StdOut, "%s to %s",
			colours.Variable("East"),
			colours.Location(r.Neighbours.East.Name))
		printAnd = true
	}

	if r.Neighbours.South != nil {
		if printAnd {
			fmt.Fprintf(colours.StdOut, ", ")
			printAnd = false
		}

		fmt.Fprintf(colours.StdOut, "%s to %s",
			colours.Variable("South"),
			colours.Location(r.Neighbours.South.Name))
		printAnd = true
	}

	if r.Neighbours.West != nil {
		if printAnd {
			fmt.Fprintf(colours.StdOut, ", ")
			printAnd = false
		}

		fmt.Fprintf(colours.StdOut, "%s to %s",
			colours.Variable("West"),
			colours.Location(r.Neighbours.West.Name))
		printAnd = true
	}

	fmt.Fprintf(colours.StdOut, ".\n")
}
