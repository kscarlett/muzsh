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

// TakeCommand provides the type for the command to be called
type TakeCommand struct{}

// Execute fulfills the command interface for the command pattern
func (t *TakeCommand) Execute(name, target string) {
	take(session.Player, target)
}

func take(p *player.Player, target string) {
	var targetItem *item.Item

	for _, item := range p.CurrentRoom.Contents {
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

	//p.Inventory.Items.add(targetItem)
	p.Inventory.Items = append(p.Inventory.Items, *targetItem)

	//p.CurrentRoom.Contents.remove(targetItem)
	p.CurrentRoom.Contents = util.RemoveItem(p.CurrentRoom.Contents, targetItem)

	fmt.Fprintf(colours.StdOut,
		"You have taken %s %s and safely put it in your inventory.\n",
		colours.Item(targetItem.NameArticle),
		colours.Item(targetItem.Name))
}
