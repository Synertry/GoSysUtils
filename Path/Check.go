package Path

import (
	"errors"
	"os"
)

// Check checks existence of provided path, returns error if error is not *PathError
func Check(path string) (pathExits bool, err error) {
	if _, err = os.Stat(path); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, err
		// Schrödinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		// SOURCE: https://stackoverflow.com/a/12518877/5516320
	}
}

// CheckDir checks if path exists and leads to a directory
func CheckDir(path string) (isDir bool, err error) {
	var exists bool
	exists, err = Check(path)
	if !exists {
		return isDir, err
	}

	var info os.FileInfo
	if info, err = os.Stat(path); err == nil {
		if info.IsDir() {
			isDir = true
		}
	}
	return isDir, err // streamlined return values
}
