package goprogress

const ModeNone = 0
const ModePercent = 1
const ModeTimer = 2
const ModeProportion = 3

const StyleSimple = 1
const StyleTrain = 2
const StyleDetailed = 3
const StyleSmooth = 4
const StyleWait = 0

func Black() []float32 {
	return []float32{0.0, 0.0, 0.0}
}

func LtGrey() []float32 {
	return []float32{0.8, 0.8, 0.8}
}
func Grey() []float32 {
	return []float32{0.5, 0.5, 0.5}
}
func DkGrey() []float32 {
	return []float32{0.2, 0.2, 0.2}
}

func LtRed() []float32 {
	return []float32{1.0, 0.5, 0.5}
}
func Red() []float32 {
	return []float32{1.0, 0.0, 0.0}
}
func DkRed() []float32 {
	return []float32{0.2, 0.0, 0.0}
}

func LtGreen() []float32 {
	return []float32{0.5, 1.0, 0.5}
}
func Green() []float32 {
	return []float32{0.0, 1.0, 0.0}
}
func DkGreen() []float32 {
	return []float32{0.0, 0.2, 0.0}
}

func LtBlue() []float32 {
	return []float32{0.5, 0.5, 1.0}
}
func Blue() []float32 {
	return []float32{0.1, 0.1, 1.0}
}
func DkBlue() []float32 {
	return []float32{0.0, 0.0, 0.5}
}

func LtCyan() []float32 {
	return []float32{0.5, 1.0, 1.0}
}
func Cyan() []float32 {
	return []float32{0.0, 0.8, 0.8}
}
func DkCyan() []float32 {
	return []float32{0.0, 0.2, 0.2}
}

func LtMagenta() []float32 {
	return []float32{1.0, 0.5, 1.0}
}
func Magenta() []float32 {
	return []float32{1.0, 0.0, 1.0}
}
func DkMagenta() []float32 {
	return []float32{0.2, 0.0, 0.2}
}

func LtYellow() []float32 {
	return []float32{1.0, 1.0, 0.5}
}
func Yellow() []float32 {
	return []float32{1.0, 1.0, 0.0}
}
func DkYellow() []float32 {
	return []float32{0.2, 0.2, 0.0}
}

func White() []float32 {
	return []float32{1.0, 1.0, 1.0}
}

func NoColour() []float32 {
	return []float32{-1.0, -1.0, -1.0}
}
