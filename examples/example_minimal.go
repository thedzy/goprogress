package main

import (
	"fmt"
	"github.com/thedzy/goprogress"
	"time"
)

func main() {
	// Set parameters for the demo
	total := 100
	goprogress.CreateProgress(goprogress.Options{
		Width:               25,
		Title:               "Loading",
		Footer:              "",
		BarCharacter:        '=',
		BarDividerCharacter: '=',
		FillCharacter:       ' ',
		IgnoreColour:        true,
		Terminators:         []string{"[", "]"},
	})

	// Simple
	fmt.Println("Simple ProgressBar")
	progress := 0
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the progress bar
		goprogress.DrawSimpleProgressBar(progress)
	}
	fmt.Println("\n")
}
