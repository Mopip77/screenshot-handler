package command

import (
	"io/ioutil"
	"screenshot-handler/infra/output"
	"screenshot-handler/util"

	"github.com/urfave/cli/v2"
)

type ImageCommandContext struct {
	Context      *cli.Context
	ImagePath    string
	ImageContent []byte
}

type ImageCommand interface {
	ExecuteCommand(ctx ImageCommandContext) error
	ValidateRequiredConfig(ctx ImageCommandContext) error
	// provide info about command for cli
	GetCommandName() string
	GetCommandHelpName() []string
	GetUsage() string
}

var (
	commands = []ImageCommand{
		ConvertToBase64Command{},
	}
)

func BuildCommands() []*cli.Command {
	var result []*cli.Command
	for _, command := range commands {
		result = append(result, &cli.Command{
			Name:    command.GetCommandName(),
			Aliases: command.GetCommandHelpName(),
			Usage:   command.GetUsage(),
			Action:  ExecuteCommand,
		})
	}
	return result
}

func ExecuteCommand(ctx *cli.Context) error {
	// load file
	imagePath, err := util.GetLatestScreenshotPath(ctx)
	if err != nil {
		output.RedFmt.Println("load screenshot file failed,", err)
		return err
	}
	imageContent, err := ioutil.ReadFile(imagePath)
	if err != nil {
		output.RedFmt.Printf("read screenshot failed, %s\n", err.Error())
		return err
	}

	imageCommandCtx := ImageCommandContext{
		Context:      ctx,
		ImagePath:    imagePath,
		ImageContent: imageContent,
	}

	for _, command := range commands {
		if command.GetCommandName() == ctx.Command.Name {
			if err := command.ValidateRequiredConfig(imageCommandCtx); err != nil {
				return err
			}
			return command.ExecuteCommand(imageCommandCtx)
		}
	}

	return nil
}
