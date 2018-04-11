package command

import (
	"fmt"

	"github.com/kscarlett/muzsh/colours"
	"github.com/kscarlett/muzsh/player"
	"github.com/kscarlett/muzsh/session"
)

// InspectCommand provides the type for the command to be called
type InspectCommand struct{}

// Execute fulfills the command interface for the command pattern
func (t *InspectCommand) Execute(target string) {
	inspect(session.Player)
}

func inspect(p *player.Player) {
	fmt.Fprintf(colours.StdOut, "%s\n", colours.Zombie("test"))
}