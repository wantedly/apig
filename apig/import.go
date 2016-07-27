package apig

import "path/filepath"

func formatImportDir(paths []string) []string {
	results := make([]string, 0, len(paths))
	flag := map[string]bool{}
	for i := 0; i < len(paths); i++ {
		dir := filepath.Dir(paths[i])
		if !flag[dir] {
			flag[dir] = true
			results = append(results, dir)
		}
	}
	return results
}
