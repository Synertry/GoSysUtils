package IO

import (
	"fmt"
	"os"
)

// PrintToErr prints to os.Stderr; the error handling is terminating, if printing fails, it panics
func PrintToErr(str string) {
	_, err := fmt.Fprintln(os.Stderr, str)
	if err != nil {
		panic(fmt.Errorf("printing to os.Stderr failed: %w", err))
	}
}
