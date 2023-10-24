// reverse_test.go - read a file line-by-line backwards

package rfile

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func TestRfile(t *testing.T) {
	tests := map[string]struct {
		content   string
		wantLines []string
	}{
		"empty file": {
			content:   "",
			wantLines: nil,
		},
		"one-line file": {
			content:   "1st line",
			wantLines: []string{"1st line"},
		},
		"multi-line file": {
			content:   "1st line\n2nd line\n",
			wantLines: []string{"2nd line", "1st line"},
		},
		"multi-line file with empty lines": {
			content:   "1st line\n2nd line\n3rd line\n\n4th line\n\n",
			wantLines: []string{"", "4th line", "", "3rd line", "2nd line", "1st line"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			filename := createTempFile(t, test.content)
			defer os.Remove(filename)

			rf, err := Open(filename)
			if err != nil {
				t.Fatal(err)
			}
			defer rf.Close()

			var lines []string

			for {
				line, err := rf.ReadLine()
				if err != nil {
					if err != io.EOF {
						t.Fatal(err)
					}
					break
				}
				lines = append(lines, line)
			}

			if !reflect.DeepEqual(test.wantLines, lines) {
				t.Errorf("not equal: want '%v' got '%v'", test.wantLines, lines)
			}
		})
	}
}

func createTempFile(t *testing.T, content string) string {
	t.Helper()
	file, err := os.CreateTemp("", "rfile-test-ReadLine")
	if err != nil {
		t.Fatalf("create a temp file: %v", err)
	}
	defer func() { _ = file.Close() }()

	_, _ = file.WriteString(content)
	return file.Name()
}
