package game

import (
	"github.com/kscarlett/muzsh/item"
)

// TODO: find a way to make the world in JSON and deserialize it at startup
// This will require a way to make references like pointers, as rooms point to neighbours.

// World holds the game world, which consists of multiple interconnected rooms
type World struct {
	Rooms        []*Room
	StartingRoom *Room
}

func NewWorld() *World {
	parkingLot := &Room{Name: "the abandoned parking lot", Visual: "The large, open space is littered with empty, abandoned cars. They aren't of much use anymore.", Contents: []item.Item{}}
	shoppingMallEntrance := &Room{Name: "the shopping mall entrance",
		Visual: "Broken glass litters the entrance to the shopping mall. Thousands of people used to come here every day. Now you'll only find the walking dead or a bunch of looters if you're 'lucky' enough.",
		Contents: []item.Item{item.Item{
			NameArticle: "a",
			Name:        "can of beans",
			ExamineText: "Many people's favourite food for during the apocalypse. You may thank the paranoid people who'd been stockpiling this stuff for ages. Doesn't taste too bad, but it's not the nicest meal either.",
			Durability:  1,
			Type:        item.Food}}}

	parkingLot.Neighbours.East = shoppingMallEntrance
	shoppingMallEntrance.Neighbours.West = parkingLot

	world := &World{
		Rooms:        []*Room{parkingLot, shoppingMallEntrance},
		StartingRoom: parkingLot,
	}

	return world
}
