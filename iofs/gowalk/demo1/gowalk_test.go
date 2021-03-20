package demo

import (
	"testing"
)

func TestFindGoFiles(t *testing.T) {
	m := map[string]bool{
		"testdata/foo/1.go":       true,
		"testdata/foo/2/2.go":     true,
		"testdata/foo/bar/3/3.go": true,
		"testdata/foo/bar/4.go":   true,
	}

	files, err := FindGoFiles("testdata/foo")
	if err != nil {
		t.Errorf("want nil, actual %s", err)
	}

	if len(files) != 4 {
		t.Errorf("want 4, actual %d", len(files))
	}

	for _, f := range files {
		_, ok := m[f]
		if !ok {
			t.Errorf("want [%s], actual not found", f)
		}
	}
}
