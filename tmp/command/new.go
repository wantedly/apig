package command

import (
	"strings"
)

type NewCommand struct {
	Meta
}

func (c *NewCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *NewCommand) Synopsis() string {
	return ""
}

func (c *NewCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
