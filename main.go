package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/tcnksm/go-gitconfig"
)

//go:generate go-bindata _templates/...

func usage() {
	fmt.Fprintf(os.Stderr, `Usage of %s:
	%s new <project name>
	%s gen -d <model directory -o <output directory>`,
	os.Args[0], os.Args[0], os.Args[0])
	os.Exit(1)
}

func main() {

	if len(os.Args) < 2 {
		usage()
	}

	cmd := os.Args[1]

	switch cmd {
	case "gen":
		var (
			modelDir string
			outDir   string
		)

		flag := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		flag.Usage = func() {
			fmt.Fprintf(os.Stderr, `Usage of %s:
	   %s gen -d <model directory> -o <output directory>

	Options:
	`, os.Args[0], os.Args[0])
			flag.PrintDefaults()
		}

		flag.StringVar(&modelDir, "d", "", "Model directory")
		flag.StringVar(&outDir, "o", "", "Output directory")

		flag.Parse(os.Args[:2])

		if modelDir == "" || outDir == "" {
			flag.Usage()
			os.Exit(1)
		}
		cmdGen(modelDir, outDir)

	case "new":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, `Usage of %s:
	   %s gen new <model name>`, os.Args[0], os.Args[0])
			os.Exit(1)
		}

		name := os.Args[2]
		cmdNew(name)

	default:
		usage()
	}

}

func cmdNew(name string) {
	vcs := "github.com"
	username, err := gitconfig.Username()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		fmt.Println("Error: $GOPATH is not found")
		os.Exit(1)
	}

	projectDir := filepath.Join(gopath, "src", vcs, username, name)

	if err := copyStaticFiles(projectDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func cmdGen(modelDir, outDir string) {
	if !fileExists(outDir) {
		if err := mkdir(outDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	files, err := ioutil.ReadDir(modelDir)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var models []*Model

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if !strings.HasSuffix(file.Name(), ".go") {
			fmt.Fprintln(os.Stderr, file.Name()+" is not a Go file.")
			os.Exit(1)
		}

		modelPath := filepath.Join(modelDir, file.Name())
		ms, err := parseFile(modelPath)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		for _, model := range ms {
			models = append(models, model)
		}
	}

	if err := generateREADME(models, outDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := generateRouter(models, outDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, model := range models {
		if err := generateController(model, outDir); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
