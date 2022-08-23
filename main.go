package main

import (
	"fmt"
	"gitUserConfig/configUtils"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	userConfigDir, _ = os.UserConfigDir()
	configFile       = configUtils.GetConfigAsStruct(userConfigDir + ".config/gitUserConfig/config.json")
)

func main() {

	outputPossibleChoices()

	fmt.Print("\nChoice: ")

	var rawInput string
	fmt.Scan(&rawInput)
	handleUserChoice(rawInput)
}

func handleUserChoice(input string) {
	if strings.ToLower(input) == "new" {
		handleMakeNewEntry()
		return
	}

	var choice int
	choice, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	_, err = addToLocalGitConfig("user.name", configFile.Choices[choice].Name)
	if err != nil {
		log.Fatal(err)
	}

	_, err = addToLocalGitConfig("user.email", configFile.Choices[choice].Email)
	if err != nil {
		log.Fatal(err)
	}
}

func handleMakeNewEntry() {
	fmt.Println("Making new config entry.")
}

// TODO: Handle errors better by outputting the git output
func addToLocalGitConfig(key string, value string) (io.ReadCloser, error) {
	out, err := exec.Command("git", "config", "--local", key, value).StdoutPipe()

	return out, err
}

func outputPossibleChoices() {
	fmt.Println("-- Choices: --")
	for i, v := range configFile.Choices {
		fmt.Printf("  * %v: %v %v\n", strconv.Itoa(i), v.Name, v.Email)
	}
	fmt.Println("  * new: Make new entry.")
}
