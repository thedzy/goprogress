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

	options := goprogress.Options{
		Total:              total,
		Width:              50,
		BarText:            " ◢◤  ◢◤ ",
		FillText:           "|",
		Title:              "Loading",
		DynamicTextColours: true,
		Terminators:        []string{"[", "]"},
	}

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

	bar := goprogress.NewProgressBar(goprogress.StyleWait, options)

	fmt.Println("Wait Bar")
	progress = 0
	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("Simple Bar")
	bar.SetStyleSimple()
	progress = 0
	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("Train Bar")
	progress = 0
	bar.SetStyleTrain()
	bar.SetModeTimer()
	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("Detailed Bar")
	progress = 0
	bar.SetStyleDetailed()
	bar.SetModePortions()
	bar.SetFooter("apps")
	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the bar
		bar.SetBarText(apps[progress])
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("Smooth Bar")
	progress = 0
	bar.SetStyleSmooth()
	bar.SetModePercent()
	for progress < total {
		progress++
		time.Sleep(time.Duration(speeds[progress]*speedProportion) * time.Millisecond)

		// Update the bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

}
