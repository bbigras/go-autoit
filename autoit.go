// A wrapper of AutoIt (AutoItX) for the Go Programming Language

// A wrapper of AutoIt (AutoItX) for the Go Programming Language.
//
// Dependencies
//     AutoIt (with AutoItX) from http://www.autoitscript.com/site/autoit/downloads/
//
// Example
//     package main
//
//     import "github.com/brunoqc/go-autoit"
//
//     func main() {
//         autoit.Run("notepad.exe", "", autoit.SW_NORMAL)
//     }
//
// Build
//     set CGO_CFLAGS=-Ic:/AutoIt3/AutoItX/StandardDLL/VC6
//     set CGO_LDFLAGS=-lAutoItX3
//     set CGO_LDFLAGS=-lAutoItX3_x64 # for 64-bit
//     go build
package autoit

/*
#include <Windows.h>
#include <AutoIt3.h>
*/
import "C"

import (
	"syscall"
)

const (
	SW_HIDE     = C.SW_HIDE     // Hidden window
	SW_MINIMIZE = C.SW_MINIMIZE // Minimized window
	SW_MAXIMIZE = C.SW_MAXIMIZE // Maximized window
	SW_NORMAL   = 4
)

// Run a program and don't wait
// Possibles flags are SW_HIDE, SW_MINIMIZE, SW_MAXIMIZE and SW_NORMAL
func Run(filename, workingdir string, flag int) {
	C.AU3_Run((*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(filename)), (*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(workingdir)), C.long(flag))
}
