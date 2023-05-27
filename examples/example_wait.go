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

	fmt.Println("1/4 Bar")
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

	fmt.Println("2/4 Bar")
	progress = 0
	bar.SetBarText(" ◢◤")
	bar.SetBarTextColour(goprogress.Yellow())
	bar.SetBarColour(goprogress.DkBlue())
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("3/4 Bar")
	progress = 0

	bar.SetBarText(" ▷▶ ▸")
	bar.SetBarTextColour(goprogress.Yellow())
	bar.SetBarColour(goprogress.DkBlue())
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("4/4 Bar")
	progress = 0

	bar.SetBarText("◠◡")
	bar.SetBarTextColour(goprogress.Yellow())
	bar.SetBarColour(goprogress.DkRed())
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

}
