package apig

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"text/template"

	"github.com/wantedly/apig/msg"
	"github.com/wantedly/apig/util"
)

var r = regexp.MustCompile(`_templates/skeleton/.*\.tmpl$`)

func generateSkeleton(detail *Detail, outDir string) error {
	var wg sync.WaitGroup
	errCh := make(chan error, 1)
	done := make(chan bool, 1)

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
				errCh <- err
			}

			tmpl, err := template.New("complex").Parse(string(body))
			if err != nil {
				errCh <- err
			}

			var buf bytes.Buffer
			var src []byte

			if err := tmpl.Execute(&buf, detail); err != nil {
				errCh <- err
			}

			if strings.HasSuffix(path, ".go") {
				src, err = format.Source(buf.Bytes())
				if err != nil {
					errCh <- err
				}
			} else {
				src = buf.Bytes()
			}

			if !util.FileExists(filepath.Dir(dstPath)) {
				if err := util.Mkdir(filepath.Dir(dstPath)); err != nil {
					errCh <- err
				}
			}

			if err := ioutil.WriteFile(dstPath, src, 0644); err != nil {
				errCh <- err
			}

			msg.Printf("\t\x1b[32m%s\x1b[0m %s\n", "create", dstPath)
		}(skeleton)
	}

	wg.Wait()
	close(done)

	select {
	case <-done:
	case err := <-errCh:
		if err != nil {
			return err
		}
	}

	return nil
}

func Skeleton(gopath, vcs, username, project, namespace, database string) int {
	detail := &Detail{
		VCS:       vcs,
		User:      username,
		Project:   project,
		Namespace: namespace,
		Database:  database,
	}
	outDir := filepath.Join(gopath, "src", detail.VCS, detail.User, detail.Project)
	if util.FileExists(outDir) {
		fmt.Fprintf(os.Stderr, "%s is already exists", outDir)
		return 1
	}
	if err := generateSkeleton(detail, outDir); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	msg.Printf("===> Created %s", outDir)
	return 0
}
