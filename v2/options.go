package goprogress

import (
	"fmt"
	"reflect"
)

// Options that apply to all progress bars
type Options struct {
	Total              int
	Width              int
	Title              string
	Footer             string
	BarText            string
	FillText           string
	Animated           bool
	BarColour          []float32
	FillColour         []float32
	LightTextColour    []float32
	DarkTextColour     []float32
	BarTextColour      []float32
	FillTextColour     []float32
	DynamicTextColours bool
	Partials           []string
	Terminators        []string
	Mode               int8
}

// SetStyleSimple Set style to simple
func (o *bar) SetStyleSimple() {
	o.style = StyleSimple
}
func (o *bar) GetStyleSimple() bool {
	return o.style == StyleSimple
}

// SetStyleTrain Set style to train
func (o *bar) SetStyleTrain() {
	o.style = StyleTrain
}
func (o *bar) GetStyleTrain() bool {
	return o.style == StyleTrain
}

// SetStyleDetailed Set style to detailed
func (o *bar) SetStyleDetailed() {
	o.style = StyleDetailed
}
func (o *bar) GetStyleDetailed() bool {
	return o.style == StyleDetailed
}

// SetStyleSmooth Set style to smooth
func (o *bar) SetStyleSmooth() {
	o.style = StyleSmooth
}
func (o *bar) GetStyleSmooth() bool {
	return o.style == StyleSmooth
}

// SetStyle Set to display style
func (o *bar) SetStyle(mode int8) {
	o.style = mode
}
func (o *bar) GetStyle() int8 {
	return o.style
}

// SetOptions Set total (max value of the progress bar)
func (o *bar) SetOptions(options Options) {
	o.options = options
}
func (o *bar) GetOptions() Options {
	return o.options
}

// SetTotal Set total (max value of the progress bar)
func (o *bar) SetTotal(total int) {
	o.options.Total = total
}
func (o *bar) GetTotal() int {
	return o.options.Total
}

// SetWidth Set the width of the progress bar
func (o *bar) SetWidth(width int) {
	o.options.Width = width
}
func (o *bar) GetWidth() int {
	return o.options.Width
}

// SetTitle Set title at the beginning of the bar
func (o *bar) SetTitle(title string) {
	o.options.Title = title
}
func (o *bar) GetTitle() string {
	return o.options.Title
}

// SetFooter Set text at the end
func (o *bar) SetFooter(footer string) {
	o.options.Footer = footer
}
func (o *bar) GetFooter() string {
	return o.options.Footer
}

// SetBarText Set text at the end
func (o *bar) SetBarText(text string) {
	o.options.BarText = text
}
func (o *bar) GetBarText() string {
	return o.options.BarText
}

// SetFillText Set text at the end
func (o *bar) SetFillText(text string) {
	o.options.FillText = text
}
func (o *bar) GetFillText() string {
	return o.options.FillText
}

// SetBarColour Set bar colour
func (o *bar) SetBarColour(colour []float32) {
	o.options.BarColour = validationColour(colour)
}
func (o *bar) GetBarColour() []float32 {
	return o.options.BarColour
}

// SetFillColour Set the fill colour (opposite of the bar)
func (o *bar) SetFillColour(colour []float32) {
	o.options.FillColour = validationColour(colour)
}
func (o *bar) GetFillColour() []float32 {
	return o.options.FillColour
}

// SetLightTextColour Set text colour that appears over dark colours
func (o *bar) SetLightTextColour(colour []float32) {
	o.options.LightTextColour = validationColour(colour)
}
func (o *bar) GetLightTextColour() []float32 {
	return o.options.LightTextColour
}

// SetDarkTextColour Set text colour that appears over light colours
func (o *bar) SetDarkTextColour(colour []float32) {
	o.options.DarkTextColour = validationColour(colour)
}
func (o *bar) GetDarkTextColour() []float32 {
	return o.options.DarkTextColour
}

// SetBarTextColour Set text colour that appears over dark colours
func (o *bar) SetBarTextColour(colour []float32) {
	o.options.BarTextColour = validationColour(colour)
}
func (o *bar) GetBarTextColour() []float32 {
	return o.options.BarTextColour
}

// SetFillTextColour Set text colour that appears over light colours
func (o *bar) SetFillTextColour(colour []float32) {
	o.options.FillTextColour = validationColour(colour)
}
func (o *bar) GetFillTextColour() []float32 {
	return o.options.FillTextColour
}

// SetDynamicTextColours Forgo the light and dark colour and use the fill/bar colours
func (o *bar) SetDynamicTextColours(dynamic bool) {
	o.options.DynamicTextColours = dynamic

}
func (o *bar) GetDynamicTextColours() bool {
	return o.options.DynamicTextColours
}

// SetPartials Set partials for the smooth bar
func (o *bar) SetPartials(partials []string) {
	o.options.Partials = partials
}
func (o *bar) GetPartials() []string {
	return o.options.Partials
}

// SetTerminators Set the start and end character for the bar
func (o *bar) SetTerminators(start string, end string) {
	o.options.Terminators = []string{start, end}
}
func (o *bar) GetTerminators() []string {
	return o.options.Terminators
}

// SetModeNone Set to hide progress/timer/proportion
func (o *bar) SetModeNone() {
	o.options.Mode = ModeNone
}
func (o *bar) GetModeNone() bool {
	return o.options.Mode == ModeNone
}

// SetModePercent Set to display progress
func (o *bar) SetModePercent() {
	o.options.Mode = ModePercent
}
func (o *bar) GetModePercent() bool {
	return o.options.Mode == ModePercent
}

// SetModeTimer Set to display to timer
func (o *bar) SetModeTimer() {
	o.options.Mode = ModeTimer
}
func (o *bar) GetModeTimer() bool {
	return o.options.Mode == ModeTimer
}

// SetModePortions Set to display portions
func (o *bar) SetModePortions() {
	o.options.Mode = ModeProportion
}
func (o *bar) GetModePortions() bool {
	return o.options.Mode == ModeProportion
}

// SetMode Set to display the mode
func (o *bar) SetMode(mode int8) {
	o.options.Mode = mode
}
func (o *bar) GetMode() int8 {
	return o.options.Mode
}

// mergeSettings Merge two settings/options and validate the options
func mergeSettings(original Options, updated Options) Options {
	// Assign to some variables for modification
	originalValue := reflect.ValueOf(&original).Elem()
	updatedValue := reflect.ValueOf(updated)

	for i := 0; i < originalValue.NumField(); i++ {
		field := originalValue.Field(i)
		updatedField := updatedValue.Field(i)

		if !reflect.ValueOf(updatedField.Interface()).IsZero() {
			field.Set(updatedField)
		}
	}

	// Verify we have 2 terminators
	if len(original.Terminators) < 2 {
		original.Terminators = append(original.Terminators, "", "")
	}

	// Keep the minimum width of the bar to 25 and the window width
	if original.Width < 10 {
		original.Width = 10
	}
	if original.Width > getWindowWidth() {
		original.Width = getWindowWidth()
	}

	// Verify that no colour is above 1.0 or below 0.0
	colours := [][]float32{original.FillColour, original.BarColour, original.LightTextColour, original.DarkTextColour}
	for _, array := range colours {
		for index, value := range array[0:3] {
			if value > 1 {
				fmt.Println("Error: One or more colour values are greater than 1 and been set to 1")
				array[index] = 1
			}
			if value < 0 {
				fmt.Println("Error: One or more colour values are greater than 0 and been set to 0")
				array[index] = 0
			}
		}
	}

	// Verify we have at least one partial
	if len(original.Partials) == 0 {
		original.Partials = []string{" "}
	}

	return original
}

// validateOptions Check all options for valid parameters
func validateOptions(options Options) Options {

	if options.Total < 1 {
		options.Total = 100
	}

	if options.Width < 1 {
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

	if options.Mode < 1 {
		options.Mode = 1
	}

	return options

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
