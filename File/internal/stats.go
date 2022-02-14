package internal

import (
	"fmt"
	"os"

	"github.com/djherbis/atime"
)

// SetStats copies the file infos from src to dst, does not implement chown yet
func SetStats(fi *os.FileInfo, dst *string) (err error) {
	err = os.Chtimes(*dst, atime.Get(*fi), (*fi).ModTime())
	if err != nil {
		return fmt.Errorf("setting preserved times failed: %w", err)
	}

	err = os.Chmod(*dst, (*fi).Mode())
	if err != nil {
		return fmt.Errorf("setting preserved modes failed: %w", err)
	}

	return
}
