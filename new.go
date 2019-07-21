package main

import (
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

	fmt.Println("(Use Ctrl + C to cancel and go back to the console)")
	inputTitle := readline("Title: ")

	createDir := filepath.Join(conf.Path, newFlagDir, formatToDirectoryName(inputTitle))

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

	if inputTitle != "" {
		h1Msg := fmt.Sprintf("# %s", inputTitle)
		_, err = fmt.Fprint(file, h1Msg)
		if err != nil {
			return err
		}
	}

	if !newFlagOpen {
		fmt.Println("created")
		return nil
	}

	fmt.Println("created and open it now")
	cmd := exec.Command(conf.Editor, indexFile)
	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}

func formatToDirectoryName(s string) string {
	t := time.Now()
	if s == "" {
		return fmt.Sprintf("%s", t.Format("2006-01-02"))
	} else {
		return fmt.Sprintf("%s-%s", t.Format("2006-01-02"), strings.Join(strings.Split(s, " "), "-"))
	}
}
