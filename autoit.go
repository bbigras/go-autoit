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
