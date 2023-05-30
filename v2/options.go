package goprogress

import (
	"fmt"
	"reflect"
)

// settings that apply to all progress bars
type settings struct {
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

type Options struct {
	Total              *int
	Width              *int
	Title              *string
	Footer             *string
	BarText            *string
	FillText           *string
	Animated           *bool
	BarColour          *[]float32
	FillColour         *[]float32
	LightTextColour    *[]float32
	DarkTextColour     *[]float32
	BarTextColour      *[]float32
	FillTextColour     *[]float32
	DynamicTextColours *bool
	Partials           *[]string
	Terminators        *[]string
	Mode               *int8
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
// mergeSettings Merge two settings/options and validate the options
func mergeSettings(original settings, updated Options) settings {
	originalValue := reflect.ValueOf(&original).Elem()
	updatedValue := reflect.ValueOf(&updated).Elem()

	for i := 0; i < originalValue.NumField(); i++ {
		field := originalValue.Field(i)
		updatedField := updatedValue.Field(i)

		if updatedField.Kind() != reflect.Ptr || !updatedField.IsNil() {
			field.Set(updatedField.Elem())
		}
	}

	// original = validateOptions(original)

	return original
}

// validateOptions Check all options for valid parameters
func validateOptions(options Options) settings {
	var processedOptions settings

	if options.Total == nil || *options.Total < 1 {
		processedOptions.Total = 100
	} else {
		processedOptions.Total = *options.Total
	}

	if options.Width == nil || *options.Width < 1 {
		processedOptions.Width = 50
	} else {
		processedOptions.Width = *options.Width
	}

	if options.Title == nil {
		processedOptions.Title = ""
	} else {
		processedOptions.Title = *options.Title
	}

	if options.Footer == nil {
		processedOptions.Footer = ""
	} else {
		processedOptions.Footer = *options.Footer
	}

	if options.BarText == nil {
		processedOptions.BarText = " "
	} else {
		processedOptions.BarText = *options.BarText
	}

	if options.FillText == nil {
		processedOptions.FillText = " "
	} else {
		processedOptions.FillText = *options.FillText
	}

	if options.Animated == nil {
		processedOptions.Animated = false
	} else {
		processedOptions.Animated = *options.Animated
	}

	if options.BarColour == nil || len(*options.BarColour) < 3 {
		processedOptions.BarColour = NoColour()
	} else {
		processedOptions.BarColour = validationColour(*options.BarColour)
	}

	if options.FillColour == nil || len(*options.FillColour) < 3 {
		processedOptions.FillColour = NoColour()
	} else {
		processedOptions.FillColour = validationColour(*options.FillColour)
	}

	if options.BarTextColour == nil || len(*options.BarTextColour) < 3 {
		processedOptions.BarTextColour = NoColour()
	} else {
		processedOptions.BarTextColour = validationColour(*options.BarTextColour)
	}

	if options.FillTextColour == nil || len(*options.FillTextColour) < 3 {
		processedOptions.FillTextColour = NoColour()
	} else {
		processedOptions.FillTextColour = validationColour(*options.FillTextColour)
	}

	if options.LightTextColour == nil || len(*options.LightTextColour) < 3 {
		processedOptions.LightTextColour = White()
	} else {
		processedOptions.LightTextColour = validationColour(*options.LightTextColour)
	}

	if options.DarkTextColour == nil || len(*options.DarkTextColour) < 3 {
		processedOptions.DarkTextColour = Black()
	} else {
		processedOptions.DarkTextColour = validationColour(*options.DarkTextColour)
	}

	if options.DynamicTextColours == nil {
		processedOptions.DynamicTextColours = false
	} else {
		processedOptions.DynamicTextColours = *options.DynamicTextColours
	}

	if options.Partials == nil || len(*options.Partials) == 0 {
		processedOptions.Partials = []string{" ", "▏", "▎", "▍", "▌", "▋", "▊", "▉"}
	} else {
		processedOptions.Partials = *options.Partials
	}

	if options.Terminators == nil || len(*options.Terminators) < 2 {
		processedOptions.Terminators = []string{"▕", "▏"}
	} else {
		processedOptions.Terminators = *options.Terminators
	}

	if options.Mode == nil || *options.Mode < 1 {
		processedOptions.Mode = 1
	} else {
		processedOptions.Mode = *options.Mode
	}

	return processedOptions

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
