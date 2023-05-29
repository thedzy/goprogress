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
		Total:       total,
		BarText:     "=",
		Title:       "Loading",
		Terminators: []string{"[", "]"},
	})

	fmt.Println("1/4")
	for progress < total {
		progress++
		time.Sleep(25 * time.Millisecond)

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("2/4")
	bar.SetBarColour(goprogress.Yellow())
	bar.SetFillColour(goprogress.Blue())
	bar.SetBarText("")
	progress = 0
	for progress < total {
		progress++
		time.Sleep(25 * time.Millisecond)

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("3/4")
	bar.SetBarColour(goprogress.NoColour())
	bar.SetFillColour(goprogress.NoColour())
	bar.SetBarTextColour(goprogress.Red())
	animation := []string{
		"◠◡",
		"◡◠",
	}
	progress = 0
	for progress < total {
		progress++
		time.Sleep(25 * time.Millisecond)

		// Update the wait bar
		bar.SetBarText(animation[progress%2])
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("4 /4")
	bar.SetBarColour(goprogress.Green())
	bar.SetFillColour(goprogress.NoColour())
	bar.SetBarText("")
	bar.SetFillText("-")
	progress = 0
	for progress < total {
		progress++
		time.Sleep(25 * time.Millisecond)

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")
}
