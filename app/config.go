package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/hytzongxuan/Codeforces-Hacker/common"
)

func readConfig(configFilePath string) ([]byte, error) {
	configData, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		return nil, err
	}

	return configData, nil
}

func LoadConfig(configFilePath string) Config {
	configData, err := readConfig(configFilePath)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	config := Config{}
	err = json.Unmarshal(configData, &config)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return config
}
