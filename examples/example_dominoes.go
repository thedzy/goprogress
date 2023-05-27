package main

import (
	"fmt"
	"github.com/thedzy/goprogress/v2"
	"strings"
	"time"
)

func main() {
	total := 100
	progress := 0

	bar := goprogress.NewProgressBar(goprogress.StyleTrain, goprogress.Options{
		Total:       total,
		Width:       80,
		BarText:     strings.Repeat("_", 80),
		FillText:    "/" + strings.Repeat("|", 80),
		Title:       "Loading",
		Terminators: []string{"_", ""},
		Mode:        goprogress.ModePercent,
	})

	bar.SetMode(goprogress.ModeProportion)

	fmt.Println("Detailed Bar")
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")
}
