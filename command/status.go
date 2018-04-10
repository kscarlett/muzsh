package command

import (
	"fmt"

	"github.com/kscarlett/muzsh/colours"
	"github.com/kscarlett/muzsh/player"
	"github.com/kscarlett/muzsh/global"
)

// StatusCommand provides the type for the command to be called
type StatusCommand struct{}

// Execute fulfills the command interface for the command pattern
func (s *StatusCommand) Execute() {
	printStatus(global.Player)
}

func printStatus(p *player.Player) {
	fmt.Printf("You are %s. You are also %s and %s",
		colours.Health(healthString(p)),
		colours.Hunger(hungerString(p)),
		colours.Thirst(thirstString(p)))
}

func healthString(p *player.Player) string {
	health := p.Health

	switch {
	case health < 10:
		return "almost dead"
	case health < 25:
		return "extremely injured"
	case health < 50:
		return "very hurt"
	case health < 75:
		return "hurt"
	case health < 90:
		return "slightly bruised"
	default:
		return "feeling healthier than ever"
	}
}

func hungerString(p *player.Player) string {
	hunger := p.Hunger

	switch {
	case hunger < 20:
		return "starving"
	case hunger < 40:
		return "very hungry"
	case hunger < 60:
		return "hungry"
	case hunger < 80:
		return "slightly hungry"
	default:
		return "sated"
	}
}

func thirstString(p *player.Player) string {
	thirst := p.Thirst

	switch {
	case thirst < 20:
		return "parched"
	case thirst < 40:
		return "very thirsty"
	case thirst < 60:
		return "thirsty"
	case thirst < 80:
		return "slightly thirsty"
	default:
		return "quenched"
	}
}
