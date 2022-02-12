package File

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// Move with preserving timestamps and attributes
func Move(source, destination string) (err error) {
	if runtime.GOOS == "windows" { // TODO: examine if renaming to different filepaths under Linux work throws error
		if filepath.VolumeName(source) != filepath.VolumeName(destination) {
			return moveByCopy(source, destination)
		}
	}

	err = os.Rename(source, destination) // simple and fast move by renaming
	if err != nil {
		return fmt.Errorf("moving by renaming filepath failed: %w", err)
	}

	return
}

// moveByCopy is a helper function for Move
func moveByCopy(source string, destination string) (err error) {
	err = Copy(source, destination)
	if err != nil {
		return fmt.Errorf("moving by copying failed: %w", err)
	}

	err = os.Remove(source)
	if err != nil {
		return fmt.Errorf("removing source file of copy failed: %w", err)
	}

	return
}
