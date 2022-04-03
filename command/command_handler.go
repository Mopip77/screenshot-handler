package command

import (
	"io/ioutil"
	"screenshot-handler/infra/output"
	"screenshot-handler/util"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	commands = []ImageCommand{
		ConvertToBase64Command{},
		OcrCommand{},
	}
)

func BuildCommands() []*cli.Command {
	var result []*cli.Command
	for _, command := range commands {
		result = append(result, &cli.Command{
			Name:     command.GetCommandName(),
			Aliases:  command.GetCommandHelpName(),
			Usage:    command.GetUsage(),
			Flags:    command.GetCommandFlags(),
			Category: command.GetCategory(),
			Action:   ExecuteCommand,
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
	splits := strings.Split(imagePath, "/")
	imageName := splits[len(splits)-1]
	imageContent, err := ioutil.ReadFile(imagePath)
	if err != nil {
		output.RedFmt.Printf("read screenshot failed, %s\n", err.Error())
		return err
	}

	imageCommandCtx := ImageCommandContext{
		Context:      ctx,
		ImageName:    imageName,
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
