package main

import (
	"errors"
	"fmt"
	"gauc/configUtils"
	"io"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("-- Choices: --")
	getPossibleChoices()
	for _, v := range getPossibleChoices() {
		fmt.Println(v)
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

	if choice >= len(configUtils.ConfigObj.Choices) {
		return errors.New("Entry not found.")
	}
	_, err = addToLocalGitConfig(
		"user.name",
		configUtils.ConfigObj.Choices[choice].Name,
	)
	if err != nil {
		return err
	}

	_, err = addToLocalGitConfig(
		"user.email",
		configUtils.ConfigObj.Choices[choice].Email,
	)
	if err != nil {
		return err
	}

	return nil
}

func handleMakeNewEntry() {

	fmt.Print("Name: ")
	var name string
	fmt.Scan(&name)

	fmt.Print("Email: ")
	var email string
	fmt.Scan(&email)

	configUtils.AppendChoiceToConfig(name, email)
}

// TODO: Handle errors better by outputting the git output
func addToLocalGitConfig(key string, value string) (io.ReadCloser, error) {
	cmd := exec.Command("git", "config", "--local", key, value)
	out, err := cmd.StdoutPipe()

	return out, err
}

func getPossibleChoices() (rslt []string) {
	for i, v := range configUtils.ConfigObj.Choices {
		rslt = append(rslt, fmt.Sprintf("  * %v: %v - %v\n", strconv.Itoa(i), v.Name, v.Email))
	}
	rslt = append(rslt, "  * new: Make new entry.")
	return rslt
}
