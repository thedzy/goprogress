//go:build windows
// +build windows

package goprogress

import (
	"golang.org/x/sys/windows"
	"os"
)

// Credit to https://github.com/buger/goterm/blob/master/terminal_sysioctl.go for solving my windows/unix build issues

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getWinSize() (*winsize, error) {
	ws := new(winsize)
	fd := os.Stdout.Fd()
	var info windows.ConsoleScreenBufferInfo
	if err := windows.GetConsoleScreenBufferInfo(windows.Handle(fd), &info); err != nil {
		return nil, err
	}

	ws.Col = uint16(info.Window.Right - info.Window.Left + 1)
	ws.Row = uint16(info.Window.Bottom - info.Window.Top + 1)

	return ws, nil
}
