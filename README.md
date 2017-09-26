rfile
=====

Read a file in reverse line-by-line.

Inspired by: a [gonuts discussion][gonuts].

Documentation: in [godoc][godoc].

<pre>
...
  f, err := rfile.Open(file)
  if err != nil {
    // handle err
  }
  for {
    line, err := f.ReadLine()
    if err != nil {
      if err != io.EOF {
        // handle error
      }
      break // must be io.EOF
    }
    // do something with "line"
  }
  f.Close()
  ...
</pre>

[gonuts]: https://groups.google.com/forum/?fromgroups#!topic/golang-nuts/a0NRB-spRhc
[godoc]: https://godoc.org/github.com/clbanning/rfile

