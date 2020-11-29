package command

type handler interface {
	execute()
}

type commandList struct {
	handlers map[string]handler
}

func (c *commandList) register(keyword string, cmd command) {
	c.handler[keyword] = cmd
}

type whocommand struct{}

func (c whocommand) execute() {

}
