package main

import (
	"fmt"

	"github.com/kscarlett/muzsh/colours"
	"github.com/kscarlett/muzsh/command"
	"github.com/kscarlett/muzsh/game"
	"github.com/kscarlett/muzsh/player"
	"github.com/kscarlett/muzsh/session"
)

func main() {
	initGame()
	var intent, target string

	fmt.Fprintf(colours.StdOut, "Welcome to %s\n%s\n",
		colours.Zombie("MUZSH"),
		colours.Zombie("[¬º-°]¬    [¬°-º]¬ [¬º-°]¬  [¬°-°]¬"))

	fmt.Fprintf(colours.StdOut, "The %s are approaching!\n",
		colours.Zombie("zombie hordes"))

	fmt.Fprintf(colours.StdOut, "Tell me, stranger, what is your %s?\n",
		colours.Variable("name"))
	session.Player.Name = game.GetInput()

	fmt.Fprintf(colours.StdOut, "Okay %s, we need to get out of here now.\n",
		colours.Variable(session.Player.Name))

	for intent, target = game.GetCommand(); intent != "quit"; intent, target = game.GetCommand() {
		command.HandleCommand(intent, target)
	}
}

func initGame() {
	game.SetupInput()
	session.Player = player.NewPlayer()
}
