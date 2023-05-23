package main

import (
	"fmt"
	"github.com/thedzy/goprogress"
	"time"
)

func main() {
	// Set parameters for the demo
	total := 100
	progress := 0
	fmt.Println("WaitBars")
	goprogress.CreateProgress(goprogress.Options{
		Total:      total,
		Width:      20,
		Title:      "Waiting",
		FillColour: []float32{1.0, 1.0, 0.0},
		BarColour:  []float32{0.0, 0.05, 0.6},
	})

	// Wait bar
	progress = 0
	for progress < total {
		progress++
		time.Sleep(time.Duration(50) * time.Millisecond)

		// Update the progress bar
		goprogress.DrawWaitBar(goprogress.Options{
			Text:       "▁▂▂▃▄▅▆▇▇██▇▇▆▅▄▃▂▂▁",
			FillColour: []float32{1.0, 1.0, 0.0},
			BarColour:  []float32{0.0, 0.0, 1.0},
		})
	}
	fmt.Println("\n")

	progress = 0
	for progress < total/5 {
		progress++
		time.Sleep(time.Duration(250) * time.Millisecond)

		// Update the progress bar
		goprogress.DrawWaitBar(goprogress.Options{
			Text:       " ◢◤",
			FillColour: []float32{1.0, 1.0, 0.0},
			BarColour:  []float32{0.0, 0.0, 0.0},
		})
	}
	fmt.Println("\n")

	progress = 0
	for progress < total/5 {
		progress++
		time.Sleep(time.Duration(250) * time.Millisecond)

		// Update the progress bar
		goprogress.DrawWaitBar(goprogress.Options{
			Text:         "▷▶ ▸ ",
			IgnoreColour: true,
		})
	}
	fmt.Println("\n")

	progress = 0
	for progress < total/5 {
		progress++
		time.Sleep(time.Duration(250) * time.Millisecond)

		// Update the progress bar
		goprogress.DrawWaitBar(goprogress.Options{
			Text:       "◠◡",
			FillColour: []float32{0.0, 0.0, 0.0},
			BarColour:  []float32{0.0, 1.0, 0.0},
		})
	}
	fmt.Println("\n")

	progress = 0
	for progress < total/5 {
		progress++
		time.Sleep(time.Duration(250) * time.Millisecond)

		// Update the progress bar
		goprogress.DrawWaitBar(goprogress.Options{
			Text:       "◴◵◶◷",
			FillColour: []float32{0.0, 0.0, 0.0},
			BarColour:  []float32{0.0, 1.0, 0.0},
		})
	}
	fmt.Println("\n")
}
