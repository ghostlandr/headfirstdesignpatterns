package remote

import (
	"fmt"
	"strings"
	"headfirstdesigntraining/command/commands"
)

const remoteSlots = 7

type Remote interface {
	SetCommand(slot int, on, off commands.Command)
	OnButtonWasPressed(slot int)
	OffButtonWasPressed(slot int)
	UndoButtonWasPressed()
}

type remote struct {
	onCommands []commands.Command
	offCommands []commands.Command
	undoCommand commands.Command
}

func New() Remote {
	r := &remote{
		onCommands: make([]commands.Command, remoteSlots),
		offCommands: make([]commands.Command, remoteSlots),
	}
	noc := commands.NewNoop()
	for i := 0; i < remoteSlots; i++ {
		r.onCommands[i] = noc
		r.offCommands[i] = noc
	}
	r.undoCommand = noc
	return r
}

func (r *remote) SetCommand(slot int, on, off commands.Command) {
	r.onCommands[slot] = on
	r.offCommands[slot] = off
}

func (r *remote) OnButtonWasPressed(slot int) {
	r.onCommands[slot].Execute()
	r.undoCommand = r.onCommands[slot]
}

func (r *remote) OffButtonWasPressed(slot int) {
	r.offCommands[slot].Execute()
	r.undoCommand = r.offCommands[slot]
}

func (r *remote) UndoButtonWasPressed() {
	r.undoCommand.Undo()
}

func (r *remote) String() string {
	var b strings.Builder
	fmt.Fprint(&b, "*----- Remote Control -----*\n")
	for i := range r.onCommands {
		fmt.Fprintf(&b, "[slot %d] %s - %s\n", i, r.onCommands[i].Name(), r.offCommands[i].Name())
	}
	fmt.Fprintf(&b, "Last button pressed: %s\n", r.undoCommand.Name())
	fmt.Fprint(&b, "*----- /Remote Control -----*\n")
	return b.String()
}

