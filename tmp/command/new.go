package command

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tcnksm/go-gitconfig"
	"github.com/wantedly/apig/apig"
)

const defaultVCS = "github.com"

type NewCommand struct {
	Meta

	project  string
	vcs      string
	username string
}

func (c *NewCommand) Run(args []string) int {
	if err := c.parseArgs(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	if c.vcs == "" {
		c.vcs = defaultVCS
	}
	if c.username == "" {
		var err error
		c.username, err = gitconfig.GithubUser()
		if err != nil {
			c.username, err = gitconfig.Username()
			if err != nil || strings.Contains(c.username, " ") {
				msg := "Cannot find github username in `~/.gitconfig` file.\n" +
					"Please use -u option"
				fmt.Fprintln(os.Stderr, msg)
				return 1
			}
		}
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		fmt.Fprintln(os.Stderr, "Error: $GOPATH is not found")
		return 1
	}

	return apig.Skeleton(gopath, c.vcs, c.username, c.project)
}

func (c *NewCommand) parseArgs(args []string) error {
	flag := flag.NewFlagSet("apig", flag.ContinueOnError)

	flag.StringVar(&c.vcs, "v", "", "VCS")
	flag.StringVar(&c.vcs, "vcs", "", "VCS")
	flag.StringVar(&c.username, "u", "", "Username")
	flag.StringVar(&c.username, "user", "", "Username")

	if err := flag.Parse(args); err != nil {
		return err
	}
	for 0 < flag.NArg() {
		c.project = flag.Arg(0)
		flag.Parse(flag.Args()[1:])
	}
	if c.project == "" {
		err := errors.New("Please specify project name.")
		return err
	}
	return nil
}

func (c *NewCommand) Synopsis() string {
	return ""
}

func (c *NewCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
