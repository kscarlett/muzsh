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

	printBanner()

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
	command.PremapAll()
	session.Player = player.NewPlayer()
	session.World, session.Player.CurrentRoom = game.NewWorld()
}

func printBanner() {
	fmt.Fprintf(colours.StdOut,
		"\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\n\t%s\n\t%s\n\n",
		colours.ZombieEx("                    Welcome  to                    "),
		colours.Zombie("@@@@@@@@@@   @@@  @@@  @@@@@@@@   @@@@@@   @@@  @@@"),
		colours.Zombie("@@@@@@@@@@@  @@@  @@@  @@@@@@@@  @@@@@@@   @@@  @@@"),
		colours.Zombie("@@! @@! @@!  @@!  @@@       @@!  !@@       @@!  @@@"),
		colours.Zombie("!@! !@! !@!  !@!  @!@      !@!   !@!       !@!  @!@"),
		colours.Zombie("@!! !!@ @!@  @!@  !@!     @!!    !!@@!!    @!@!@!@!"),
		colours.Zombie("!@!   ! !@!  !@!  !!!    !!!      !!@!!!   !!!@!!!!"),
		colours.Zombie("!!:     !!:  !!:  !!!   !!:           !:!  !!:  !!!"),
		colours.Zombie(":!:     :!:  :!:  !:!  :!:           !:!   :!:  !:!"),
		colours.Zombie(":::     ::   ::::: ::   :: ::::  :::: ::   ::   :::"),
		colours.Zombie(" :      :     : :  :   : :: : :  :: : :     :   : :"),
		colours.Zombie("[¬º-°]¬    [¬°-º]¬ [¬º-°]¬  [¬°-°]¬ [¬°-º]¬ [¬º-°]¬"),
		colours.ZombieEx("muzsh.kscarlett.com           Kellen Scarlett, 2018"))
}