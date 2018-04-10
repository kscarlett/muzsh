package colours

import (
	"github.com/fatih/color"
)

var Text = color.New(color.FgWhite).SprintFunc()
var Location = color.New(color.FgHiYellow).SprintFunc()
var Item = color.New(color.FgCyan).SprintFunc()
var Zombie = color.New(color.FgGreen).SprintFunc()
var Variable = color.New(color.FgMagenta).Add(color.Underline).SprintFunc()
var Prompt = color.New(color.FgMagenta).SprintFunc()
var Health = color.New(color.FgRed).Add(color.Italic).SprintFunc()
var Hunger = color.New(color.FgHiGreen).Add(color.Italic).SprintFunc()
var Thirst = color.New(color.FgBlue).Add(color.Italic).SprintFunc()
var Action = color.New(color.FgHiWhite).Add(color.Italic).SprintFunc()
