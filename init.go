package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func runInit(c *cli.Context) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	if existsFile(configPath) {
		return fmt.Errorf("%v already exists.", configPath)
	}

	if err := createEmptyConfig(configPath); err != nil {
		return err
	}
	fmt.Printf("created the %v\n", configPath)
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
	_, err = file.WriteString(string(bytes))
	if err != nil {
		return err
	}

	return nil
}

func createEmptySettingJSONBytes() ([]byte, error) {
	config := config{
		Path:   "",
		Editor: "",
	}
	return json.MarshalIndent(config, "", "    ")
}

func existsFile(path string) bool {
	_, e := os.Stat(path)
	return e == nil
}
