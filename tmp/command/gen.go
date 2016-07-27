package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wantedly/apig/apig"
	"github.com/wantedly/apig/util"
)

const (
	modelDir   = "models"
	targetFile = "main.go"
)

type GenCommand struct {
	Meta
}

func (c *GenCommand) Run(args []string) int {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	if !util.FileExists(filepath.Join(wd, targetFile)) || !util.FileExists(filepath.Join(wd, modelDir)) {
		fmt.Fprintf(os.Stderr, `%s is not project root. Please move.
`, wd)
		return 1
	}
	return apig.Generate(wd, modelDir, targetFile)
}

func (c *GenCommand) Synopsis() string {
	return ""
}

func (c *GenCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
