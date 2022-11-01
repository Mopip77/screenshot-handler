package command

import (
	"fmt"

	"github.com/Mopip77/screenshot-handler/config"
	"github.com/Mopip77/screenshot-handler/infra/output"
	"github.com/Mopip77/screenshot-handler/util"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

const (
	CHANNEL_SMMS   = "smms"
	CHANNEL_GITHUB = "github"
)

var (
	useJsDeliver bool
)

type UploadCommand struct {
	abstractImageCommand
}

type imageUploadResult struct {
	imageUrl     string
	imageChannel string
	err          error
}

func (cmd UploadCommand) ExecuteCommand(ctx ImageCommandContext) error {
	var imageUploadResults = []imageUploadResult{}

	for _, uploadChannel := range config.GlobalConfig.Upload.Use {
		var imageUrl string
		var channel string
		var err error
		switch uploadChannel {
		case CHANNEL_SMMS:
			imageUrl, err = util.UploadToSmms(ctx.ImageContent)
			channel = CHANNEL_SMMS
		case CHANNEL_GITHUB:
			imageUrl, err = util.UploadToGithub(ctx.ImageContent, useJsDeliver)
			channel = CHANNEL_GITHUB
		}
		imageUploadResults = append(imageUploadResults, imageUploadResult{
			imageUrl:     imageUrl,
			imageChannel: channel,
			err:          err,
		})
	}

	for idx, result := range imageUploadResults {
		if result.err == nil {
			markdownText := fmt.Sprintf("![%s](%s)", ctx.ImageName, result.imageUrl)

			output.Fmt.Printf("upload image to %s success, url: %s\n\n", color.CyanString(result.imageChannel), color.GreenString(result.imageUrl))
			// print templates
			output.Fmt.Printf("Markdown: %s\n", markdownText)
			output.Fmt.Printf("BBCode  : [IMG]%s[/IMG]\n", result.imageUrl)
			output.Fmt.Printf("HTML    : <img src=\"%s\" alt=\"%s\">\n", result.imageUrl, ctx.ImageName)

			if idx == len(imageUploadResults)-1 {
				// last reuslt
				util.WriteToClipboard(util.CLIPBOARD_FORMAT_TEXT, []byte(markdownText))
				output.CyanFmt.Println("\nmarkdown template saved to clipboard.")
			} else {
				output.Fmt.Printf("\n")
			}
		} else {
			output.RedFmt.Printf("upload image to %s failed, reason: %s \n\n", result.imageChannel, result.err.Error())
		}
	}

	return nil
}

func (cmd UploadCommand) ValidateRequiredConfig(ctx ImageCommandContext) error {
	var err error
	for _, use := range config.GlobalConfig.Upload.Use {
		switch use {
		case CHANNEL_SMMS:
			err = util.CheckRequiredUploadSmmsConfig()
		case CHANNEL_GITHUB:
			err = util.CheckRequiredUploadGithubConfig()
		default:
			err = fmt.Errorf(color.BlueString("upload.use ") + color.RedString(use) + " not supported, only support: [%s, %s]", CHANNEL_SMMS, CHANNEL_GITHUB)
		}
		if err != nil {
			return err
		}
	}
	return nil
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
			Name:        "js-delivr",
			Aliases:     []string{"jsd"},
			Value:       true,
			Usage:       "use js delivr as cdn (only for github)",
			Destination: &useJsDeliver,
		},
	}
}
