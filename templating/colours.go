package templating

import (
	"github.com/fatih/color"
)

var (
	blue      = color.New(color.FgBlue).SprintFunc()
	cyan      = color.New(color.FgCyan).SprintFunc()
	green     = color.New(color.FgGreen).SprintFunc()
	magenta   = color.New(color.FgMagenta).SprintFunc()
	red       = color.New(color.FgRed).SprintFunc()
	yellow    = color.New(color.FgYellow).SprintFunc()
	white     = color.New(color.FgWhite).SprintFunc()
	hiblue    = color.New(color.FgHiBlue).SprintFunc()
	hicyan    = color.New(color.FgHiCyan).SprintFunc()
	higreen   = color.New(color.FgHiGreen).SprintFunc()
	himagenta = color.New(color.FgHiMagenta).SprintFunc()
	hired     = color.New(color.FgHiRed).SprintFunc()
	hiyellow  = color.New(color.FgHiYellow).SprintFunc()
	hiwhite   = color.New(color.FgHiWhite).SprintFunc()
)

func HighBlue(v interface{}) string {
	return hiblue(v)
}

func HighCyan(v interface{}) string {
	return hicyan(v)
}

func HighGreen(v interface{}) string {
	return higreen(v)
}

func HighMagenta(v interface{}) string {
	return himagenta(v)
}

func HighRed(v interface{}) string {
	return hired(v)
}

func HighYellow(v interface{}) string {
	return hiyellow(v)
}

func HighWhite(v interface{}) string {
	return hiwhite(v)
}

func Blue(v interface{}) string {
	return blue(v)
}

func Cyan(v interface{}) string {
	return cyan(v)
}

func Green(v interface{}) string {
	return green(v)
}

func Magenta(v interface{}) string {
	return magenta(v)
}

func Red(v interface{}) string {
	return red(v)
}

func Yellow(v interface{}) string {
	return yellow(v)
}

func White(v interface{}) string {
	return white(v)
}
