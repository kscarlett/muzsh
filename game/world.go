package game

// TODO: find a way to make the world in JSON and deserialize it at startup

// World holds the game world, which consists of multiple interconnected rooms
type World struct {
	Rooms []*Room
}

func NewWorld() (world *World, startingRoom *Room) {
	startingRoom = &Room{Name: "the abandoned parking lot", Visual: "The large, open space is littered with empty, abandoned cars. They aren't of much use anymore.", Contents: []Item{}}
	otherRoom := &Room{ Name: "the shopping mall entrance",
		Visual: "Broken glass litters the entrance to the shopping mall. Thousands of people used to come here every day. Now you'll only find the walking dead or a bunch of looters if you're 'lucky' enough.",
		Contents: []Item{Item{
			NameArticle: "a",
			Name: "can of beans",
			ExamineText: "Many people's favourite food for during the apocalypse. You may thank the paranoid people who'd been stockpiling this stuff for ages. Doesn't taste too bad, but it's not the nicest meal either.",
			Durability: 1,
			Type: Food}}}

	startingRoom.Neighbours.East = otherRoom
	otherRoom.Neighbours.West = startingRoom

	return &World{[]*Room{startingRoom, otherRoom}}, startingRoom
}