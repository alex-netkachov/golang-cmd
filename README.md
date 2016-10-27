# golang-cmd

Go library for parsing the command line and seamless execution of
shell commands. Works pretty much like a shell script or makefile:
prints output to stdout, stops when the command fails.

The following example demonstrates how to use the library:

```
import "github.com/AlexAtNet/golang-cmd"

src := cmd.GetLines("find . -name *.go")
version := cmd.Get("go version")
cmd.Run("go build")
```


