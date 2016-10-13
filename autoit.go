// Package autoit is a AutoItX Go wrapper
//
// Dependencies
//     AutoItX (http://www.autoitscript.com/site/autoit/downloads/)
//
// Example
//     package main
//
//     import (
//     	"fmt"
//
//     	"github.com/brunoqc/go-autoit"
//     )
//
//     func main() {
//     	success, pid := autoit.Run("notepad.exe", "", autoit.SwNormal)
//     	if !success {
//     		panic("can't run process")
//     	}
//
//     	fmt.Println("pid", pid)
//     }
//
// Build
//     set CGO_CFLAGS=-Ic:/AutoItX
//     set CGO_LDFLAGS=-lAutoItX3_DLL
//     set CGO_LDFLAGS=-lAutoItX3_x64_DLL # for 64-bit
//     go build
package autoit

import (
	"runtime"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

const (
	EnableUserInput  = 0
	DisabelUserInput = 1
)

const (
	StateExists    = 1
	StateVisible   = 2
	StateEnabled   = 4
	StateActive    = 8
	StateMinimized = 16
	StateMaximized = 32
)

type utf16str string

func (p utf16str) Swigcptr() uintptr {
	c, err := syscall.UTF16FromString(string(p))
	if err != nil {
		panic(err)
	}

	return uintptr(unsafe.Pointer(&c))
}

// Run a program and don't wait
// Possibles flags are SwHide, SwMinimize, SwMaximize and SwNormal
// returns true on success with the pid
func Run(filename, workingdir string, flag int) (bool, int) {
	var pid int
	var err bool
	do(func() {
		pid = AU3_Run(utf16str(filename), utf16str(workingdir), flag)
		err = AU3_error() == 0
	})
	return err, pid
}

// WinClose closes a window
func WinClose(title, text string) {
	do(func() {
		AU3_WinClose(utf16str(title), utf16str(text))
	})
}

// WinGetState returns a window's state
func WinGetState(title, text string) (bool, int) {
	var result int
	var err bool
	do(func() {
		result = AU3_WinGetState(utf16str(title), utf16str(text))
		err = AU3_error() == 0
	})
	return err, result
}

// WinSetState sets a window's state
func WinSetState(title, text string, flag int) {
	do(func() {
		AU3_WinSetState(utf16str(title), utf16str(text), flag)
	})
}

// WinActive returns true if the window is active
func WinActive(title, text string) bool {
	var result bool
	do(func() {
		result = int(AU3_WinActive(utf16str(title), utf16str(text))) == 1
	})
	return result
}

// WinExists returns true if the window exist
func WinExists(title, text string) bool {
	var result bool
	do(func() {
		result = int(AU3_WinExists(utf16str(title), utf16str(text))) == 1
	})
	return result
}

func findTermChr(buff []uint16) int {
	for i, char := range buff {
		if char == 0x0 {
			return i
		}
	}
	panic("not supposed to happen")
}

// WinGetText returns the text contained in a window
func WinGetText(title, text string, bufSize int) (result string) {
	do(func() {
		// TODO: test if bufSize is not greater than 64KB
		if bufSize < 1 {
			panic("bufSize must be greater than 0")
		}

		buff := make([]uint16, bufSize)

		AU3_WinGetText(utf16str(title), utf16str(text), SwigcptrLPWSTR(unsafe.Pointer(&buff)), bufSize)
		pos := findTermChr(buff)
		result = string(utf16.Decode(buff[0:pos]))
	})

	return result
}

// WinActivate set the focus on a window
func WinActivate(title, text string) {
	do(func() {
		AU3_WinActivate(utf16str(title), utf16str(text))
	})
}

// Send simulates input on the keyboard
// flag: 0: normal, 1: raw
func Send(keys string, flag int) {
	do(func() {
		AU3_Send(utf16str(keys), flag)
	})
}

// PixelGetColor returns the color of the pixel at the specified location
// return -1 if the location is invalid
func PixelGetColor(x, y int) (result int) {
	do(func() {
		result = int(AU3_PixelGetColor(x, y))
	})
	return result
}

// Opt is used to set/get a property
func Opt(option string, param int) (result int) {
	do(func() {
		result = int(AU3_Opt(utf16str(option), param))
	})
	return result
}

// ControlClick clicks on a control without using the mouse pointer
// TODO: x, y should be center by defaut
func ControlClick(title, text, controlID, button string, clicks, x, y int) (result int) {
	do(func() {
		result = int(AU3_ControlClick(utf16str(title), utf16str(text), utf16str(controlID), utf16str(button), clicks, x, y))
	})
	return result
}

// https://msdn.microsoft.com/en-us/library/windows/desktop/dd162897(v=vs.85).aspx
type rect struct {
	Left, Top, Right, Bottom int32
}

// PixelChecksum returns a checksum of the pixel in a region
func PixelChecksum(left, top, right, bottom int32, step int) (result int64) {
	do(func() {
		r := &rect{left, top, right, bottom}
		result = int64(AU3_PixelChecksum(SwigcptrLPRECT(unsafe.Pointer(&r)), step))
	})
	return result
}

// MouseMove moves the mouse's pointer to a specific location
func MouseMove(x, y, speed int) {
	do(func() {
		AU3_MouseMove(x, y, speed)
	})
}

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

// Main runs the main AutoIt service loop.
// The binary's main.main must call autoit.Main() to run this loop.
// Main does not return. If the binary needs to do other work, it
// must do it in separate goroutines.
func Main() {
	for f := range mainfunc {
		f()
	}
}

// queue of work to run in main thread.
var mainfunc = make(chan func())

// do runs f on the main thread.
func do(f func()) {
	done := make(chan bool, 1)
	mainfunc <- func() {
		f()
		done <- true
	}
	<-done
}
