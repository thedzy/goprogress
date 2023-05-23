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
	total := len(apps) - 1
	progress := 0

	goprogress.CreateProgress(goprogress.Options{
		Total:           total,
		Width:           80,
		Title:           "Searching",
		Footer:          "apps",
		BarColour:       []float32{1.0, 1.0, 0.0},
		FillColour:      []float32{0.0, 0.0, 0.5},
		LightTextColour: []float32{0.9, 0.0, 0.0},
		DarkTextColour:  []float32{0.0, 0.0, 0.9},
		InvertColours:   true,
		Terminators:     []string{"▕", "▏"},
		Mode:            3,
	})

	// Detailed with text
	fmt.Println("Detailed ProgressBar")
	for progress <= total {
		time.Sleep(time.Duration(rand.Intn(150-20)+20) * time.Millisecond)

		// Update the progress bar
		goprogress.DrawDetailedProgressBar(progress, goprogress.Options{
			Text: apps[progress]})
		progress++
	}
	fmt.Println("\n")
}
