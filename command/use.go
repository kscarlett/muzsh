package command

import (
	"fmt"
	"strings"

	"github.com/kscarlett/muzsh/colours"
	"github.com/kscarlett/muzsh/player"
	"github.com/kscarlett/muzsh/session"
)

// UseCommand provides the type for the command to be called
type UseCommand struct{}

// Execute fulfills the command interface for the command pattern
func (t *UseCommand) Execute(name, target string) {
	use(session.Data.Player, name, target)
}

func use(p *player.Player, action, target string) {
	fmt.Fprintf(colours.StdOut, "%s\n", colours.Zombie("test"))

	// TODO: implement
	switch {
	case strings.Contains(action, "eat"):
		// if target item != itemType.Food || itemType.Medical
		// else
	case strings.Contains(action, "drink"):
		// if target item != itemType.Drink || itemType.Medical
		// else
	case strings.Contains(action, "shoot"):
		// if target item != itemType.Ranged || target == enemy
		// else
	case strings.Contains(action, "hit"):
		// if target item != itemType.Melee || target == enemy
		// else
	default:
		// use (medical, misc, etc)
	}
}
