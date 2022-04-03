package command

type ConvertToBase64Command struct {

}

func (cmd ConvertToBase64Command) ExecuteCommand(ctx ImageCommandContext) {

}

func (cmd ConvertToBase64Command) ValidateRequiredConfig(ctx ImageCommandContext) {

}

func (cmd ConvertToBase64Command) GetCommandName() string {
	return "base64"
}

func (cmd ConvertToBase64Command) GetCommandHelpName() []string {
	return []string{"b"}
}

func (cmd ConvertToBase64Command) GetUsage() string {
	return "encode to base64"
}