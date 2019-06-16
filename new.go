package main

import (
	"bufio"
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

	if err := os.Chdir(conf.Path); err != nil {
		return err
	}

	inputTitle := readLine("Title: ")

	dirName := formatToDirectoryName(inputTitle)

	if err = os.Mkdir(dirName, os.ModeDir); err != nil {
		return err
	}
	indexFile := filepath.Join(dirName, "index.md")
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

func readLine(askMsg string) string {
	fmt.Print(askMsg)
	scan := bufio.NewScanner(os.Stdin)
	var s string
	if scan.Scan() {
		s = scan.Text()
	}
	return s
}

func formatToDirectoryName(s string) string {
	fmt.Println(s)
	t := time.Now()
	return fmt.Sprintf("%s-%s", t.Format("2006-01-02"), strings.Join(strings.Split(s, " "), "-"))
}
