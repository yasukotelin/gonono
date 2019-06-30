package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/urfave/cli"
)

func validNewCommand(conf *config) error {
	switch {
	case conf.Path == "":
		return fmt.Errorf("%s: path is empty", configName)
	case conf.Editor == "":
		return fmt.Errorf("%s: editor is empty", configName)
	}
	return nil
}

func runNew(c *cli.Context) error {
	conf, err := openConfigFile()
	if err != nil {
		return err
	}

	err = validNewCommand(conf)
	if err != nil {
		return err
	}

	inputTitle := readline("Title: ")
	if inputTitle == "" {
		return errors.New("inputed title is empty")
	}

	createDir := filepath.Join(conf.Path, flagDir, formatToDirectoryName(inputTitle))

	for {
		in := readline(fmt.Sprintf("is this OK? (%s) [y/n] ", createDir))
		if in == "y" {
			break
		} else if in == "n" {
			return nil
		} else {
			fmt.Println("you should input the y or n")
		}
	}

	if err = os.MkdirAll(createDir, os.ModeDir); err != nil {
		return err
	}
	indexFile := filepath.Join(createDir, "index.md")
	file, err := os.Create(indexFile)
	if err != nil {
		return err
	}
	defer file.Close()

	h1Msg := fmt.Sprintf("# %s", inputTitle)
	_, err = fmt.Fprint(file, h1Msg)
	if err != nil {
		return err
	}

	cmd := exec.Command(conf.Editor, indexFile)
	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}

func formatToDirectoryName(s string) string {
	t := time.Now()
	return fmt.Sprintf("%s-%s", t.Format("2006-01-02"), strings.Join(strings.Split(s, " "), "-"))
}
