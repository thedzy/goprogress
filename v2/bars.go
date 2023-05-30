package goprogress

import (
	"fmt"
	"math"
	"strings"

	"time"
)

type bar struct {
	style   int8
	counter int64
	options settings
	timer   time.Time
}

type ColourCodes struct {
	BarColour      string
	FillColour     string
	BarTextColour  string
	FillTextColour string
}

// Initialise variables and settings
func init() {
}

// NewProgressBar Create a new instance of a progress bar
func NewProgressBar(style int8, options Options) *bar {

	processedOptions := validateOptions(options)

	progressBar := bar{
		style:   style,
		options: processedOptions,
	}

	return &progressBar
}

// ResetTimer Reset the timer if a bar does not complete
func (o *bar) ResetTimer() {
	o.timer = time.Time{}
}

// Draw the progress with text (and colour)
func (o *bar) Draw(progress int, overrides ...Options) {

	var options settings
	if len(overrides) > 0 {
		options = mergeSettings(o.options, overrides[0])
	} else {
		options = o.options
	}

	counter := o.nextDigit()
	_ = counter
	if o.timer.IsZero() {
		o.timer = time.Now()
	}

	colours := getColours(options)

	footer := getFooter(progress, o.timer, options)
	footer = fmt.Sprintf("%s %s", footer, options.Footer)

	// Keep the minimum width of the bar to 25
	width := getFinalWidth(options.Width, options.Title, footer, options.Terminators, 0)

	// Calculate the fill of the bar
	barWidth := int(float64(progress) / float64(options.Total) * float64(width))
	fillWidth := width - barWidth

	var bar string
	var fill string
	switch o.style {
	case StyleWait:
		barText := stringToArray(options.BarText)

		// Ensure wait bar is not empty
		if len(strings.TrimSpace(options.BarText)) == 0 {
			barText = []string{"/", " "}
		}
		barTextLen := len(barText)

		barText = stringArraySizer(
			barText,
			int(math.Ceil(float64(width)/float64(barTextLen)))*barTextLen,
			true, true,
		)

		offset := barTextLen - int(counter%int64(barTextLen))
		barText = append(barText[offset:], barText...)

		bar = arrayToString(barText[0:width])
		footer = strings.Repeat(" ", len(footer))
	case StyleTrain:
		barText := stringToArray(options.BarText)
		barFill := stringToArray(options.FillText)

		bar = arrayToString(stringArraySizer(barText, barWidth, false, true))
		fill = arrayToString(stringArraySizer(barFill, fillWidth, false, false))
	case StyleDetailed:
		barText := stringArraySizer(stringToArray(options.BarText), width, false, false)

		bar = arrayToString(barText[0:barWidth])
		fill = arrayToString(barText[barWidth:])

		alterOptions := options
		if colours.BarColour == "" {
			alterOptions.BarColour = LtGrey()
		}
		if colours.FillColour == "" {
			alterOptions.FillColour = Grey()
		}

		colours = getColours(alterOptions)
	case StyleSmooth:
		partialsCount := len(options.Partials)
		width = width * partialsCount
		barWidth := int(float64(progress) / float64(options.Total) * float64(width))

		quotient := barWidth / partialsCount
		remainder := barWidth % partialsCount

		bar = strings.Repeat("▉", quotient)
		fill = arrayToString(
			stringToArray(
				options.Partials[remainder] +
					strings.Repeat(" ",
						(width/partialsCount)-quotient),
			)[0 : (width/partialsCount)-quotient],
		)

		alterOptions := options
		if colours.BarColour == "" {
			alterOptions.BarColour = Grey()
		}
		alterOptions.BarTextColour = alterOptions.BarColour
		alterOptions.FillTextColour = alterOptions.BarColour
		alterOptions.DynamicTextColours = false

		colours = getColours(alterOptions)
	case StyleSimple:
		fallthrough
	default:
		barText := stringToArray(options.BarText)
		barFill := stringToArray(options.FillText)

		bar = arrayToString(stringArraySizer(barText, barWidth, true, true))
		fill = arrayToString(stringArraySizer(barFill, fillWidth, true, false))
	}

	// Prevent bar for exceeding 100%
	if progress >= options.Total {
		progress = options.Total
		o.timer = time.Time{}
	}

	// Update the progress bar
	fmt.Printf("\r %s %s%s%s%s%s%s%s%s%s%s%s\r",
		options.Title,
		options.Terminators[0],
		colours.BarColour,
		colours.BarTextColour,
		bar,
		"\u001B[0;m",
		colours.FillColour,
		colours.FillTextColour,
		fill,
		"\u001B[0;m",
		options.Terminators[1],
		footer)
}

func (o *bar) nextDigit() int64 {
	o.counter += 1
	if o.counter == math.MaxInt64 {
		o.counter = 0
	}
	return o.counter
}
