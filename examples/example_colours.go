package main

import (
	"fmt"
	"github.com/thedzy/goprogress/v2"
	"time"
)

func main() {
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
		bar.SetBarColour(colour)
		bar.SetFillColour(lastColour)

		// Smooth with fraction al character widths
		fmt.Println("Smooth ProgressBar")

		progress = 0
		for progress < total {
			progress++
			time.Sleep(time.Duration(10) * time.Millisecond)

			// Update the progress bar
			bar.Draw(progress, goprogress.Options{
				FillColour: goprogress.Black(),
			})
		}
		fmt.Println("\n")

	}

	// Smooth with fraction al character widths
	fmt.Println("Smooth ProgressBar with colour shifting")
	progress = 0
	fractional := 1.0 / (float32(total) + 1)
	for progress < total {
		progress++
		time.Sleep(time.Duration(10) * time.Millisecond)

		// Update the progress bar
		bar.Draw(progress, goprogress.Options{
			BarColour:  []float32{1.0 - float32(progress)*fractional, 0.0, float32(progress) * fractional},
			FillColour: goprogress.Black(),
		})
	}
	fmt.Println("\n")

	fmt.Println("Smooth ProgressBar with every ansi code")
	progress = 0
	bar.SetTotal(239)
	progress = 0

	timeMultiplier := 150
	x := 16
	for i := 0; i < 18; i++ {
		progress++
		time.Sleep(time.Duration(timeMultiplier) * time.Millisecond)
		bar.SetTitle(fmt.Sprintf("%3d", x)).SetBarColour(goprogress.GetRGB(x)).Draw(progress)

		for y := 0; y < 5; y++ {
			x = x + 36
			progress++
			time.Sleep(time.Duration(timeMultiplier) * time.Millisecond)
			bar.SetTitle(fmt.Sprintf("%3d", x)).SetBarColour(goprogress.GetRGB(x)).Draw(progress)
		}
		x++
		progress++
		time.Sleep(time.Duration(timeMultiplier) * time.Millisecond)
		bar.SetTitle(fmt.Sprintf("%3d", x)).SetBarColour(goprogress.GetRGB(x)).Draw(progress)

		for z := 0; z < 5; z++ {
			x = x - 36
			progress++
			time.Sleep(time.Duration(timeMultiplier) * time.Millisecond)
			bar.SetTitle(fmt.Sprintf("%3d", x)).SetBarColour(goprogress.GetRGB(x)).Draw(progress)
		}
		x++
	}
	for x = 255; x >= 231; x-- {
		progress++
		time.Sleep(time.Duration(timeMultiplier) * time.Millisecond)
		bar.SetTitle(fmt.Sprintf("%3d", x)).SetBarColour(goprogress.GetRGB(x)).Draw(progress)
	}

}
