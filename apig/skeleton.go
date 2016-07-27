package apig

import (
	"bytes"
	"fmt"
	"text/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/wantedly/apig/util"
)

func generateSkeleton(detail *Detail, outDir string) error {
	ch := make(chan error)
	go func() {
		var wg sync.WaitGroup
		r := regexp.MustCompile(`_templates/skeleton/*`)
		for _, skeleton := range AssetNames() {
			wg.Add(1)
			go func(s string) {
				defer wg.Done()
				if !r.MatchString(s) {
					return
				}

				trim := strings.Replace(s, "_templates/skeleton/", "", 1)
				path := strings.Replace(trim, ".tmpl", "", 1)
				dstPath := filepath.Join(outDir, path)

				body, err := Asset(s)

				if err != nil {
					ch <- err
				}

				tmpl, err := template.New("complex").Parse(string(body))

				if err != nil {
					ch <- err
				}

				var buf bytes.Buffer

				if err := tmpl.Execute(&buf, detail); err != nil {
					ch <- err
				}

				if !util.FileExists(filepath.Dir(dstPath)) {
					if err := util.Mkdir(filepath.Dir(dstPath)); err != nil {
						ch <- err
					}
				}

				if err := ioutil.WriteFile(dstPath, buf.Bytes(), 0644); err != nil {
					ch <- err
				}

				fmt.Printf("\t\x1b[32m%s\x1b[0m %s\n", "create", dstPath)
			}(skeleton)
		}
		wg.Wait()
		ch <- nil
	}()

	err := <-ch
	if err != nil {
		return err
	}

	return nil
}

func Skeleton(gopath, vcs, username, project string) int {
	detail := &Detail{VCS: vcs, User: username, Project: project}
	outDir := filepath.Join(gopath, "src", detail.VCS, detail.User, detail.Project)
	if util.FileExists(outDir) {
		fmt.Fprintf(os.Stderr, "%s is already exists", outDir)
		return 1
	}
	if err := generateSkeleton(detail, outDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Fprintf(os.Stdout, `===> Created %s
`, outDir)
	return 0
}
