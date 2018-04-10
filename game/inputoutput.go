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

// SetupInput assigns a new buffer reader for all input
func SetupInput() {
	reader = bufio.NewReader(os.Stdin)
}

// GetInput reads in a single string from the command line and retruns it
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

// GetCommand reads in a string from the command line and attempts to split it into intent (a command) and the target (location, item, enemy, etc), which it then returns
func GetCommand() (intent string, target string) {
	printPrompt()

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Panic(err.Error())
		return "", ""
	}

	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)
	text = strings.ToLower(text)
	words := strings.Split(text, " ")

	if len(words) < 1 {
		log.Panic("No command given")
		return "", ""
	}

	intent = words[0]

	if len(words) == 1 {
		return intent, ""
	}

	target = strings.Join(words[1:], " ")

	return
}

func printPrompt() {
	fmt.Printf("%s", colours.Prompt("> "))
}
