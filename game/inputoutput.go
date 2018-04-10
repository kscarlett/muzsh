package game

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

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
		log.Panic(err.Error())
		return ""
	}

	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	return text
}

func GetCommand() (intent string, target string) {
	printPrompt()

	return "", ""
}

func printPrompt() {
	fmt.Printf("%s", colours.Prompt("> "))
}
