package main

type CommandDescriptor struct {
	name        string
	handler     func(c *Context) (err error)
	summary     string
	description string
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
		cc.commandMap[command.name] = command
	}
	return cc
}
