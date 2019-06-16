package main

import (
	"encoding/json"
	"os"

	"github.com/urfave/cli"
)

func runInit(c *cli.Context) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	if err := createEmptyConfig(configPath); err != nil {
		return err
	}
	return nil
}

func createEmptyConfig(path string) error {
	bytes, err := createEmptySettingJSONBytes()
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	file.WriteString(string(bytes))

	return nil
}

func createEmptySettingJSONBytes() ([]byte, error) {
	config := config{
		Path:    "",
		Editor:  "",
		OpenCmd: "",
	}
	return json.MarshalIndent(config, "", "    ")
}