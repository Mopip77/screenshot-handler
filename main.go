package main

import (
	"log"
	"os"
	"screenshot-handler/command"
	"screenshot-handler/config"
	"screenshot-handler/infra/output"

	"github.com/urfave/cli/v2"
)

func main() {
	if err := config.InitConfig(); err != nil {
		output.RedFmt.Println("load config failed,", err)
		return
	}

	app := &cli.App{
		Name:     "screenshot-handler",
		HelpName: "sch",
		Usage:    "It provides a set of image-handling tools to handle screenshot after saved.",
		Commands: command.BuildCommands(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
