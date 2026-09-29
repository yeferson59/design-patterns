package command

type OffCommand struct {
	device Device
}

func NewOffCommand(device Device) *OffCommand {
	return new(OffCommand{
		device: device,
	})
}

func (c *OffCommand) execute() {
	c.device.off()
}
