package command

type OnCommand struct {
	device Device
}

func NewOnCommand(device Device) *OnCommand {
	return new(OnCommand{
		device: device,
	})
}

func (c *OnCommand) execute() {
	c.device.on()
}
