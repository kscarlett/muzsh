package command

import (
	"fmt"
	"strings"
	"log"

	"github.com/kscarlett/muzsh/colours"
	"github.com/kscarlett/muzsh/player"
	"github.com/kscarlett/muzsh/session"
	"github.com/kscarlett/muzsh/game"
)

// GoCommand provides the type for the command to be called
type GoCommand struct{}

// Execute fulfills the command interface for the command pattern
func (t *GoCommand) Execute(name, target string) {
	goTo(session.Player, target, name)
}

func goTo(p *player.Player, target, command string) {
	ok := false

	switch {
	case strings.Contains(command, "north"):
		ok = tryGoToRoom(p, 0)
		target = "North"
	case strings.Contains(command, "east"):
		ok = tryGoToRoom(p, 1)
		target = "East"
	case strings.Contains(command, "south"):
		ok = tryGoToRoom(p, 2)
		target = "South"
	case strings.Contains(command, "west"):
		ok = tryGoToRoom(p, 3)
		target = "West"
	case strings.Contains(target, "north"):
		ok = tryGoToRoom(p, 0)
		target = "North"
	case strings.Contains(target, "east"):
		ok = tryGoToRoom(p, 1)
		target = "East"
	case strings.Contains(target, "south"):
		ok = tryGoToRoom(p, 2)
		target = "South"
	case strings.Contains(target, "west"):
		ok = tryGoToRoom(p, 3)
		target = "West"
	default:
		fmt.Fprintf(colours.StdOut, "%s is not a valid direction!\n",
			colours.Variable(strings.ToTitle(target)))
		return
	}

	if !ok {
		fmt.Fprintf(colours.StdOut, "There is nowhere to go to the %s.\n",
			colours.Location(target))
	}
}

func tryGoToRoom(p *player.Player, direction int) bool {
	var targetRoom *game.Room
	
	switch direction {
	case 0:
		targetRoom = p.CurrentRoom.Neighbours.North
	case 1:
		targetRoom = p.CurrentRoom.Neighbours.East
	case 2:
		targetRoom = p.CurrentRoom.Neighbours.South
	case 3:
		targetRoom = p.CurrentRoom.Neighbours.West
	default:
		log.Panic("Invalid direction given!")
	}

	if targetRoom != nil {
		p.CurrentRoom = targetRoom
		fmt.Fprintf(colours.StdOut, "You are now in %s.\n",
			colours.Location(p.CurrentRoom.Name))
		p.CurrentRoom.PrintRoomVisual()
		return true
	}
	
	return false
}