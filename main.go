package main

import (
	"log"
	"os"

	"github.com/Mopip77/screenshot-handler/command"
	"github.com/Mopip77/screenshot-handler/config"
	"github.com/Mopip77/screenshot-handler/consts"
	"github.com/Mopip77/screenshot-handler/infra/output"

	"github.com/urfave/cli/v2"
)

func main() {
	if err := config.InitConfig(); err != nil {
		output.RedFmt.Println("load config failed,", err)
		return
	}

	app := &cli.App{
		Name:            "screenshot-handler",
		HelpName:        "sch",
		Usage:           "It provides a set of image-handling tools to handle screenshot from clipboard.",
		HideHelpCommand: true,
		Commands:        command.BuildCommands(),
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "from-dir",
				Aliases: []string{"d"},
				Usage:   "load latest screenshot from default folder set in " + consts.CONFIG_FILE_PATH,
			},
			&cli.StringFlag{
				Name:  "file",
				Usage: "load screenshot by filepath.",
			},
		},
		ExitErrHandler: func(ctx *cli.Context, err error) {
			if err != nil {
				output.RedFmt.Println(err)
				os.Exit(1)
			}
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
