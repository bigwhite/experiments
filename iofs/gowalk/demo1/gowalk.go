package demo

import (
	"os"
	"path/filepath"
	"strings"
)

func FindGoFiles(dir string) ([]string, error) {
	var goFiles []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		ss := strings.Split(path, ".")
		if ss[len(ss)-1] != "go" {
			return nil
		}

		goFiles = append(goFiles, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return goFiles, nil
}
