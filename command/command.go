package command

import (
	"github.com/urfave/cli/v2"
)

type ImageCommandContext struct {
	Context       *cli.Context
	FromClipboard bool
	ImageName     string
	ImagePath     string
	ImageContent  []byte
}

type ImageCommand interface {
	ExecuteCommand(ctx ImageCommandContext) error
	ValidateRequiredConfig(ctx ImageCommandContext) error
	// provide info about command for cli
	GetCommandName() string
	GetCommandHelpName() []string
	GetUsage() string
	GetCategory() string
	GetArgUsages() string
	GetCommandFlags() []cli.Flag
}

type abstractImageCommand struct {
}

func (cmd abstractImageCommand) ExecuteCommand(ctx ImageCommandContext) error {
	return nil
}

func (cmd abstractImageCommand) ValidateRequiredConfig(ctx ImageCommandContext) error {
	return nil
}

func (cmd abstractImageCommand) GetCommandName() string {
	return "implementted"
}

func (cmd abstractImageCommand) GetCommandHelpName() []string {
	return []string{"implementted"}
}

func (cmd abstractImageCommand) GetUsage() string {
	return "implementted"
}

func (cmd abstractImageCommand) GetCategory() string {
	return ""
}

func (cmd abstractImageCommand) GetArgUsages() string {
	return ""
}

func (cmd abstractImageCommand) GetCommandFlags() []cli.Flag {
	return []cli.Flag{}
}
