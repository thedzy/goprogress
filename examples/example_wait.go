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

	bar := goprogress.NewProgressBar(goprogress.StyleWait, goprogress.Options{
		Total: total,
		Width: 25,
		Title: "Loading",
	})

	fmt.Println("Wait bar 1/4")
	bar.SetBarText("▁▂▂▃▄▅▆▇▇██▇▇▆▅▄▃▂▂▁")
	bar.SetBarTextColour(goprogress.Red())
	bar.SetBarColour(goprogress.Black())
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("Wait bar 2/4")
	progress = 0
	bar.SetBarText(" ◢◤").
		SetBarTextColour(goprogress.Yellow()).
		SetBarColour(goprogress.DkBlue())
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("Wait bar 3/4")
	progress = 0

	bar.SetBarText(" ▷▶ ▸").
		SetBarTextColour(goprogress.Yellow()).
		SetBarColour(goprogress.DkBlue())
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("Wait bar 4/4")
	progress = 0

	bar.SetBarText("◠◡").
		SetBarTextColour(goprogress.Yellow()).
		SetBarColour(goprogress.DkRed())
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

}
