package main

import (
	"fmt"
	"github.com/thedzy/goprogress/v2"
	"time"
)

func main() {
	// Set parameters for the demo
	total := 100
	progress := 0

	bar := goprogress.NewProgressBar(goprogress.StyleSimple, goprogress.Options{
		Total:          total,
		BarText:        "━",
		FillText:       "─",
		Title:          "Loading",
		BarTextColour:  []float32{1.0, 0.0, 0.0},
		FillTextColour: []float32{0.0, 0.0, 1.0},
		Terminators:    []string{"◯", "◯"},
	})

	fmt.Println("Detailed Bar")
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")
}
