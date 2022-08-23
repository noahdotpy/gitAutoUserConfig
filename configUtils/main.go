package configUtils

import (
	"encoding/json"
	"log"
	"os"
)

type choices struct {
	Choices []choice `json:"choices"`
}
type choice struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetConfigAsStruct(filePath string) choices {
	file, _ := os.ReadFile(filePath)

	var payload choices
	err := json.Unmarshal(file, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return payload
}
