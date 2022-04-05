package command

import (
	"fmt"

	"github.com/Mopip77/screenshot-handler/config"
	"github.com/Mopip77/screenshot-handler/infra/output"
	"github.com/Mopip77/screenshot-handler/util"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

var (
	useJsDeliver bool
)

type UploadCommand struct {
	abstractImageCommand
}

func (cmd UploadCommand) ExecuteCommand(ctx ImageCommandContext) error {
	var imageUrl string
	var err error
	switch config.GlobalConfig.Upload.Use {
	case "smms":
		imageUrl, err = util.UploadToSmms(ctx.ImageContent)
	case "github":
		imageUrl, err = util.UploadToGithub(ctx.ImageContent, useJsDeliver)
	}

	if err != nil {
		return err
	}

	markdownText := fmt.Sprintf("![%s](%s)", ctx.ImageName, imageUrl)
	util.WriteToClipboard(util.CLIPBOARD_FORMAT_TEXT, []byte(markdownText))

	output.Fmt.Printf("upload image to %s success, url: %s\n\n", color.CyanString(config.GlobalConfig.Upload.Use), color.GreenString(imageUrl))
	// print templates
	output.Fmt.Printf("Markdown: %s\n", markdownText)
	output.Fmt.Printf("BBCode  : [IMG]%s[/IMG]\n", imageUrl)
	output.Fmt.Printf("HTML    : <img src=\"%s\" alt=\"%s\">\n", imageUrl, ctx.ImageName)

	output.CyanFmt.Println("\nmarkdown template saved to clipboard.")

	return nil
}

func (cmd UploadCommand) ValidateRequiredConfig(ctx ImageCommandContext) error {
	switch config.GlobalConfig.Upload.Use {
	case "smms":
		return util.CheckRequiredUploadSmmsConfig()
	case "github":
		return util.CheckRequiredUploadGithubConfig()
	case "":
		return fmt.Errorf(color.BlueString("upload.use") + " is not set, please set it in config file")
	default:
		return fmt.Errorf(color.BlueString("upload.use ") + color.RedString(config.GlobalConfig.Upload.Use) + " not supported, only support: [smms, github]")
	}
}

func (cmd UploadCommand) GetCommandName() string {
	return "upload"
}

func (cmd UploadCommand) GetCommandHelpName() []string {
	return []string{"u"}
}

func (cmd UploadCommand) GetUsage() string {
	return "screenshot upload with [options]"
}

func (cmd UploadCommand) GetCommandFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "js-deliver",
			Aliases:     []string{"jsd"},
			Value:       true,
			Usage:       "use js deliver as cdn (only for github)",
			Destination: &useJsDeliver,
		},
	}
}
