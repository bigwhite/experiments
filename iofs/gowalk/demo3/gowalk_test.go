package demo

import (
	"testing"
	"testing/fstest"
)

/*
$tree testdata
testdata
└── foo
    ├── 1
    │   └── 1.txt
    ├── 1.go
    ├── 2
    │   ├── 2.go
    │   └── 2.txt
    └── bar
        ├── 3
        │   └── 3.go
        └── 4.go

5 directories, 6 files

*/

func TestFindGoFiles(t *testing.T) {
	m := map[string]bool{
		"testdata/foo/1.go":       true,
		"testdata/foo/2/2.go":     true,
		"testdata/foo/bar/3/3.go": true,
		"testdata/foo/bar/4.go":   true,
	}

	mfs := fstest.MapFS{
		"testdata/foo/1.go":       {Data: []byte("package foo\n")},
		"testdata/foo/1/1.txt":    {Data: []byte("1111\n")},
		"testdata/foo/2/2.txt":    {Data: []byte("2222\n")},
		"testdata/foo/2/2.go":     {Data: []byte("package bar\n")},
		"testdata/foo/bar/3/3.go": {Data: []byte("package zoo\n")},
		"testdata/foo/bar/4.go":   {Data: []byte("package zoo1\n")},
	}

	files, err := FindGoFiles("testdata/foo", mfs)
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
