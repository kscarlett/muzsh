package command

import (
	"fmt"

	"github.com/kscarlett/muzsh/colours"
)

var commands map[string]Command

// Command provides an interface for commands
type Command interface {
	Excecute()
}

// HandleCommand calls a command function associated with the given name
func HandleCommand(cmd, target string) {
	if selected := commands[cmd]; selected == nil {
		fmt.Fprintf(colours.StdOut, "You don't know how to %s!\n", colours.Action(selected))
	} else {
		selected.Excecute()
	}
}

// RegisterCommand registers a command to the commands map
func RegisterCommand(name string, cmd Command) {
	// not sure if this will work well
	commands[name] = cmd
}
