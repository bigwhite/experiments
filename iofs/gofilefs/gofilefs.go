package gfs

import (
	"io/fs"
	"os"
	"strings"
)

func GoFilesFS(dir string) fs.FS {
	return goFilesFS(dir)
}

type goFile struct {
	*os.File
}

func Open(name string) (*goFile, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	return &goFile{f}, nil
}

func (f goFile) ReadDir(count int) ([]fs.DirEntry, error) {
	entries, err := f.File.ReadDir(count)
	if err != nil {
		return nil, err
	}
	var newEntries []fs.DirEntry

	for _, entry := range entries {
		if !entry.IsDir() {
			ss := strings.Split(entry.Name(), ".")
			if ss[len(ss)-1] != "go" {
				continue
			}
		}
		newEntries = append(newEntries, entry)
	}
	return newEntries, nil
}

type goFilesFS string

func (dir goFilesFS) Open(name string) (fs.File, error) {
	f, err := Open(string(dir) + "/" + name)
	if err != nil {
		return nil, err // nil fs.File
	}
	return f, nil
}
