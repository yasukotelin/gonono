package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func validFinderCmd(conf *config) error {
	switch {
	case conf.Path == "":
		return fmt.Errorf("%s: path is empty", configName)
	case conf.Editor == "":
		return fmt.Errorf("%s: editor is empty", configName)
	}
	return nil
}

func runFinder(c *cli.Context) error {
	conf, err := openConfigFile()
	if err != nil {
		return err
	}

	err = validFinderCmd(conf)
	if err != nil {
		return err
	}

	if err := os.Chdir(conf.Path); err != nil {
		return err
	}

	files, err := readAllFiles(conf.Path)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := newFinderCmd(files, &buf).Run(); err != nil {
		return err
	}

	if err = newCmd(fmt.Sprintf("%s %s", conf.Editor, buf.String()), nil, nil).Start(); err != nil {
		return err
	}

	return nil
}

func readAllFiles(path string) ([]string, error) {
	var paths []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return paths, err
	}
	return paths, nil
}
