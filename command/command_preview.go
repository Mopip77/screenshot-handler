package command

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

type PriviewCommand struct {
	abstractImageCommand
}

func (cmd PriviewCommand) ExecuteCommand(ctx ImageCommandContext) error {
	previewFilePath := ctx.ImagePath
	if ctx.FromClipboard {
		tmpPreviewFilePath := fmt.Sprintf("/tmp/sch-%s.png", time.Now().Format("2006-01-02_15-04-05"))
		if err := ioutil.WriteFile(tmpPreviewFilePath, ctx.ImageContent, 0644); err != nil {
			return err
		}
		previewFilePath = tmpPreviewFilePath
	}
	
	exec.Command("open", previewFilePath).Run()

	return nil
}

func (cmd PriviewCommand) GetCommandName() string {
	return "preview"
}

func (cmd PriviewCommand) GetCommandHelpName() []string {
	return []string{"p"}
}

func (cmd PriviewCommand) GetUsage() string {
	return "preview screenshot"
}
