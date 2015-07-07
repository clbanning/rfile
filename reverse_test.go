// reverse_test.go - read a file line-by-line backwards

package rfile

import (
	"fmt"
	"io"
	"testing"
)

func TestRfile(t *testing.T) {
		rf, err := Open("test.file")
		if err != nil {
			t.Fatal(err)
		}
		for {
			line, err := rf.ReadLine()
			if err != nil {
				if err != io.EOF {
					t.Fatal(err)
				}
				break
			}
			fmt.Println(line)
		}
		rf.Close()
}
