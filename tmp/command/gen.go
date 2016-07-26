package command

import (
	"strings"
)

type GenCommand struct {
	Meta
}

func (c *GenCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *GenCommand) Synopsis() string {
	return ""
}

func (c *GenCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
