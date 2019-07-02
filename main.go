package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	newFlagOpen bool
	newFlagDir  string
)

func main() {
	app := cli.NewApp()
	app.Name = "gonono"
	app.Version = "1.2.0"
	app.Usage = "provides the note environment with your favorite editor. "
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Usage:   "creates the empty config file",
			Aliases: []string{"i"},
			Action:  runInit,
		},
		{
			Name:    "open",
			Usage:   "opens the note directory with explorer",
			Aliases: []string{"o"},
			Action:  runOpen,
		},
		{
			Name:    "new",
			Usage:   "creates the new note",
			Aliases: []string{"n"},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "open, o",
					Usage:       "opens after created",
					Destination: &newFlagOpen,
				},
				cli.StringFlag{
					Name:        "dir, d",
					Usage:       "create to the dir",
					Destination: &newFlagDir,
				},
			},
			Action: runNew,
		},
		{
			Name:    "finder",
			Usage:   "find the created note with fzf",
			Aliases: []string{"f"},
			Action:  runFinder,
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
		return fmt.Errorf("%s: path is empty", configName)
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
	command := fmt.Sprintf("%s .", conf.Editor)
	if err = newCmd(command, nil, nil).Run(); err != nil {
		return err
	}

	return nil
}
