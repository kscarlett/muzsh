package command

import (
	"fmt"

	"github.com/kscarlett/muzsh/colours"
)

var commands map[string]Command
var aliases map[string]string

// Command provides an interface for commands
type Command interface {
	Execute(name, target string)
}

// HandleCommand calls a command function associated with the given name
func HandleCommand(cmd, target string) {
	if selected := commands[cmd]; selected == nil {
		// TODO: handle aliases as well
		fmt.Fprintf(colours.StdOut, "You don't know how to %s!\n", colours.Action(cmd))
	} else {
		selected.Execute(cmd, target)
	}
}

// RegisterCommand registers a command to the commands map
func RegisterCommand(name string, cmd Command) {
	// not sure if this will work well
	commands[name] = cmd
}

// PremapAll maps a set of built-in commands to their canonical names without having to call RegisterCommand on them all individually
func PremapAll() {
	premapCommands()
	premapAliases()
}

func premapCommands() {
	commands = map[string]Command{
		"status": &StatusCommand{},
		"take": &TakeCommand{},
		"drop": &DropCommand{},
		"go": &GoCommand{},
		"north": &GoCommand{},
		"east": &GoCommand{},
		"south": &GoCommand{},
		"west": &GoCommand{},
		"use": &UseCommand{},
		"eat": &UseCommand{},
		"drink": &UseCommand{},
		"shoot": &UseCommand{},
		"hit": &UseCommand{},
		"inventory": &InventoryCommand{},
		"inspect": &InspectCommand{},
	}
}

func premapAliases() {
	aliases = map[string]string{
		"grab": "take",
	}
}