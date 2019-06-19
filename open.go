package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func validOpenCmd(conf *config) error {
	switch {
	case conf.Path == "":
		return fmt.Errorf("%s: path is empty", configName)
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

	err = newOpenCmd(conf.Path).Start()
	if err != nil {
		return err
	}

	return nil
}
