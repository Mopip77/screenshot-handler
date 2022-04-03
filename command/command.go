package command

import (
	"github.com/urfave/cli/v2"
)

type ImageCommandContext struct {
	Context      *cli.Context
	ImagePath    string
	ImageContent []byte
}

type ImageCommand interface {
	ExecuteCommand(ctx ImageCommandContext)
	ValidateRequiredConfig(ctx ImageCommandContext)
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
	// ioutil.ReadFile()

	// ctx.Command.HelpName
	return nil
}
