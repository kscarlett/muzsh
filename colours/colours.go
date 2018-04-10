package colours

import (
	"github.com/fatih/color"
)

var Text = color.New(color.FgWhite).SprintfFunc()
var Location = color.New(color.FgHiYellow).SprintfFunc()
var Item = color.New(color.FgCyan).SprintFunc()
var Zombie = color.New(color.FgGreen).SprintFunc()
var Variable = color.New(color.FgMagenta).Add(color.Underline).SprintfFunc()
var Prompt = color.New(color.FgMagenta).SprintfFunc()
