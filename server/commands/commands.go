package command

type handler interface {
	execute()
}

type commandList struct {
	handlers map[string]handler
}

func (c *commandList) register(keyword string, handler handler) {
	c.handlers[keyword] = handler
}

type whocommand struct{}

func (c whocommand) execute() {

}
