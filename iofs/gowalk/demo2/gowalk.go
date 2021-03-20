package demo

import (
	"io/fs"
	"strings"
)

func FindGoFiles(dir string, fsys fs.FS) ([]string, error) {
	var newEntries []string
	err := fs.WalkDir(fsys, dir, func(path string, entry fs.DirEntry, err error) error {
		if entry == nil {
			return nil
		}

		if !entry.IsDir() {
			ss := strings.Split(entry.Name(), ".")
			if ss[len(ss)-1] != "go" {
				return nil
			}
			newEntries = append(newEntries, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return newEntries, nil
}
