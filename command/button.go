package command

type Button struct {
	command Command
}

func NewButton(command Command) *Button {
	return new(Button{
		command: command,
	})
}

func (b *Button) Press() {
	b.command.execute()
}
