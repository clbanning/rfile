rfile
=====

Read a file in reverse line-by-line.

Inspired by: https://groups.google.com/forum/?fromgroups#!topic/golang-nuts/a0NRB-spRhc.

<pre>
...
  f, err := NewReverseFile(file string)
  if err != nil {
    // handle err
  }
  for {
    line, err := f.ReadLine()
    if err != nil {
      if err != io.EOF {
        // handle error
      }
      break // may be io.EOF
    }
    // do something with "line"
  }
  f.Close()
  ...
</pre>
