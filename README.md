rfile
=====

read a file in reverse line-by-line
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
