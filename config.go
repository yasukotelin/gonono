package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

const configName = ".gonono.json"

type config struct {
	Path   string `json:"path"`
	Editor string `json:"editor"`
}

func getConfigPath() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configName), nil
}

func openConfigFile() (*config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	conf, err := openConfig(configPath)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func openConfig(path string) (*config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var conf config
	err = json.NewDecoder(file).Decode(&conf)

	return &conf, err
}
