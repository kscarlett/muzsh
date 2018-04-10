package main

import (
	"fmt"

	"github.com/kscarlett/muzsh/colours"
	"github.com/kscarlett/muzsh/game"
)

func main() {
	initGame()

	fmt.Printf("Welcome to %s\n%s\n", colours.Zombie("MUZSH"), colours.Zombie("[¬º-°]¬    [¬°-º]¬ [¬º-°]¬  [¬°-°]¬"))

	fmt.Printf("The %s are approaching!\n", colours.Zombie("zombie hordes"))

	fmt.Printf("Tell me, stranger, what is your %s?\n", colours.Variable("name"))
	name := game.GetInput()

	fmt.Printf("Okay %s, we need to get out of here now.\n", colours.Variable(name))
}

func initGame() {
	game.SetupInput()
}
