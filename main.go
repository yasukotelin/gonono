package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

const configName = ".gonono.json"

type config struct {
	Path string   `json:"path"`
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

func main() {
	configPath, err := getConfigPath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// TODO create the new config file if command is specified `init`
	// if !isFileExists(configPath) {
	// 	if err := createEmptyConfig(configPath); err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}
	// }

	conf, err := openConfig(configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = validate(conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := os.Chdir(conf.Path); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := exec.Command(conf.Cmd, conf.Args...)
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func getConfigPath() (string, error) {
	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, configName), nil
}

func isFileExists(path string) bool {
	if _, err := os.Open(path); err != nil {
		return false
	}
	return true
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

func validate(conf *config) error {
	switch {
	case conf.Path == "":
		return fmt.Errorf("%s: Path is empty", configName)
	case conf.Cmd == "":
		return fmt.Errorf("%s: Cmd is empty", configName)
	case conf.Args == nil:
		return fmt.Errorf("%s: Args are empty", configName)
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
	config := config{Args: []string{}}
	return json.MarshalIndent(config, "", "    ")
}
