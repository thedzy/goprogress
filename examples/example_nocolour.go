package main

import (
	"fmt"
	"github.com/thedzy/goprogress"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"time"
)

func main() {
	dir := "/Applications"

	// Get a list of files in the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Filter files by application extension (.app)
	apps := make([]string, 0)
	for _, file := range files {
		if file.IsDir() && filepath.Ext(file.Name()) == ".app" {
			apps = append(apps, file.Name())
		}
	}

	// Set parameters for the demo
	total := 120
	progress := 0
	min, max := 20, 200
	var sum float32

	goprogress.CreateProgress(goprogress.Options{
		Total:               total,
		Width:               50,
		Title:               "Testing",
		Text:                "◢◤  ◢◤  ",
		Footer:              "",
		BarCharacter:        '=',
		BarDividerCharacter: '>',
		FillCharacter:       '-',
		BarColour:           []float32{1.0, 1.0, 0.0},
		FillColour:          []float32{0.0, 0.0, 0.5},
		LightTextColour:     []float32{0.9, 0.0, 0.0},
		DarkTextColour:      []float32{0.0, 0.0, 0.9},
		IgnoreColour:        true,
		Partials:            []string{" ", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},
		Terminators:         []string{"▕", "▏"},
		Mode:                2,
	})

	var speeds []float32
	desiredSpeed := float32(5.0)
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

	// Wait bar
	fmt.Println("WaitBar")
	progress = 0
	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the progress bar
		if progress%3 == 0 {
			goprogress.DrawWaitBar()
		}
	}
	fmt.Println("\n")

	goprogress.SetFooter("")
	goprogress.SetModePercent()

	// Simple
	fmt.Println("Simple ProgressBar")
	progress = 0
	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the progress bar
		goprogress.DrawSimpleProgressBar(progress)
	}
	fmt.Println("\n")

	goprogress.SetText("◧◨◡-◡◧◨◡-◡◧◨◡-◡◧◨◡-◡◧◨◡-◡◧◨◡-◡◧◨◡-◡◧◨◡-◡◧◨◡-◡◧◨◘◺")
	goprogress.SetFillCharacter('_')
	goprogress.SetModeTimer()
	goprogress.SetFooter("left")

	// Train progress
	fmt.Println("Train ProgressBar")
	progress = 0
	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the progress bar
		goprogress.DrawTrainProgressBar(progress)
	}
	fmt.Println("\n")

	goprogress.SetFooter("files")
	goprogress.SetModePortions()

	// Detailed with text
	fmt.Println("Detailed ProgressBar (does not respect the IgnoreColour option)")
	progress = 0
	index := 0
	for progress < total {
		if progress%3 == 0 {
			index += 1
		}

		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the progress bar
		goprogress.DrawDetailedProgressBar(progress, goprogress.Options{
			Text: apps[index]})
	}
	fmt.Println("\n")

	goprogress.SetFooter(" ")
	goprogress.SetModePercent()

	// Smooth with sub divided
	fmt.Println("Smooth ProgressBar")
	progress = 0
	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the progress bar
		goprogress.DrawSmoothProgressBar(progress)
	}
	fmt.Println("\n")

	fmt.Println("All completed")
}
