//go:build !windows && !plan9 && !solaris
// +build !windows,!plan9,!solaris

package goprogress

import (
	// "golang.org/x/sys/windows"
	"golang.org/x/sys/unix"
	"os"
)

// getWinSize Get the full size of the window
func getWinSize() (*unix.Winsize, error) {
	// Get window dimensions for Unix
	winDimensions, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, err
	}
	return winDimensions, nil
}
