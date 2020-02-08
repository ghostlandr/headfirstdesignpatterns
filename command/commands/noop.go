package commands

func NewNoop() Command {
	return &noCmd{}
}

type noCmd struct {}

func (n *noCmd) Execute() {}
func (n *noCmd) Name() string { return "No command" }
func (n *noCmd) Undo() {}
