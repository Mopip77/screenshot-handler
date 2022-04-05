package command

import (
	"encoding/base64"

	"github.com/Mopip77/screenshot-handler/infra/output"
	"github.com/Mopip77/screenshot-handler/util"
)

type ConvertToBase64Command struct {
	abstractImageCommand
}

func (cmd ConvertToBase64Command) ExecuteCommand(ctx ImageCommandContext) error {
	base64Output := base64.StdEncoding.EncodeToString(ctx.ImageContent)
	output.Fmt.Println(base64Output)
	util.WriteToClipboard(util.CLIPBOARD_FORMAT_TEXT, []byte(base64Output))
	output.CyanFmt.Printf("\n[file path] %s\n", ctx.ImagePath)
	output.CyanFmt.Println("base64 encoded image has been copied to clipboard")
	return nil
}

func (cmd ConvertToBase64Command) GetCommandName() string {
	return "base64"
}

func (cmd ConvertToBase64Command) GetCommandHelpName() []string {
	return []string{"b"};
}

func (cmd ConvertToBase64Command) GetUsage() string {
	return "encode to base64"
}
