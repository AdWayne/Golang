package command

type DoorOpenCommand struct {
	door *Door
}

func (c *DoorOpenCommand) Execute() {
	c.door.Open()
}

func (c *DoorOpenCommand) Undo() {
	c.door.Close()
}

type DoorCloseCommand struct {
	door *Door
}

func (c *DoorCloseCommand) Execute() {
	c.door.Close()
}

func (c *DoorCloseCommand) Undo() {
	c.door.Open()
}