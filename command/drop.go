package command

import (
	"fmt"
	"strings"

	"github.com/kscarlett/muzsh/colours"
	"github.com/kscarlett/muzsh/item"
	"github.com/kscarlett/muzsh/player"
	"github.com/kscarlett/muzsh/session"
	"github.com/kscarlett/muzsh/util"
)

// DropCommand provides the type for the command to be called
type DropCommand struct{}

// Execute fulfills the command interface for the command pattern
func (t *DropCommand) Execute(cmd, target string) {
	drop(session.Player, target)
}

func drop(p *player.Player, target string) {
	var targetItem *item.Item

	for _, item := range p.Inventory.Items {
		if strings.ToLower(item.Name) == target || strings.ToLower(item.NameArticle+" "+item.Name) == target {
			targetItem = &item
			break
		}
	}

	if targetItem == nil {
		fmt.Fprintf(colours.StdOut, "There is no %s to take.\n",
			colours.Item(target))
		return
	}

	p.CurrentRoom.Contents = append(p.CurrentRoom.Contents, *targetItem)
	p.Inventory.Items = util.RemoveItem(p.Inventory.Items, targetItem)

	fmt.Fprintf(colours.StdOut,
		"You have dropped %s %s.\n",
		colours.Item(targetItem.NameArticle),
		colours.Item(targetItem.Name))
}
