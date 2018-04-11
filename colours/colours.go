package colours

import (
	"github.com/fatih/color"
)

var (
	StdOut = color.Output

	Text     = color.New(color.FgWhite).SprintFunc()
	Location = color.New(color.FgHiYellow).SprintFunc()
	Item     = color.New(color.FgCyan).SprintFunc()
	Zombie   = color.New(color.FgGreen).SprintFunc()
	ZombieEx = color.New(color.FgHiGreen).SprintFunc()
	Variable = color.New(color.FgMagenta).Add(color.Underline).SprintFunc()
	Prompt   = color.New(color.FgMagenta).SprintFunc()
	Health   = color.New(color.FgRed).Add(color.Italic).SprintFunc()
	Hunger   = color.New(color.FgHiGreen).Add(color.Italic).SprintFunc()
	Thirst   = color.New(color.FgBlue).Add(color.Italic).SprintFunc()
	Action   = color.New(color.FgHiWhite).Add(color.Italic).SprintFunc()
)
