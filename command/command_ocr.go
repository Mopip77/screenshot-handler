package command

import (
	"fmt"

	"github.com/Mopip77/screenshot-handler/config"
	"github.com/Mopip77/screenshot-handler/infra/output"
	"github.com/Mopip77/screenshot-handler/util"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

type OcrCommand struct {
	abstractImageCommand
}

var (
	useLineFeed  bool
	useFullWidth bool
)

func (cmd OcrCommand) ExecuteCommand(ctx ImageCommandContext) error {
	var ocrResult string
	var err error
	switch config.GlobalConfig.Ocr.Use {
	case "tencent":
		ocrResult, err = util.OcrTencent(ctx.ImageContent, useLineFeed, useFullWidth)
	}

	if err != nil {
		return err
	}

	util.WriteToClipboard(util.CLIPBOARD_FORMAT_TEXT, []byte(ocrResult))

	output.GreenFmt.Println("ocr result (saved to clipboard):")
	output.Fmt.Println(ocrResult)

	return nil
}

func (cmd OcrCommand) ValidateRequiredConfig(ctx ImageCommandContext) error {
	switch config.GlobalConfig.Ocr.Use {
	case "tencent":
		return util.CheckRequiredOcrTencentConfig()
	case "":
		return fmt.Errorf(color.BlueString("ocr.use") + " is not set, please set it in config file")
	default:
		return fmt.Errorf(color.BlueString("ocr.use ") + color.RedString(config.GlobalConfig.Ocr.Use) + " not supported, only support: [tencent]")
	}
}

func (cmd OcrCommand) GetCommandName() string {
	return "ocr"
}

func (cmd OcrCommand) GetCommandHelpName() []string {
	return []string{"o"}
}

func (cmd OcrCommand) GetUsage() string {
	return "screenshot ocr with [options]"
}

func (cmd OcrCommand) GetCommandFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "linefeed",
			Aliases:     []string{"lf"},
			Value:       true,
			Usage:       "output linefeed",
			Destination: &useLineFeed,
		},
		&cli.BoolFlag{
			Name:        "fullwidth",
			Value:       false,
			Usage:       "transform halfwidth to fullwidth",
			Destination: &useFullWidth,
		},
	}
}
