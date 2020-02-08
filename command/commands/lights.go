package commands

import (
	"headfirstdesigntraining/command/objects"
)

func NewLightOn(l *objects.Light) Command {
	return &lightOnCommand{l: l}
}

type lightOnCommand struct {
	l *objects.Light
}

func (c *lightOnCommand) Execute() {
	c.l.On()
}

func (c *lightOnCommand) Undo() {
	c.l.Off()
}

func (c *lightOnCommand) Name() string {
	return "Turn light on"
}

func NewLightOff(l *objects.Light) Command {
	return &lightOffCommand{l: l}
}

type lightOffCommand struct {
	l *objects.Light
}

func (c *lightOffCommand) Execute() {
	c.l.Off()
}

func (c *lightOffCommand) Undo() {
	c.l.On()
}

func (c *lightOffCommand) Name() string {
	return "Turn light off"
}
