package File

import (
	"fmt"
	"os"

	"github.com/Microsoft/go-winio"
)

// SetStats copies file times and attributes (and padding)
func SetStats(src, dst string) (err error) {
	fsrc, err := os.Open(src)
	if err != nil {
		return
	}

	defer func() {
		errCi := fsrc.Close()
		if err == nil {
			err = errCi
		}
	}()

	fdst, err := os.Open(dst)
	if err != nil {
		return
	}

	defer func() {
		errCo := fdst.Close()
		if err == nil {
			err = errCo
		}
	}()

	info, err := winio.GetFileBasicInfo(fsrc)
	if err != nil {
		return fmt.Errorf("could not retrieve file times and atrributes: %w", err)
	}

	if err = winio.SetFileBasicInfo(fdst, info); err != nil {
		return fmt.Errorf("could not set file times and atrributes: %w", err)
	}

	return
}
