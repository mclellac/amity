package ui

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"
)

// ANSI escape colours
const (
	Grey       = "\x1B[38;5;243m"
	White      = "\x1B[38;5;011m"
	Red        = "\x1B[31;5;031m"
	Green      = "\x1B[31;5;032m"
	Blue       = "\x1B[31;5;034m"
	Cyan       = "\x1B[31;5;036m"
	LightRed   = "\x1B[31;5;091m"
	LightGreen = "\x1B[31;5;092m"
	LightBlue  = "\x1B[31;5;094m"
	LightCyan  = "\x1B[31;5;096m"
	Reset      = "\x1B[0m" // Reset escape sequence
)

// winsize struct
type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// DrawDivider draws a cheap horizontal reference line
// that is the width of the terminal.
func DrawDivider() {
	termWidth := GetTermWidth()
	divider := strings.Repeat("─", termWidth-1)
	fmt.Println(Grey, divider, Reset)
}

// GetTermWidth gets the char column size of the current terminal
// window.
func GetTermWidth() int {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}

	return int(ws.Col)
}
