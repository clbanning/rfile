// reverse.go - read a file line-by-line backwards

package rfile

import (
	"bytes"
	"io"
	"os"
)

type Rfile struct {
	fh             *os.File
	offset         int64
	bufsize        int
	lines          [][]byte
	i              int
	atStartOfFile bool
}

// Open a file that you want read in reverse line-by-line.
func NewReverseFile(file string) (*Rfile, error) {
	var err error
	rf := new(Rfile)
	if rf.fh, err = os.Open(file); err != nil {
		return nil, err
	}
	fi, _ := rf.fh.Stat()
	rf.offset = fi.Size()
	rf.lines = make([][]byte,0)
	rf.bufsize = 4096
	return rf, nil
}

// Close file that was opened.
func (rf *Rfile) Close() {
	rf.fh.Close()
}

// Read next previous line, beginning with the last line in the file.
// When the beginning of the file is reached: "", io.EOF is returned.
func (rf *Rfile) ReadLine() (string, error) {
	if rf.i > 0 {
		rf.i--
		return string(rf.lines[rf.i+1]), nil
	}
	if rf.i < 0 {
		return "", io.EOF
	}
	if rf.atStartOfFile {
		rf.i--
		return string(rf.lines[0]), nil
	}

	// save what hasn't been retrieved
	var remainder []byte
	if len(rf.lines) > 0 {
		remainder = rf.lines[0]
	} else {
		remainder = make([]byte,0)
	}

	// get more from file
	// back up from end-of-file
	rf.offset -= int64(rf.bufsize)
	if rf.offset < 0 {
		rf.bufsize += int(rf.offset) // shrink the buffer
		rf.offset = 0
		rf.atStartOfFile = true
	}
	_, err := rf.fh.Seek(rf.offset, 0) // read end of file
	if err != nil {
		return "", err
	}

	// compute buffer size
	buf := make([]byte, rf.bufsize)
	if n, err := rf.fh.Read(buf); err != nil && err != io.EOF {
		return "", err
	} else if n != rf.bufsize { // shouldn't happen
		buf = buf[:n]
	}

	// get the lines in the buffer, append what was carried over
	buf = append(buf, remainder...)
	rf.lines = bytes.Split(buf, []byte("\n"))
	rf.i = len(rf.lines)-1

	return rf.ReadLine() // now read the next line back
}
