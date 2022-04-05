package command

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Mopip77/screenshot-handler/infra/output"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

type CopyCommand struct {
	abstractImageCommand
}

var (
	forceOverwrite bool
)

func (cmd CopyCommand) ExecuteCommand(ctx ImageCommandContext) error {
	dstpath := ctx.Context.Args().First()
	if dstpath == "" {
		return fmt.Errorf("dstpath is empty")
	}

	file, _ := os.Stat(dstpath)
	if file != nil {
		if !file.IsDir() && !forceOverwrite {
			return fmt.Errorf("dstpath already exists, use -f to overwrite")
		}
		if file.IsDir() {
			dstpath = filepath.Join(dstpath, ctx.ImageName)
		}
	}

	dstFile, err := os.Create(dstpath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	io.Copy(dstFile, bytes.NewReader(ctx.ImageContent))

	output.Fmt.Printf("Successfully copy to %s\n", color.GreenString(dstpath))

	return nil
}

func (cmd CopyCommand) GetCommandName() string {
	return "copy"
}

func (cmd CopyCommand) GetCommandHelpName() []string {
	return []string{"cp"}
}

func (cmd CopyCommand) GetUsage() string {
	return "copy screenshot to path"
}

func (cmd CopyCommand) GetArgUsages() string {
	return "<dstpath>"
}

func (cmd CopyCommand) GetCommandFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "f",
			Aliases:     []string{"force"},
			Value:       false,
			Usage:       "overwrite dstpath if it exists",
			Destination: &forceOverwrite,
		},
	}
}
