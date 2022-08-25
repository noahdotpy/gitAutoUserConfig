package main

import (
	"errors"
	"fmt"
	"gitAutoUserConfig/configUtils"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	userConfigDir, _ = os.UserConfigDir()
	configFile       = configUtils.GetConfigAsStruct(userConfigDir + "/gitUserConfig/config.json")
)

func main() {

	fmt.Println("-- Choices: --")
	getPossibleChoices()
	for i, v := range getPossibleChoices() {
		fmt.Println(i, v)
	}

	fmt.Print("\nChoice: ")

	var rawInput string
	fmt.Scan(&rawInput)
	err := handleUserChoice(rawInput)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func handleUserChoice(input string) error {
	if strings.ToLower(input) == "new" {
		handleMakeNewEntry()
		return nil
	}

	var choice int
	choice, err := strconv.Atoi(input)
	if err != nil {
		return err
	}

	if choice >= len(configFile.Choices) {
		return errors.New("Entry not found.")
	}
	_, err = addToLocalGitConfig("user.name", configFile.Choices[choice].Name)
	if err != nil {
		return err
	}

	_, err = addToLocalGitConfig("user.email", configFile.Choices[choice].Email)
	if err != nil {
		return err
	}

	return nil
}

func handleMakeNewEntry() {
	fmt.Println("Making new config entry.")
}

// TODO: Handle errors better by outputting the git output
func addToLocalGitConfig(key string, value string) (io.ReadCloser, error) {
	cmd := exec.Command("git", "config", "--local", key, value)
	out, err := cmd.StdoutPipe()

	return out, err
}

func getPossibleChoices() (rslt []string) {
	for i, v := range configFile.Choices {
		rslt = append(rslt, fmt.Sprintf("  * %v: %v - %v\n", strconv.Itoa(i), v.Name, v.Email))
	}
	rslt = append(rslt, "  * new: Make new entry.")
	return rslt
}
