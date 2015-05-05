go-autoit
=========
[![GoDoc](https://godoc.org/github.com/brunoqc/go-autoit?status.png)](https://godoc.org/github.com/brunoqc/go-autoit)

A Go [AutoItX](https://www.autoitscript.com/site/autoit/) wrapper.

##Sample code##
```go
package main

import "github.com/brunoqc/go-autoit"

func main() {
	success, pid := autoit.Run("notepad.exe", "", autoit.SwNormal)
	if !success {
		log.Panic("can't run process")
	}

	log.Println("pid", pid)
}

```

##Build##

It seems the latest `AutoItX3_DLL.h` file doesn't build with Go since it uses defaults arguments with some functions (a C++ feature). C++ with Go is supported by [SWIG](swig.org/Doc3.0/Go.html) but I wasn't able to get it to work and I'll try again with Go 1.5. For now you can patch `AutoItX3_DLL.h` (see [AutoItX3_DLL.h.diff](AutoItX3_DLL.h.diff)).

You need the DLL (`AutoItX3.dll` or `AutoItX3_x64.dll` for 64-bit) in your `PATH`.

Note: I wasn't able to set `CGO_CFLAGS` to `C:\Program Files (x86)\AutoIt3\AutoItX`. It doesn't seem to like white spaces in the path so I copied the `AutoItX` directory to `c:\`

```bash
set CGO_CFLAGS=-Ic:/AutoItX
set CGO_LDFLAGS=-Lc:/AutoItX -lAutoItX3_DLL

# (for 64-bit)
set CGO_LDFLAGS=-Lc:/AutoItX -lAutoItX3_x64_DLL

go build
```
