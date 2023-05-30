package main

import (
	"fmt"
	"github.com/thedzy/goprogress/v2"
	"time"
)

func main() {
	// Request greeting messages for the names.
	total := 100
	progress := 0
	duration := 100

	bar := goprogress.NewProgressBar(goprogress.StyleSmooth, goprogress.Options{
		Total:      total,
		Width:      total / 8,
		BarColour:  goprogress.Yellow(),
		FillColour: goprogress.DkBlue(),
		Mode:       goprogress.ModeNone,
	})

	// Smooth with fractional widths
	fmt.Println("Smooth ProgressBar")
	progress = 0
	for progress < total {
		progress++
		time.Sleep(time.Duration(duration) * time.Millisecond)

		// Update the progress bar
		bar.Draw(progress)
	}
	time.Sleep(2)
	fmt.Println("\n")

	// Simple with no text
	fmt.Println("Simple ProgressBar")
	bar.SetStyle(goprogress.StyleSimple)
	progress = 0
	for progress < total {
		progress++
		time.Sleep(time.Duration(duration) * time.Millisecond)

		// Update the progress bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

}
