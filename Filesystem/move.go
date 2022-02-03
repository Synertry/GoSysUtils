package Filesystem

import "os"

// Move with preserving timestamps
func Move(source, destination string) (err error) {
	err = CopyFile(source, destination)
	if err != nil {
		return
	}

	err = os.Remove(source)
	return
}
