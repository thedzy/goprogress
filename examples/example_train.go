package main

import (
	"fmt"
	"github.com/thedzy/goprogress/v2"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// Set parameters for the demo
	total := 120
	progress := 0
	width := 50
	min, max := 20, 200
	var sum float32

	bar := goprogress.NewProgressBar(goprogress.StyleTrain, goprogress.Options{
		Total:          total,
		Width:          width,
		Title:          "Loading",
		BarText:        "=",
		FillText:       strings.Repeat("_", width),
		BarTextColour:  goprogress.Yellow(),
		FillTextColour: goprogress.DkGrey(),
		Terminators:    []string{"▕", "▏"},
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

	fmt.Println("Train bar 1/4")
	animation := []string{
		strings.Repeat("◟◧◨◝-", int(total/5)) + "◟◧◨◘◺",
		strings.Repeat("◜◧◨◞-", int(total/5)) + "◜◧◨◘◺",
	}
	progress = 0
	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the wait bar
		bar.SetBarText(animation[progress%2])
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("Train bar 2/4")
	animation = []string{
		strings.Repeat("◷◧◨◷-", int(total/5)) + "◷◧◨◘◺",
		strings.Repeat("◶◧◨◶-", int(total/5)) + "◶◧◨◘◺",
		strings.Repeat("◵◧◨◵-", int(total/5)) + "◵◧◨◘◺",
		strings.Repeat("◴◧◨◴-", int(total/5)) + "◴◧◨◘◺",
	}
	progress = 0
	bar.SetBarTextColour(goprogress.Grey())
	bar.SetFillText("")

	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the wait bar
		bar.SetBarText(animation[progress%4])
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("Train bar 3/4")

	progress = 0
	bar.SetBarTextColour(goprogress.Black())
	bar.SetBarColour(goprogress.White())
	bar.SetBarText(strings.Repeat("¸¸♬·¯·♩¸¸♪·¯·♫", total/15))
	bar.SetFillText(">")

	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("Train bar 4/4")

	progress = 0
	bar.SetBarTextColour(goprogress.Black())
	bar.SetBarColour(goprogress.White())
	bar.SetFillColour(goprogress.Black())
	bar.SetFillTextColour(goprogress.White())
	bar.SetBarText("<")
	bar.SetFillText(">")

	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

}
