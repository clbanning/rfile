rfile
=====

read a file in reverse line-by-line

...
f, err := NewReverseFile(file string)
for {
  line, err := f.ReadLine()
  if err != nil {
      break // may be io.EOF
  }
  // do something with "line"
}
f.Close()
...
