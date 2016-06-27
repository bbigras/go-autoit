go-autoit
=========
[![GoDoc](https://godoc.org/github.com/brunoqc/go-autoit?status.png)](https://godoc.org/github.com/brunoqc/go-autoit)
[![Build status](https://ci.appveyor.com/api/projects/status/nhb09oh0gei24md9?svg=true)](https://ci.appveyor.com/project/brunoqc/go-autoit)

A Go [AutoItX](https://www.autoitscript.com/site/autoit/) wrapper.

##Sample code##
```go
package main

import (
	"fmt"

	"github.com/brunoqc/go-autoit"
)

func main() {
	success, pid := autoit.Run("notepad.exe", "", autoit.SwNormal)
	if !success {
		panic("can't run process")
	}

	fmt.Println("pid", pid)
}
```

##Build##

You need the DLL (`AutoItX3.dll` or `AutoItX3_x64.dll` for 64-bit) and the [Swig](http://www.swig.org/) executable in your `PATH`.

Note: I wasn't able to set `CGO_CFLAGS` to `C:\Program Files (x86)\AutoIt3\AutoItX`. It doesn't seem to like white spaces in the path so I copied the `AutoItX` directory to `c:\`

Note 2: swig seems to ignore `CGO_CFLAGS` so you may have to copy `AutoItX3_DLL.h` to the current directory.

```bash
set CGO_CFLAGS=-Ic:/AutoItX
set CGO_LDFLAGS=-Lc:/AutoItX -lAutoItX3_DLL

# (for 64-bit)
set CGO_LDFLAGS=-Lc:/AutoItX -lAutoItX3_x64_DLL

go build
```
