package command

import (
	"fmt"
	"strings"

	"github.com/kscarlett/muzsh/colours"
	"github.com/kscarlett/muzsh/player"
	"github.com/kscarlett/muzsh/session"
)

// InspectCommand provides the type for the command to be called
type InspectCommand struct{}

// Execute fulfills the command interface for the command pattern
func (t *InspectCommand) Execute(cmd, target string) {
	inspect(session.Data.Player, cmd, target)
}

func inspect(p *player.Player, command, target string) {

	// TODO: find other aliases for place and implement in a better way
	if strings.Contains(target, "room") || strings.Contains(target, "area") {
		p.CurrentRoom.InspectRoom()
		return
	}

	// if target cast to item in inventory, print description and uses left if not weapon or ammo left if weapon (and durability != -1 (infinite))
	// TODO: implement

	// if target cast to item in room inventory, print description and uses left if not weapon or ammo left if weapon (and durability != -1 (infinite))
	// TODO: implement

	// Check for player name last, in case the player gives themself the same name as an item
	if strings.Contains(target, "self") || strings.Contains(target, strings.ToLower(p.Name)) {
		HandleCommand("status", "")
		return
	}

	// else error

	fmt.Fprintf(colours.StdOut, "%s\n", colours.Zombie("test"))
}
