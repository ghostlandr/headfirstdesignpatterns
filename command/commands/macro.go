package commands

import "fmt"

type macroCommand struct {
	cmds []Command
	name string
}

func NewMacro(name string, cmds []Command) Command {
	return &macroCommand{cmds: cmds, name: name}
}

func (c *macroCommand) Execute() {
	for _, cmd := range c.cmds {
		cmd.Execute()
	}
}

func (c *macroCommand) Undo() {
	for i := len(c.cmds)- 1; i >= 0; i-- {
		c.cmds[i].Undo()
	}
}

func (c *macroCommand) Name() string {
	return fmt.Sprintf("Macro: %s", c.name)
}
