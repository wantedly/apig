package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestGenCommand_implement(t *testing.T) {
	var _ cli.Command = &GenCommand{}
}
