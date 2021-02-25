package cli

type CommandDescriptor struct {
	Name        string
	Handler     func(c *CommandContext) (err error)
	Summary     string
	Description string
}

type CommandCollection struct {
	commandList []CommandDescriptor
	commandMap  map[string]*CommandDescriptor
}

func NewCommandCollection(commandList []CommandDescriptor) *CommandCollection {
	cc := new(CommandCollection)
	cc.commandList = commandList
	cc.commandMap = make(map[string]*CommandDescriptor)
	for i, _ := range cc.commandList {
		command := &cc.commandList[i]
		cc.commandMap[command.Name] = command
	}
	return cc
}

func (cc *CommandCollection) Lookup(name string) (descriptor *CommandDescriptor, exists bool) {
	descriptor, exists = cc.commandMap[name]
	return
}
