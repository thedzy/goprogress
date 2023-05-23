package main

import (
	"fmt"
	"github.com/thedzy/goprogress"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// Set parameters for the demo
	total := 120
	progress := 0
	min, max := 20, 200
	var sum float32

	goprogress.CreateProgress(goprogress.Options{
		Total:         total,
		Width:         50,
		Title:         "Testing",
		FillCharacter: '_',
		BarColour:     []float32{1.0, 1.0, 0.0},
		FillColour:    []float32{0.0, 0.0, 0.5},
		Terminators:   []string{"▕", "▏"},
		Mode:          1,
	})

	var speeds []float32
	desiredSpeed := float32(10.0)
	progress = 0
	for progress < total+1 {
		progress++
		speeds = append(speeds, float32(rand.Intn(max-min+1)+min))
	}
	for _, num := range speeds {
		sum += num
	}
	// sum = sum / 1000
	speedProportion := desiredSpeed / sum * 1000

	fmt.Println("ProgressBars")

	// Train progress
	fmt.Println("Train ProgressBar")
	animation := []string{
		strings.Repeat("◟◧◨◝-", int(total/5)) + "◟◧◨◘◺",
		strings.Repeat("◜◧◨◞-", int(total/5)) + "◜◧◨◘◺",
	}
	animation = []string{
		strings.Repeat("◟◧◨◝-", int(total/5)) + "◟◧◨◘◺",
		strings.Repeat("◜◧◨◞-", int(total/5)) + "◜◧◨◘◺",
	}
	progress = 0
	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the progress bar
		goprogress.DrawTrainProgressBar(progress, goprogress.Options{
			Text:         animation[progress%2],
			IgnoreColour: true,
		})
	}
	fmt.Println("\n")

	// Train progress
	fmt.Println("Train ProgressBar")
	animation = []string{
		strings.Repeat("◷◧◨◷-", int(total/5)) + "◷◧◨◘◺",
		strings.Repeat("◶◧◨◶-", int(total/5)) + "◶◧◨◘◺",
		strings.Repeat("◵◧◨◵-", int(total/5)) + "◵◧◨◘◺",
		strings.Repeat("◴◧◨◴-", int(total/5)) + "◴◧◨◘◺",
	}
	progress = 0
	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the progress bar
		goprogress.DrawTrainProgressBar(progress, goprogress.Options{
			Text:         animation[progress%4],
			IgnoreColour: true,
		})
	}
	fmt.Println("\n")

	// Train progress
	fmt.Println("Train ProgressBar")
	progress = 0
	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the progress bar
		goprogress.DrawTrainProgressBar(progress, goprogress.Options{
			Text:          strings.Repeat("¸¸♬·¯·♩¸¸♪·¯·♫", total/15),
			BarColour:     []float32{1.0, 1.0, 0.0},
			FillCharacter: '-',
			IgnoreColour:  false,
		})
	}
	fmt.Println("\n")
}
