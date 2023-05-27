package goprogress

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"time"
)

// getColours Get colours from the options
func getColours(options Options) ColourCodes {

	// Colours
	barColour := getAnsiCode(options.BarColour[0], options.BarColour[1], options.BarColour[2], true)
	fillColour := getAnsiCode(options.FillColour[0], options.FillColour[1], options.FillColour[2], true)

	var barTextColour string
	var fillTextColour string
	if options.DynamicTextColours {
		darkTextColour := getAnsiCode(options.DarkTextColour[0], options.DarkTextColour[1], options.DarkTextColour[2], false)
		lightTextColour := getAnsiCode(options.LightTextColour[0], options.LightTextColour[1], options.LightTextColour[2], false)

		barContrast := colourContrast(options.BarColour[0], options.BarColour[1], options.BarColour[2])
		fillContrast := colourContrast(options.FillColour[0], options.FillColour[1], options.FillColour[2])
		if barContrast > 0.5 {
			barTextColour = darkTextColour
		} else {
			barTextColour = lightTextColour
		}
		if fillContrast > 0.5 {
			fillTextColour = darkTextColour
		} else {
			fillTextColour = lightTextColour
		}

	} else {
		barTextColour = getAnsiCode(options.BarTextColour[0], options.BarTextColour[1], options.BarTextColour[2], false)
		fillTextColour = getAnsiCode(options.FillTextColour[0], options.FillTextColour[1], options.FillTextColour[2], false)
	}

	return ColourCodes{
		BarColour:      barColour,
		FillColour:     fillColour,
		BarTextColour:  barTextColour,
		FillTextColour: fillTextColour,
	}
}

// getAnsiCode Get the ansi code for a colour as string
func getAnsiCode(red float32, green float32, blue float32, back bool) string {
	if red == -1.0 && green == -1.0 && blue == -1.0 {
		return ""
	}

	var ansiCode = 0
	if int(red*5) == int(green*5) && int(green*5) == int(blue*5) {
		ansiCode = 232 + int(red*23)
	} else {
		ansiCode = int(16 + (36 * math.Round(float64(red*5))) + (6 * math.Round(float64(green*5))) + math.Round(float64(blue*5)))
	}
	var escCode string
	if back {
		escCode = ansiBackgroundCode(ansiCode)
	} else {
		escCode = ansiForegroundCode(ansiCode)
	}
	return escCode
}

// ansiBackgroundCode Get the ansi code string with escape for the background
func ansiBackgroundCode(code int) string {
	return "\033[48;5;" + strconv.Itoa(code) + "m"
}

// ansiForegroundCode Get the ansi code string with escape for the foreground
func ansiForegroundCode(code int) string {
	return "\033[38;5;" + strconv.Itoa(code) + "m"
}

// colourContrast Get the perceived contrast on a scale of 0.0 - 1.0
func colourContrast(red float32, green float32, blue float32) float32 {
	colourValue := ((red * 299) + (green * 587) + (blue * 114)) / 1000

	return colourValue
}

// getFooter Get the footer based on the choice and do the calculation
func getFooter(progress int, timer time.Time, options Options) string {
	var footer string

	switch options.Mode {
	case ModePercent:
		// Calculate the percentage
		percentage := int(float64(progress) / float64(options.Total) * 100)
		footer = fmt.Sprintf(" %3d%%", percentage)

	case ModeTimer:
		// Calculate time
		timeRemainingSlices := float32(time.Since(timer).Seconds()) / float32(progress)
		timeRemaining := timeRemainingSlices * (float32(options.Total) - float32(progress))

		// Get minutes and seconds remaining
		minutes := int(timeRemaining) / 60
		seconds := int(timeRemaining) % 60

		footer = fmt.Sprintf("%3.1d:%02d", minutes, seconds)
	case ModeProportion:
		numLength := len(strconv.Itoa(options.Total))
		footer = fmt.Sprintf("%*.0d/%*.0d", numLength, progress, numLength, options.Total)
	case ModeNone:
		fallthrough
	default:
		footer = ""
	}
	return footer
}

// getFinalWidth Calculate the complete width of the bar accounting for all other items
func getFinalWidth(setWidth int, title string, footer string, terminators []string, adjustment int) int {
	if setWidth < 10 {
		setWidth = 10
	}
	fullWidth := setWidth + len(title) + len(footer) + len(terminators[0]) + len(terminators) + adjustment
	if fullWidth > getWindowWidth() {
		setWidth = getWindowWidth() - len(title) - len(footer) - len(terminators[0]) - len(terminators) - adjustment
	} else {
		setWidth -= adjustment
	}

	return setWidth
}

// getWindowWidth Get the width of the window
func getWindowWidth() int {
	winSize, err := getWinSize()

	if err != nil {
		return 100
	}

	return int(winSize.Col - 1)
}

// validationColour Validate a colour and modify it to conformity
func validationColour(colour []float32) []float32 {
	// Verify that no colour is above 1.0 or below 0.0

	if colour[0] == -1.0 && colour[0] == colour[1] && colour[1] == colour[2] {
		return colour
	}

	for index, value := range colour[0:3] {
		if value > 1 {
			fmt.Println("Error: One or more colour values are greater than 1 and been set to 1")
			colour[index] = 1

		} else if value < 0 {
			fmt.Println("Error: One or more colour values are greater than 0 and been set to 0")
			colour[index] = 0
		}
	}

	return colour
}

// stringToArray Turn a string into an array to strings for slicing or long unicodes
func stringToArray(text string) []string {
	var textChars []string
	for _, char := range text {
		textChars = append(
			textChars,
			string(char),
		)
	}

	return textChars
}

// arrayToString Restore an array to a string
func arrayToString(strArray []string) string {
	return strings.Join(strArray, "")
}

// stringArraySizer Resize, repeat and fit and align
func stringArraySizer(strArray []string, length int, repeat bool, alignRight bool) []string {
	strLength := len(strArray)
	if strLength == 0 {
		paddedArray := make([]string, length)
		for i := range paddedArray {
			paddedArray[i] = " "
		}
		return paddedArray
	}

	if repeat {
		var newStrArray []string
		for x := 0; x < (length/strLength)+1; x++ {
			newStrArray = append(newStrArray, strArray...)
		}
		strLength = len(newStrArray)
		strArray = newStrArray
	}

	if strLength > length {
		if alignRight {
			return strArray[strLength-length:]
		} else {
			return strArray[0:length]
		}
	} else {
		paddedArray := make([]string, length-strLength)
		for i := range paddedArray {
			paddedArray[i] = " "
		}
		if alignRight {
			return append(paddedArray, strArray...)
		} else {
			return append(strArray, paddedArray...)
		}
	}
}
