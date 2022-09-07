package configUtils

import (
	"encoding/json"
	"os"
)

var (
	UserConfigDir, _ = os.UserConfigDir()
	ConfigFilePath   = UserConfigDir + "/gitUserConfig/config.json"
	ConfigObj, _     = GetConfigAsStruct(ConfigFilePath)
)

type Config struct {
	Choices []Choice `json:"choices"`
}
type Choice struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetConfigAsStruct(filePath string) (Config, error) {
	file, _ := os.ReadFile(filePath)

	var payload Config
	err := json.Unmarshal(file, &payload)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func SaveConfig(filePath string, data Config) error {

	newFile, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, newFile, 0666)
	if err != nil {
		return err
	}

	return nil
}

func AppendChoiceToConfig(n, e string) error {
	newChoice := Choice{
		Name:  n,
		Email: e,
	}

	ConfigObj.Choices = append(ConfigObj.Choices, newChoice)
	err := SaveConfig(ConfigFilePath, ConfigObj)
	if err != nil {
		return err
	}

	return nil
}
