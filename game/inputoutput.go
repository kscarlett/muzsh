package game

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/kscarlett/muzsh/colours"
)

var reader *bufio.Reader

func SetupInput() {
	reader = bufio.NewReader(os.Stdin)
}

func GetInput() string {
	printPrompt()

	text, err := reader.ReadString('\n')
	if err != nil {
		return text
	} else {
		log.Panic(err)
		return ""
	}
}

func GetCommand() (intent string, target string) {
	printPrompt()

	return "", ""
}

func printPrompt() {
	fmt.Printf("%s", colours.Prompt("> "))
}
