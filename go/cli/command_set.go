package cli

type CommandDescriptor struct {
	Name        string
	Handler     func(c *CommandContext) (err error)
	Summary     string
	Description string
}

type CommandSet struct {
	commandList []CommandDescriptor
	commandMap  map[string]*CommandDescriptor
}

func NewCommandSet(initialCommands []CommandDescriptor) *CommandSet {
	cs := new(CommandSet)
	cs.commandMap = make(map[string]*CommandDescriptor)
	for _, command := range initialCommands {
		cs.Add(command)
	}
	return cs
}

func (cs *CommandSet) Add(command CommandDescriptor) {
	cs.commandList = append(cs.commandList, command)
	cs.commandMap[command.Name] = &command
}

func (cs *CommandSet) Lookup(name string) (descriptor *CommandDescriptor, exists bool) {
	descriptor, exists = cs.commandMap[name]
	return
}
