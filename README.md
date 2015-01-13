go-autoit
=========
A wrapper of AutoIt (AutoItX) for the Go Programming Language. [![GoDoc](https://godoc.org/github.com/brunoqc/go-autoit?status.png)](https://godoc.org/github.com/brunoqc/go-autoit)

##Dependencies##
- AutoIt (with AutoItX) from http://www.autoitscript.com/site/autoit/downloads/

##Sample code##
```go
package main

import "github.com/brunoqc/go-autoit"

func main() {
	success, pid := autoit.Run("notepad.exe", "", autoit.SW_NORMAL)
	if !success {
		log.Panic("can't run process")
	}

	log.Println("pid", pid)
}

```

##Build with##

You need the DLL (`AutoItX3.dll` or `AutoItX3_x64.dll` for 64-bit) in your `PATH`.

Note: I wasn't able to set `CGO_CFLAGS` to `C:\Program Files (x86)\AutoIt3\AutoItX\StandardDLL\VC6`. It doesn't seem to like white spaces in the path so I copied the `AutoIt3` directory to `c:\`

```bash
set CGO_CFLAGS=-Ic:/AutoIt3/AutoItX/StandardDLL/VC6
set CGO_LDFLAGS=-Lc:/AutoIt3/AutoItX -lAutoItX3

# (for 64-bit)
set CGO_LDFLAGS=-Lc:/AutoIt3/AutoItX -lAutoItX3_x64

go build
```
