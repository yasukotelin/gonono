package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gonono"
	app.Version = "0.1.0"
	app.Usage = "provides the note environment with your favorite editor. "
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "initialize config file",
			Action: runInit,
		},
		{
			Name:    "open",
			Usage:   "open with the explorer",
			Aliases: []string{"o"},
			Action:  runOpen,
		},
		{
			Name:    "new",
			Usage:   "create the new content",
			Aliases: []string{"n"},
			Action:  runNew,
		},
		{
			Name:    "edit",
			Usage:   "edit the created content",
			Aliases: []string{"e"},
		},
	}
	app.Action = runGonono

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func validGononoCmd(conf *config) error {
	switch {
	case conf.Path == "":
		return fmt.Errorf("%s: Path is empty", configName)
	case conf.Editor == "":
		return fmt.Errorf("%s: editor is empty", configName)
	}
	return nil
}

func runGonono(c *cli.Context) error {
	conf, err := openConfigFile()
	if err != nil {
		return err
	}

	err = validGononoCmd(conf)
	if err != nil {
		return err
	}

	if err := os.Chdir(conf.Path); err != nil {
		return err
	}

	// カレントディレクトリでエディタを開くために"."を引数に指定する
	cmd := exec.Command(conf.Editor, ".")
	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}
