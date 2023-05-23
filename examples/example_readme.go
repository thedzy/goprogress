package main

import (
	"fmt"
	"github.com/thedzy/goprogress"
	"time"
)

func main() {
	total := 25
	goprogress.CreateProgress(goprogress.Options{
		Total:               total,
		Width:               50,
		Title:               "Testing",
		Text:                "◢◤  ◢◤  ",
		Footer:              "",
		BarCharacter:        '=',
		BarDividerCharacter: '>',
		FillCharacter:       '-',
		BarColour:           []float32{1.0, 1.0, 0.0},
		FillColour:          []float32{0.0, 0.0, 0.5},
		LightTextColour:     []float32{0.9, 0.0, 0.0},
		DarkTextColour:      []float32{0.0, 0.0, 0.9},
		InvertColours:       true,
		IgnoreColour:        false,
		Partials:            []string{" ", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},
		Terminators:         []string{"▕", "▏"},
		Mode:                2,
	})

	// Wait bar
	fmt.Println("WaitBar")
	progress := 0
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the wait bar
		goprogress.DrawWaitBar()
	}
	fmt.Println("\n")

	// Simple
	fmt.Println("Simple ProgressBar")
	progress = 0
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the progress bar
		goprogress.DrawSimpleProgressBar(progress, goprogress.Options{
			BarColour:  []float32{1.00, 0.0, 0.0},
			FillColour: []float32{0.25, 0.0, 0.0},
		})
	}
	fmt.Println("\n")
}
