package goprogress

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
