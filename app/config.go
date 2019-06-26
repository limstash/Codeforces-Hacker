package app

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/hytzongxuan/Codeforces-Hacker/common"
)

func readConfig(configFilePath string) ([]byte, error) {
	configData, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		return nil, err
	}

	return configData, nil
}

func LoadConfig(configFilePath string) (Config, error) {
	config := Config{}

	configData, err := readConfig(configFilePath)

	if err != nil {
		return config, err
	}

	err = json.Unmarshal(configData, &config)

	if err != nil {
		return config, err
	}

	return config, nil
}
