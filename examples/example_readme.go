package main

import (
	"fmt"
	"github.com/thedzy/goprogress"
	"time"
)

func main() {
	// Example from the readme
	total := 25
	progress := 0

	options := goprogress.Options{
		Total:          total,
		Width:          50,
		BarText:        "◢◤  ◢◤  ",
		BarColour:      []float32{1.0, 1.0, 0.0},
		FillColour:     []float32{0.0, 0.0, 0.5},
		BarTextColour:  []float32{0.9, 0.0, 0.0},
		FillTextColour: []float32{0.0, 0.0, 0.9},
		Title:          "Loading",
		Terminators:    []string{"[", "]"},
	}

	waitBar := goprogress.NewProgressBar(goprogress.StyleWait, options)

	fmt.Println("Wait Bar")
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the wait bar
		waitBar.Draw(progress)
	}
	fmt.Println("\n")

	simpleBar := goprogress.NewProgressBar(goprogress.StyleSimple, options)
	simpleBar.SetBarText("◤◢◣◥")
	simpleBar.SetBarColour(goprogress.Red())
	simpleBar.SetBarTextColour([]float32{0.25, 0.0, 0.0})
	simpleBar.SetFillColour([]float32{0.25, 0.0, 0.0})
	progress = 0

	fmt.Println("Detailed Bar")
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the wait bar
		simpleBar.Draw(progress)
	}
	fmt.Println("\n")
}
