package main

import (
	"fmt"
	"github.com/thedzy/goprogress"
	"time"
)

func main() {
	// Request greeting messages for the names.
	total := 120
	progress := 0

	bar := goprogress.NewProgressBar(goprogress.StyleSmooth, goprogress.Options{
		Total:   total,
		Width:   50,
		BarText: " ",
		Mode:    goprogress.ModeNone,
	})

	lastColour := goprogress.White()
	colours := [][]float32{
		goprogress.LtGrey(),
		goprogress.Grey(),
		goprogress.DkGrey(),

		goprogress.LtRed(),
		goprogress.Red(),
		goprogress.DkRed(),

		goprogress.LtGreen(),
		goprogress.Green(),
		goprogress.DkGreen(),

		goprogress.LtBlue(),
		goprogress.Blue(),
		goprogress.DkBlue(),

		goprogress.LtCyan(),
		goprogress.Cyan(),
		goprogress.DkCyan(),

		goprogress.LtMagenta(),
		goprogress.Magenta(),
		goprogress.DkMagenta(),

		goprogress.LtYellow(),
		goprogress.Yellow(),
		goprogress.DkYellow(),
	}

	for _, colour := range colours {

		// Smooth with sub divided
		bar.SetBarColour(colour)
		bar.SetFillColour(lastColour)
		fmt.Println("Smooth ProgressBar")

		progress = 0
		for progress < total {
			progress++
			time.Sleep(time.Duration(10) * time.Millisecond)

			// Update the progress bar
			bar.Draw(progress)
		}
		fmt.Println("\n")
		lastColour = colour
	}
}
