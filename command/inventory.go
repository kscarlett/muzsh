package command

import (
	"fmt"

	"github.com/kscarlett/muzsh/colours"
	"github.com/kscarlett/muzsh/player"
	"github.com/kscarlett/muzsh/session"
)

// InventoryCommand provides the type for the command to be called
type InventoryCommand struct{}

// Execute fulfills the command interface for the command pattern
func (t *InventoryCommand) Execute(cmd, target string) {
	inventory(session.Player)
}

func inventory(p *player.Player) {
	if len(p.Inventory.Items) <= 0 {
		fmt.Fprintf(colours.StdOut, "Your inventory is empty.\n")
	}
	
	fmt.Fprintf(colours.StdOut, "You have the following in your inventory: ")

	for i, item := range p.Inventory.Items {
		if i > 0 {
			fmt.Fprintf(colours.StdOut, ", ")
		}

		fmt.Fprintf(colours.StdOut, "%s %s",
			colours.Item(item.NameArticle),
			colours.Item(item.Name))
	}

	fmt.Fprintf(colours.StdOut, ".\n")
}