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
//         success, pid := autoit.Run("notepad.exe", "", autoit.SW_NORMAL)
//         if !success {
//         	log.Panic("can't run process")
//         } else {
//         	log.Println("pid", pid)
//         }
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
	"encoding/binary"
	"syscall"
)

const (
	SW_HIDE     = C.SW_HIDE     // Hidden window
	SW_MINIMIZE = C.SW_MINIMIZE // Minimized window
	SW_MAXIMIZE = C.SW_MAXIMIZE // Maximized window
	SW_NORMAL   = 4
)

const (
	ENABLE_USER_INPUT  = 0
	DISABLE_USER_INPUT = 1
)

const (
	STATE_EXISTS    = 1
	STATE_VISIBLE   = 2
	STATE_ENABLED   = 4
	STATE_ACTIVE    = 8
	STATE_MINIMIZED = 16
	STATE_MAXIMIZED = 32
)

// Run a program and don't wait
// Possibles flags are SW_HIDE, SW_MINIMIZE, SW_MAXIMIZE and SW_NORMAL
// returns true on success with the pid
func Run(filename, workingdir string, flag int) (bool, int) {
	pid := C.AU3_Run((*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(filename)), (*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(workingdir)), C.long(flag))
	return C.AU3_error() == 0, int(pid)
}

// Block the keyboard and mouse
func BlockInput(flag int) {
	C.AU3_BlockInput(C.long(flag))
}

// Close a window
func WinClose(title, text string) {
	C.AU3_WinClose((*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(title)), (*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(text)))
}

func WinGetState(title, text string) (bool, int) {
	result := C.AU3_WinGetState((*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(title)), (*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(text)))
	return C.AU3_error() == 0, int(result)
}

func WinActive(title, text string) int {
	return int(C.AU3_WinActive((*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(title)), (*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(text))))
}

func WinExists(title, text string) int {
	return int(C.AU3_WinExists((*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(title)), (*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(text))))
}

func WinGetText(title, text string, bufSize int) (result string) {
	// TODO: test if bufSize is not greater than 64KB
	if bufSize < 1 {
		panic("bufSize must be greater than 0")
	}

	data := make([]uint16, bufSize)

	C.AU3_WinGetText((*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(title)), (*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(text)), (*_Ctype_WCHAR)(&data[0]), (C.int)(bufSize))

	for _, char := range data {
		if char == 0x0 {
			break
		}

		buf := make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, char)

		// FIXME: shoudln't have to only use the first byte
		result += string(buf[0])
	}

	return
}

func WinActivate(title, text string) {
	C.AU3_WinActivate((*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(title)), (*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(text)))
}

// Simulate input on the keyboard
// flag: 0: normal, 1: raw
func Send(keys string, flag int) {
	C.AU3_Send((*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(keys)), C.long(flag))
}

// Get the color of the pixel at the specified location
// return -1 if the location is invalid
func PixelGetColor(x, y int) int {
	return int(C.AU3_PixelGetColor(C.long(x), C.long(y)))
}

// Set/get a property
func Opt(option string, param int) int {
	return int(C.AU3_Opt((*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(option)), C.long(param)))
}

// click on a control without using the mouse pointer
// TODO: x, y should be center by defaut
func ControlClick(title, text, controlID, button string, clicks, x, y int) int {
	return int(C.AU3_ControlClick((*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(title)), (*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(text)), (*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(controlID)), (*_Ctype_WCHAR)(syscall.StringToUTF16Ptr(button)), C.long(clicks), C.long(x), C.long(y)))
}

// Get a checksum of the pixel in a region
func PixelChecksum(left, top, right, bottom, step int) int64 {
	return int64(C.AU3_PixelChecksum(C.long(left), C.long(top), C.long(right), C.long(bottom), C.long(step)))
}

// Move the mouse pointer to a specific location
func MouseMove(x, y, speed int) {
	C.AU3_MouseMove(C.long(x), C.long(y), C.long(speed))
}
