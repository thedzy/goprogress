package main

import (
	"fmt"
	"github.com/thedzy/goprogress/v2"
	"io/ioutil"
	"path/filepath"
	"strings"
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
	total := len(apps) - 1
	progress := 0

	bar := goprogress.NewProgressBar(goprogress.StyleDetailed, goprogress.Options{
		Total:              total,
		Width:              80,
		Title:              "Searching",
		Footer:             "apps",
		BarColour:          []float32{1.0, 1.0, 0.0},
		FillColour:         []float32{0.0, 0.0, 0.5},
		LightTextColour:    []float32{0.9, 0.0, 0.0},
		DarkTextColour:     []float32{0.0, 0.0, 0.9},
		DynamicTextColours: true,
		Terminators:        []string{"▕", "▏"},
		Mode:               3,
	})

	bar.SetMode(goprogress.ModeProportion)

	fmt.Println("Detailed Bar 1/3")
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the wait bar
		bar.SetBarText(apps[progress])
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("Detailed Bar 2/3")
	bar.SetBarText(strings.Repeat("◠◡", total/2))
	progress = 0
	for progress < total {
		progress++
		time.Sleep(100 * time.Millisecond)

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

	fmt.Println("Detailed Bar 3/3")
	progress = 0
	bar.SetTextColours(goprogress.Black(), goprogress.White(), goprogress.Black(), goprogress.White())
	bar.SetColours(goprogress.White(), goprogress.Black()).
		SetDynamicTextColours(false).
		SetBarText(strings.Repeat("¸¸♬·¯·♩¸¸♪·¯·♫", total/14))

	for progress < total {
		progress++
		time.Sleep(time.Duration(100 * time.Millisecond))

		// Update the wait bar
		bar.Draw(progress)
	}
	fmt.Println("\n")

}
