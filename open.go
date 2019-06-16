package main

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli"
)

func validOpenCmd(conf *config) error {
	switch {
	case conf.Path == "":
		return fmt.Errorf("%s: Path is empty", configName)
	case conf.OpenCmd == "":
		return fmt.Errorf("%s: Open Cmd is empty", configName)
	}
	return nil
}

func runOpen(c *cli.Context) error {
	conf, err := openConfigFile()
	if err != nil {
		return err
	}

	err = validOpenCmd(conf)
	if err != nil {
		return err
	}

	cmd := exec.Command(conf.OpenCmd, conf.Path)
	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}
