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
	fileName := fmt.Sprintf("/tmp/sch-%s.png", time.Now().Format("2006-01-02_15-04-05"))
	if ctx.FromClipboard {
		ioutil.WriteFile(fileName, ctx.ImageContent, 0644)
	}
	
	exec.Command("open", fileName).Run()

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
