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
	options Options
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

	if options.Total == 0 {
		options.Total = 100
	}

	if options.Width == 0 {
		options.Width = 50
	}

	if options.BarColour == nil || len(options.BarColour) < 3 {
		options.BarColour = NoColour()
	} else {
		options.BarColour = validationColour(options.BarColour)
	}

	if options.FillColour == nil || len(options.FillColour) < 3 {
		options.FillColour = NoColour()
	} else {
		options.FillColour = validationColour(options.FillColour)
	}

	if options.BarTextColour == nil || len(options.BarTextColour) < 3 {
		options.BarTextColour = NoColour()
	} else {
		options.BarTextColour = validationColour(options.BarTextColour)
	}

	if options.FillTextColour == nil || len(options.FillTextColour) < 3 {
		options.FillTextColour = NoColour()
	} else {
		options.FillTextColour = validationColour(options.FillTextColour)
	}

	if options.LightTextColour == nil || len(options.LightTextColour) < 3 {
		options.LightTextColour = White()
	} else {
		options.LightTextColour = validationColour(options.LightTextColour)
	}

	if options.DarkTextColour == nil || len(options.DarkTextColour) < 3 {
		options.DarkTextColour = Black()
	} else {
		options.DarkTextColour = validationColour(options.DarkTextColour)
	}

	if len(options.Partials) == 0 {
		options.Partials = []string{" ", "▏", "▎", "▍", "▌", "▋", "▊", "▉"}
	}

	if len(options.Terminators) < 2 {
		options.Terminators = []string{"▕", "▏"}
	}

	if options.Mode == 0 {
		options.Mode = 1
	}

	progressBar := bar{
		style:   style,
		options: options,
	}

	return &progressBar
}

// ResetTimer Reset the timer if a bar does not complete
func (o *bar) ResetTimer() {
	o.timer = time.Time{}
}

// Draw the progress with text (and colour)
func (o *bar) Draw(progress int, options ...Options) {

	_ = options

	counter := o.nextDigit()
	_ = counter
	if o.timer.IsZero() {
		o.timer = time.Now()
	}

	colours := getColours(o.options)

	footer := getFooter(progress, o.timer, o.options)
	footer = fmt.Sprintf("%s %s", footer, o.options.Footer)

	// Keep the minimum width of the bar to 25
	width := getFinalWidth(o.options.Width, o.options.Title, footer, o.options.Terminators, 0)

	// Calculate the fill of the bar
	barWidth := int(float64(progress) / float64(o.options.Total) * float64(width))
	fillWidth := width - barWidth

	var bar string
	var fill string
	switch o.style {
	case StyleWait:
		barText := stringToArray(o.options.BarText)
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
		barText := stringToArray(o.options.BarText)
		barFill := stringToArray(o.options.FillText)

		bar = arrayToString(stringArraySizer(barText, barWidth, false, true))
		fill = arrayToString(stringArraySizer(barFill, fillWidth, false, false))
	case StyleDetailed:
		// Detailed
		barText := stringArraySizer(stringToArray(o.options.BarText), width, false, false)

		bar = arrayToString(barText[0:barWidth])
		fill = arrayToString(barText[barWidth:])

		alterOptions := o.options
		if colours.BarColour == "" {
			alterOptions.BarColour = LtGrey()
		}
		if colours.FillColour == "" {
			alterOptions.FillColour = Grey()
		}

		colours = getColours(alterOptions)
	case StyleSmooth:
		// Smooth
		partialsCount := len(o.options.Partials)
		width = width * partialsCount
		barWidth := int(float64(progress) / float64(o.options.Total) * float64(width))

		quotient := barWidth / partialsCount
		remainder := barWidth % partialsCount

		bar = strings.Repeat("▉", quotient)
		fill = arrayToString(
			stringToArray(
				o.options.Partials[remainder] +
					strings.Repeat(" ",
						(width/partialsCount)-quotient),
			)[0 : (width/partialsCount)-quotient],
		)

		alterOptions := o.options
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
		barText := stringToArray(o.options.BarText)
		barFill := stringToArray(o.options.FillText)

		bar = arrayToString(stringArraySizer(barText, barWidth, true, true))
		fill = arrayToString(stringArraySizer(barFill, fillWidth, true, false))
	}
	// Prevent bar for exceeding 100%
	if progress >= o.options.Total {
		progress = o.options.Total
		o.timer = time.Time{}
	}

	// Update the progress bar
	fmt.Printf("\r %s %s%s%s%s%s%s%s%s%s%s%s\r",
		o.options.Title,
		o.options.Terminators[0],
		colours.BarColour,
		colours.BarTextColour,
		bar,
		"\u001B[0;m",
		colours.FillColour,
		colours.FillTextColour,
		fill,
		"\u001B[0;m",
		o.options.Terminators[1],
		footer)
}

func (o *bar) nextDigit() int64 {
	o.counter += 1
	if o.counter == math.MaxInt64 {
		o.counter = 0
	}
	return o.counter
}
