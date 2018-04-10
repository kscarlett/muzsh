package command

import (
	"fmt"

	"github.com/kscarlett/muzsh/colours"
)

var commands map[string]Command

type Command interface {
	Excecute()
}

func HandleCommand(cmd string) {
	if selected := commands[cmd]; selected == nil {
		fmt.Printf("You don't know how to %s", colours.Action(selected))
	} else {
		selected.Excecute()
	}
}

func RegisterCommand(name string, cmd Command) {
	// not sure if this will work well
	commands[name] = cmd
}
