package templating

import (
	"fmt"
	"log/slog"
	"text/template"
)

func Dereference(value any) string {
	switch v := value.(type) {
	case int, int8, int16, int32, int64, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case *int:
		if v != nil {
			return fmt.Sprintf("%d", *v)
		}
	case *int8:
		if v != nil {
			return fmt.Sprintf("%d", *v)
		}
	case *int16:
		if v != nil {
			return fmt.Sprintf("%d", *v)
		}
	case *int32:
		if v != nil {
			return fmt.Sprintf("%d", *v)
		}
	case *int64:
		if v != nil {
			return fmt.Sprintf("%d", *v)
		}
	case *uint8:
		if v != nil {
			return fmt.Sprintf("%d", *v)
		}
	case *uint16:
		if v != nil {
			return fmt.Sprintf("%d", *v)
		}
	case *uint32:
		if v != nil {
			return fmt.Sprintf("%d", *v)
		}
	case *uint64:
		if v != nil {
			return fmt.Sprintf("%d", *v)
		}
	case float32, float64:
		return fmt.Sprintf("%f", v)
	case *float32:
		if v != nil {
			return fmt.Sprintf("%f", *v)
		}
	case *float64:
		if v != nil {
			return fmt.Sprintf("%f", *v)
		}
	case string:
		return v
	case *string:
		if v != nil {
			return *v
		}
	default:
		slog.Error("unsupported type", "type", fmt.Sprintf("%t", value))
	}
	return ""
}

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"include":     Include,
		"dump":        DumpArgs,
		"blue":        Blue,
		"cyan":        Cyan,
		"green":       Green,
		"magenta":     Magenta,
		"purple":      Magenta,
		"red":         Red,
		"yellow":      Yellow,
		"white":       White,
		"hiblue":      HighBlue,
		"hicyan":      HighCyan,
		"higreen":     HighGreen,
		"himagenta":   HighMagenta,
		"hipurple":    HighMagenta,
		"hired":       HighRed,
		"hiyellow":    HighYellow,
		"hiwhite":     HighWhite,
		"dereference": Dereference,
	}
}
