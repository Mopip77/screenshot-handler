package command

import (
	"fmt"

	"github.com/Mopip77/screenshot-handler/util"

	"github.com/urfave/cli/v2"
)

var (
	commands = []ImageCommand{
		ConvertToBase64Command{},
		UploadCommand{},
		OcrCommand{},
		CopyCommand{},
		PriviewCommand{},
	}
)

func BuildCommands() []*cli.Command {
	var result []*cli.Command
	for _, command := range commands {
		result = append(result, &cli.Command{
			Name:      command.GetCommandName(),
			Aliases:   command.GetCommandHelpName(),
			Usage:     command.GetUsage(),
			Flags:     command.GetCommandFlags(),
			Category:  command.GetCategory(),
			ArgsUsage: command.GetArgUsages(),
			Action:    ExecuteCommand,
		})
	}
	return result
}

func ExecuteCommand(ctx *cli.Context) error {
	// load file
	imageName, imagePath, imageContent, fromClipboard, err := util.LoadScreenshot(ctx)
	if err != nil {
		return fmt.Errorf("load screenshot file failed, %w", err)
	}

	imageCommandCtx := ImageCommandContext{
		Context:       ctx,
		FromClipboard: fromClipboard,
		ImageName:     imageName,
		ImagePath:     imagePath,
		ImageContent:  imageContent,
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
