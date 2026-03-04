package command

type TempUpCommand struct {
	thermo *Thermostat
}

func (c *TempUpCommand) Execute() {
	c.thermo.Increase()
}

func (c *TempUpCommand) Undo() {
	c.thermo.Decrease()
}

type TempDownCommand struct {
	thermo *Thermostat
}

func (c *TempDownCommand) Execute() {
	c.thermo.Decrease()
}

func (c *TempDownCommand) Undo() {
	c.thermo.Increase()
}