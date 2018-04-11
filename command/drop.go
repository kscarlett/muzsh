package command

import (
	"fmt"

	"github.com/kscarlett/muzsh/colours"
	"github.com/kscarlett/muzsh/player"
	"github.com/kscarlett/muzsh/session"
)

// DropCommand provides the type for the command to be called
type DropCommand struct{}

// Execute fulfills the command interface for the command pattern
func (t *DropCommand) Execute(target string) {
	drop(session.Player)
}

func drop(p *player.Player) {
	fmt.Fprintf(colours.StdOut, "%s\n", colours.Zombie("test"))
}