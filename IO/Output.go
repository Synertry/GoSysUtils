package IO

import (
	"bytes"
	"io"
	"os"
)

// GetOutput captures the output of a function to os.Stdout
// Source: https://stackoverflow.com/a/10476304/5516320 (modified)
// Will return empty string in case of errors
func GetOutput(f func()) string {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, r); err != nil {
			return
		}
		outC <- buf.String()
	}()

	// back to normal state
	if w.Close() != nil {
		return ""
	}
	os.Stdout = old // restoring the real stdout
	return <-outC
}
