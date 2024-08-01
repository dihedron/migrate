package templating

import (
	"fmt"
	"os"
)

func DumpArgs(args ...interface{}) (string, error) {
	result := ""
	if args != nil {
		for i, arg := range args {
			result += fmt.Sprintf("%d => '%v' (%T)\n", i, arg, arg)
		}
		fmt.Fprintln(os.Stderr, result)
		return result, nil
	} else {
		return "<empty>", nil
	}
}
