package goprogress

import (
	"fmt"
	"math"
	"strconv"

	// "golang.org/x/sys/windows"
	"golang.org/x/sys/unix"
	"os"
	"reflect"
	"strings"
	"time"
)

// Options that apply to all progress bars
type Options struct {
	Total               int
	Progress            int
	Width               int
	Title               string
	Text                string
	Footer              string
	BarCharacter        rune
	BarDividerCharacter rune
	FillCharacter       rune
	BarColour           []float32
	FillColour          []float32
	LightTextColour     []float32
	DarkTextColour      []float32
	InvertColours       bool
	IgnoreColour        bool
	Partials            []string
	Terminators         []string
	Mode                int8
}

// Define globals
var (
	settings = Options{}
	timer    = time.Time{}
)

type ColourCodes struct {
	BarColour      string
	FillColour     string
	BarTextColour  string
	FillTextColour string
}

// Initialise variables and settings
func init() {
	settings = Options{
		Total:               100,
		Progress:            0,
		Width:               getWindowWidth(),
		Title:               "",
		Text:                "",
		Footer:              "",
		BarCharacter:        '=',
		BarDividerCharacter: '=',
		FillCharacter:       ' ',
		BarColour:           []float32{1.0, 1.0, 1.0},
		FillColour:          []float32{0.0, 0.0, 0.0},
		LightTextColour:     []float32{1.0, 1.0, 1.0},
		DarkTextColour:      []float32{0.0, 0.0, 0.0},
		InvertColours:       false,
		IgnoreColour:        false,
		Partials:            []string{" ", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},
		Terminators:         []string{"▕", "▏"},
		Mode:                1,
	}
}

// CreateProgress Initialise settings for user that will be used across progress bars
func CreateProgress(options Options) {
	settings = mergeSettings(settings, options)
}

// SetTotal Set total (max value of the progress bar)
func SetTotal(total int) {
	options := Options{
		Total: total}
	settings = mergeSettings(settings, options)
}
func GetTotal() int {
	return settings.Total
}

// SetWidth Set the width of the progress bar
func SetWidth(width int) {
	options := Options{
		Width: width}
	settings = mergeSettings(settings, options)
}
func GetWidth() int {
	return settings.Width
}

// SetTitle Set title at the beginning of the bar
func SetTitle(title string) {
	options := Options{
		Title: title}
	settings = mergeSettings(settings, options)
}
func GetTitle() string {
	return settings.Title
}

// SetText Set text where title is use
func SetText(text string) {
	options := Options{
		Text: text}
	settings = mergeSettings(settings, options)
}
func GetText() string {
	return settings.Text
}

// SetFooter Set text at the end
func SetFooter(footer string) {
	options := Options{
		Footer: footer}
	settings = mergeSettings(settings, options)
}
func GetFooter() string {
	return settings.Footer
}

// SetBarCharacter Set the bar rune where used
func SetBarCharacter(char rune) {
	options := Options{
		BarCharacter: char}
	settings = mergeSettings(settings, options)
}
func GetBarCharacter() rune {
	return settings.BarCharacter
}

// SetBarDivCharacter Set the tip of the bar rune where used
func SetBarDivCharacter(char rune) {
	options := Options{
		BarDividerCharacter: char}
	settings = mergeSettings(settings, options)
}
func GetBarDivCharacter() rune {
	return settings.BarDividerCharacter
}

// SetFillCharacter Set the fill rune of the bar where used
func SetFillCharacter(char rune) {
	options := Options{
		FillCharacter: char}
	settings = mergeSettings(settings, options)
}
func GetFillCharacter() rune {
	return settings.FillCharacter
}

// SetBarColour Set bar colour
func SetBarColour(red float32, green float32, blue float32) {
	options := Options{
		BarColour: []float32{red, green, blue}}
	settings = mergeSettings(settings, options)
}
func GetBarColour() []float32 {
	return settings.BarColour
}

// SetFillColour Set the fill colour (opposite of the bar)
func SetFillColour(red float32, green float32, blue float32) {
	options := Options{
		FillColour: []float32{red, green, blue}}
	settings = mergeSettings(settings, options)
}
func GetFillColour() []float32 {
	return settings.FillColour
}

// SetLightTextColour Set text colour that appears over dark colours
func SetLightTextColour(red float32, green float32, blue float32) {
	options := Options{
		LightTextColour: []float32{red, green, blue}}
	settings = mergeSettings(settings, options)
}
func GetLightTextColour() []float32 {
	return settings.LightTextColour
}

// SetDarkTextColour Set text colour that appears over light colours
func SetDarkTextColour(red float32, green float32, blue float32) {
	options := Options{
		DarkTextColour: []float32{red, green, blue}}
	settings = mergeSettings(settings, options)
}
func GetDarkTextColour() []float32 {
	return settings.DarkTextColour
}

// SetInvertColours Forgo the light and dark colour and use the fill/bar colours
func SetInvertColours(invert bool) {
	options := Options{
		InvertColours: invert}
	settings = mergeSettings(settings, options)
}
func GetInvertColours() bool {
	return settings.InvertColours
}

// SetIgnoreColour Remove all colour information
func SetIgnoreColour(ignore bool) {
	options := Options{
		IgnoreColour: ignore}
	settings = mergeSettings(settings, options)
}
func GetIgnoreColour() bool {
	return settings.IgnoreColour
}

// SetPartials Set partials for the smooth bar
func SetPartials(partials []string) {
	options := Options{
		Partials: partials}
	settings = mergeSettings(settings, options)
}
func GetPartials() []string {
	return settings.Partials
}

// SetTerminators Set the start and end character for the bar
func SetTerminators(start string, end string) {
	options := Options{
		Terminators: []string{start, end}}
	settings = mergeSettings(settings, options)
}
func GetTerminators() []string {
	return settings.Terminators
}

// SetModePercent Set to display progress
func SetModePercent() {
	options := Options{
		Mode: 1}
	settings = mergeSettings(settings, options)
}
func GetModePercent() bool {
	return settings.Mode == 1
}

// SetModeTimer Set to display to timer
func SetModeTimer() {
	options := Options{
		Mode: int8(2)}
	settings = mergeSettings(settings, options)
}
func GetModeTimer() bool {
	return settings.Mode == 2
}

// SetModePortions Set to display portions
func SetModePortions() {
	options := Options{
		Mode: 3}
	settings = mergeSettings(settings, options)
}
func GetModePortions() bool {
	return settings.Mode == 3
}

// SetMode Set to display the mode
func SetMode(mode int8) {
	options := Options{
		Mode: mode}
	settings = mergeSettings(settings, options)
}
func GetMode() int8 {
	return settings.Mode
}

// ResetTimer Reset the timer if a bar does not complete
func ResetTimer() {
	timer = time.Time{}
	settings.Progress = 0
}

// DrawSimpleProgressBar Draw simple progress with text (and colour)
// Repeat a character with a dividing character and a repeating character
func DrawSimpleProgressBar(progress int, options ...Options) {
	// Load settings
	var tempSettings = Options{}
	if len(options) > 0 {
		tempSettings = mergeSettings(settings, options[0])
	} else {
		tempSettings = settings
	}

	if timer.IsZero() {
		timer = time.Now()
	}
	footer := getFooter(progress, timer, tempSettings)
	footer = fmt.Sprintf("%s %s", footer, tempSettings.Footer)

	// Keep the minimum width of the bar to 25
	width := getFinalWidth(tempSettings.Width, tempSettings.Title, footer, tempSettings.Terminators, 0)

	// Get colours
	barColor := getAnsiCode(tempSettings.BarColour[0], tempSettings.BarColour[1], tempSettings.BarColour[2])
	fillColour := getAnsiCode(tempSettings.FillColour[0], tempSettings.FillColour[1], tempSettings.FillColour[2])

	barCharColour := ansiForegroundCode(barColor, tempSettings.IgnoreColour)
	fillCharColour := ansiForegroundCode(fillColour, tempSettings.IgnoreColour)

	// Calculate the fill of the bar
	barWidth := int(float64(progress) / float64(tempSettings.Total) * float64(width))

	bar := strings.Repeat(string(tempSettings.BarCharacter), barWidth)
	var fillLength int
	if width-barWidth > 0 {
		bar = bar + string(tempSettings.BarDividerCharacter)
		fillLength = width - barWidth - 1
	} else {
		fillLength = width - barWidth
	}
	fill := strings.Repeat(string(tempSettings.FillCharacter), fillLength)

	// Prevent bar for exceeding 100%
	if progress >= tempSettings.Total {
		progress = tempSettings.Total
		timer = time.Time{}
	}

	// Update the progress bar
	fmt.Printf("\r %s %s%s%s%s%s%s%s%s\r",
		tempSettings.Title,
		tempSettings.Terminators[0],
		barCharColour,
		bar,
		fillCharColour,
		fill,
		"\u001B[0;m",
		tempSettings.Terminators[1],
		footer)
}

// DrawTrainProgressBar Draw progress that moves text (and colour) from left to right
func DrawTrainProgressBar(progress int, options ...Options) {
	// Load settings
	var tempSettings = Options{}
	if len(options) > 0 {
		tempSettings = mergeSettings(settings, options[0])
	} else {
		tempSettings = settings
	}
	if len(tempSettings.Text) == 0 {
		tempSettings.Text = "|"
	}

	if timer.IsZero() {
		timer = time.Now()
	}
	footer := getFooter(progress, timer, tempSettings)
	footer = fmt.Sprintf("%s %s", footer, tempSettings.Footer)

	// Keep the minimum width of the bar to 25
	width := getFinalWidth(tempSettings.Width, tempSettings.Title, footer, tempSettings.Terminators, 0)
	barWidth := int(float64(progress) / float64(tempSettings.Total) * float64(width))

	textArray := []string{}
	for _, char := range tempSettings.Text {
		textArray = append(textArray, string(char))
	}
	if len(textArray) < width {
		paddedArray := make([]string, width-len(textArray))
		for i := range paddedArray {
			paddedArray[i] = " "
		}
		textArray = append(paddedArray, textArray...)
	} else {
		textArray = textArray[len(textArray)-width:]
	}
	text := strings.Join(textArray[len(textArray)-barWidth:], "")

	// Prevent bar for exceeding 100%
	if progress >= tempSettings.Total {
		progress = tempSettings.Total
		timer = time.Time{}
	}

	// Get colours
	barColor := getAnsiCode(tempSettings.BarColour[0], tempSettings.BarColour[1], tempSettings.BarColour[2])
	fillColour := getAnsiCode(tempSettings.FillColour[0], tempSettings.FillColour[1], tempSettings.FillColour[2])

	barCharColour := ansiForegroundCode(barColor, tempSettings.IgnoreColour)
	fillCharColour := ansiForegroundCode(fillColour, tempSettings.IgnoreColour)

	// Update the progress bar
	fmt.Printf("\r %s %s%s%s%s%s%s%s%s\r",
		tempSettings.Title,
		tempSettings.Terminators[0],
		barCharColour,
		text,
		fillCharColour,
		strings.Repeat(string(tempSettings.FillCharacter), width-barWidth),
		"\u001B[0;m",
		tempSettings.Terminators[1],
		footer)
}

// DrawDetailedProgressBar Draw progress that contains text in the bar
func DrawDetailedProgressBar(progress int, options ...Options) {
	// Load settings
	var tempSettings = Options{}
	if len(options) > 0 {
		tempSettings = mergeSettings(settings, options[0])
	} else {
		tempSettings = settings
	}

	if timer.IsZero() {
		timer = time.Now()
	}
	footer := getFooter(progress, timer, tempSettings)
	footer = fmt.Sprintf("%s %s", footer, tempSettings.Footer)

	// Keep the minimum width of the bar to 25
	width := getFinalWidth(tempSettings.Width, tempSettings.Title, footer, tempSettings.Terminators, 0)

	// Prevent bar for exceeding 100%
	if progress >= tempSettings.Total {
		progress = tempSettings.Total
		timer = time.Time{}
	}

	// Get colours
	colourCodes := getColours(tempSettings)

	// Trimmed text
	text := fmt.Sprintf(fmt.Sprintf("%-*.*s", width, width, tempSettings.Text))

	// Calculate the fill of the bar
	barWidth := int(float64(progress) / float64(tempSettings.Total) * float64(width))

	// Update the progress bar
	print(settings.Terminators[0])
	fmt.Printf("\r %s %s%s%s%s%s%s%s%s%s%s\r",
		tempSettings.Title,
		tempSettings.Terminators[0],
		colourCodes.BarColour,
		colourCodes.BarTextColour,
		text[0:barWidth],
		colourCodes.FillColour,
		colourCodes.FillTextColour,
		text[barWidth:],
		"\u001B[0;m",
		tempSettings.Terminators[1],
		footer)
}

// DrawSmoothProgressBar Draw progress that subdivides the character into sections for making a more seamless motion
func DrawSmoothProgressBar(progress int, options ...Options) {
	// Load settings
	var tempSettings = Options{}
	if len(options) > 0 {
		tempSettings = mergeSettings(settings, options[0])
	} else {
		tempSettings = settings
	}
	partialsCount := len(tempSettings.Partials)

	if timer.IsZero() {
		timer = time.Now()
	}
	footer := getFooter(progress, timer, tempSettings)
	footer = fmt.Sprintf("%s %s", footer, tempSettings.Footer)

	// Keep the minimum width of the bar to 25
	width := getFinalWidth(tempSettings.Width, tempSettings.Title, footer, tempSettings.Terminators, 1)

	width = width * partialsCount

	// Prevent bar for exceeding 100%
	if progress >= tempSettings.Total {
		progress = tempSettings.Total
	}

	// Get colours
	barColor := getAnsiCode(tempSettings.BarColour[0], tempSettings.BarColour[1], tempSettings.BarColour[2])
	fillColour := getAnsiCode(tempSettings.FillColour[0], tempSettings.FillColour[1], tempSettings.FillColour[2])

	barForeground := ansiBackgroundCode(barColor, tempSettings.IgnoreColour)
	barBackground := ansiForegroundCode(barColor, tempSettings.IgnoreColour)
	fillBackground := ansiBackgroundCode(fillColour, tempSettings.IgnoreColour)

	// Calculate the fill of the bar
	barWidth := int(float64(progress) / float64(tempSettings.Total) * float64(width))
	quotient := barWidth / partialsCount
	remainder := barWidth % partialsCount

	// Fill the bar at 100%
	if progress >= tempSettings.Total {
		quotient = width / partialsCount
		remainder = partialsCount - 1
	}

	// Update the progress bar
	fmt.Printf("\r %s %s%s%s%s%s%s%s%s%s%s%s\r",
		tempSettings.Title,
		tempSettings.Terminators[0],
		barForeground,
		barBackground,
		strings.Repeat("▉", quotient),
		barForeground,
		fillBackground,
		tempSettings.Partials[remainder],
		strings.Repeat(" ", (width/partialsCount)-quotient),
		"\u001B[0;m",
		tempSettings.Terminators[1],
		footer)
}

// DrawWaitBar Draw a bar with no timer, or progress that just moves text left to right repeating
func DrawWaitBar(options ...Options) {
	// Load settings
	var tempSettings = Options{}
	if len(options) > 0 {
		tempSettings = mergeSettings(settings, options[0])
	} else {
		tempSettings = settings
	}
	if len(tempSettings.Text) == 0 {
		tempSettings.Text = "> "
	}
	if len(tempSettings.Text) == 1 {
		tempSettings.Text += " "
	}

	settings.Progress += 1
	if settings.Progress == math.MaxInt32 {
		settings.Progress = 0
	}
	var footer string
	switch tempSettings.Mode {
	case 1:
		footer = strings.Repeat(" ", 5)
	case 2:
		footer = strings.Repeat(" ", 6)
	case 3:
		numLength := len(strconv.Itoa(tempSettings.Total))
		footer = strings.Repeat(" ", (numLength*2)+1)
	default:
		footer = ""
	}
	footer = fmt.Sprintf(" %s%s", tempSettings.Footer, footer)

	// Keep the minimum width of the bar to 25
	width := getFinalWidth(tempSettings.Width, tempSettings.Title, footer, tempSettings.Terminators, 0)
	var text []string
	for _, char := range tempSettings.Text {
		text = append(text, string(char))
	}

	// Set the fraction of the string to start and append a repeating loop of the string
	offset := len(text) - (settings.Progress % len(text))
	barStart := strings.Join(text[offset:len(text)], "")
	trailing := strings.Repeat(strings.Join(text, ""), width+2)
	fullBar := fmt.Sprintf("%*.*s", width, width, barStart+trailing)

	// Get colours
	barColor := getAnsiCode(tempSettings.BarColour[0], tempSettings.BarColour[1], tempSettings.BarColour[2])
	backColour := getAnsiCode(tempSettings.FillColour[0], tempSettings.FillColour[1], tempSettings.FillColour[2])

	barBackground := ansiForegroundCode(barColor, tempSettings.IgnoreColour)
	backBackground := ansiBackgroundCode(backColour, tempSettings.IgnoreColour)

	// Update the progress bar
	fmt.Printf("\r %s %s%s%s%s%s%s%s\r",
		tempSettings.Title,
		tempSettings.Terminators[0],
		barBackground,
		backBackground,
		fullBar,
		"\u001B[0;m",
		tempSettings.Terminators[1],
		footer)
}

// EraseProgressBar Clear the line dood for clearing when done
func EraseProgressBar() {
	fmt.Printf("\r%s\r", strings.Repeat(" ", getWindowWidth()))
	ResetTimer()
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

// getColours Get colours from the options
func getColours(options Options) ColourCodes {
	// Colours
	foregroundCode := getAnsiCode(options.BarColour[0], options.BarColour[1], options.BarColour[2])
	backgroundCode := getAnsiCode(options.FillColour[0], options.FillColour[1], options.FillColour[2])

	darkCode := getAnsiCode(options.DarkTextColour[0], options.DarkTextColour[1], options.DarkTextColour[2])
	lightCode := getAnsiCode(options.LightTextColour[0], options.LightTextColour[1], options.LightTextColour[2])

	foreground := ansiBackgroundCode(foregroundCode, false)
	background := ansiBackgroundCode(backgroundCode, false)

	var foregroundText string
	var backgroundText string
	foregroundContrast := colourContrast(options.BarColour[0], options.BarColour[1], options.BarColour[2])
	backgroundContrast := colourContrast(options.FillColour[0], options.FillColour[1], options.FillColour[2])

	darkText := ansiForegroundCode(darkCode, false)
	lightText := ansiForegroundCode(lightCode, false)

	if settings.InvertColours {
		foregroundText = ansiForegroundCode(backgroundCode, false)
		backgroundText = ansiForegroundCode(foregroundCode, false)
	} else {
		if foregroundContrast > 0.5 {
			foregroundText = darkText
		} else {
			foregroundText = lightText
		}
		if backgroundContrast > 0.5 {
			backgroundText = darkText
		} else {
			backgroundText = lightText
		}
	}

	return ColourCodes{
		BarColour:      foreground,
		FillColour:     background,
		BarTextColour:  foregroundText,
		FillTextColour: backgroundText,
	}
}

// getAnsiCode Get the ansi code for a colour as string
func getAnsiCode(red float32, green float32, blue float32) string {
	var ansiCode = 0
	if int(red*5) == int(green*5) && int(green*5) == int(blue*5) {
		ansiCode = 232 + int(red*23)
	} else {
		ansiCode = int(16 + (36 * math.Round(float64(red)*5)) + (6 * math.Round(float64(green)*5)) + math.Round(float64(blue)*5))
	}
	return fmt.Sprintf("%d", ansiCode)
}

// ansiBackgroundCode Get the ansi code string with escape for the background
func ansiBackgroundCode(code string, ignore bool) string {
	if ignore {
		return "\033[0m"
	} else {
		return "\033[48;5;" + code + "m"
	}
}

// ansiForegroundCode Get the ansi code string with escape for the foreground
func ansiForegroundCode(code string, ignore bool) string {
	if ignore {
		return "\033[0m"
	} else {
		return "\033[38;5;" + code + "m"
	}
}

// colourContrast Get the precised contrast on a scale of 0.0 - 1.0
func colourContrast(red float32, green float32, blue float32) float32 {
	colourValue := ((red * 299) + (green * 587) + (blue * 114)) / 1000

	return colourValue
}

// getFooter Get the footer based on the choice and do the calculation
func getFooter(progress int, timer time.Time, options Options) string {
	var footer string

	switch options.Mode {
	case 1:
		// Calculate the percentage
		percentage := int(float64(progress) / float64(options.Total) * 100)
		footer = fmt.Sprintf(" %3d%%", percentage)

	case 2:
		// Calculate time
		timeRemainingSlices := float32(time.Since(timer).Seconds()) / float32(progress)
		timeRemaining := timeRemainingSlices * (float32(options.Total) - float32(progress))

		// Get minutes and seconds remaining
		minutes := int(timeRemaining) / 60
		seconds := int(timeRemaining) % 60

		footer = fmt.Sprintf("%3.1d:%02d", minutes, seconds)
	case 3:
		numLength := len(strconv.Itoa(options.Total))
		footer = fmt.Sprintf("%*.0d/%*.0d", numLength, progress, numLength, options.Total)
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

// getWinSize Get the full size of the window
func getWinSize() (*unix.Winsize, error) {
	// Get window dimensions for Unix
	winDimensions, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, err
	}
	return winDimensions, nil
}
