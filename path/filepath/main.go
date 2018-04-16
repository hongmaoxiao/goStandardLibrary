package main

import (
	"fmt"
	"runtime"

	"github.com/hongmaoxiao/goStandardLibrary/path/filepath"
)

type PathTest struct {
	path, result string
}

var basetests = []PathTest{
	{"", "."},
	{".", "."},
	{"/.", "."},
	{"/", "/"},
	{"////", "/"},
	{"x/", "x"},
	{"abc", "abc"},
	{"abc/def", "def"},
	{"a/b/.x", ".x"},
	{"a/b/c.", "c."},
	{"a/b/c.x", "c.x"},
}

var winbasetests = []PathTest{
	{`c:\`, `\`},
	{`c:.`, `.`},
	{`c:\a\b`, `b`},
	{`c:a\b`, `b`},
	{`c:a\b\c`, `c`},
	{`\\host\share\`, `\`},
	{`\\host\share\a`, `a`},
	{`\\host\share\a\b`, `b`},
}

func main() {
	tests := basetests
	if runtime.GOOS == "windows" {
		// make unix tests work on windows
		for i := range tests {
			tests[i].result = filepath.Clean(tests[i].result)
		}
		// add windows specific tests
		tests = append(tests, winbasetests...)
	}
	for _, test := range tests {
		if s := filepath.Base(test.path); s != test.result {
			fmt.Printf("Base(%q) = %q, want %q", test.path, s, test.result)
		}
	}
}
