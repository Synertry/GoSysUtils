//go:build windows

package Path

import (
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows/registry"
)

// TODO: Add struct to save registry readings
// var ErrNotMapped = errors.New("drive of path is not mapped")

// GetUNC gets the UNC path for a path with a mapped drive.
// If Path is already written in UNC, returns path unchanged immediately.
// All errors will/should be terminated.
// In case of errors in the lower functions just `\\?\` will be prepended to path.
func GetUNC(path string) string {
	if strings.HasPrefix(path, `\\`) {
		return path
	}

	vol := filepath.VolumeName(path)
	driveLtr := strings.TrimSuffix(vol, ":")

	pathUNC, ok := mappedDrives()[driveLtr]
	if ok {
		return filepath.Join(pathUNC, strings.TrimPrefix(path, vol))
	} else {
		return `\\?\` + path
	}
}

// func checkDrive(driveLtr string) (path string, ok bool) {
// 	path, ok = mappedDrives()[driveLtr]
// 	return
// }

// mappedDrives builds the map of drives and remote paths from registry
func mappedDrives() (mDrives map[string]string) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Network`, registry.READ)
	if err != nil {
		return nil
	}
	defer func() {
		err := k.Close()
		if err != nil {
			mDrives = nil
		}
	}()

	ki, err := k.Stat()
	if err != nil {
		return nil
	}

	mDrives = make(map[string]string)
	if ki.SubKeyCount == 0 { // check amount of mapped drives
		return nil
	} else {
		// fmt.Println("Total number of mapped drives or printers: ", ki.SubKeyCount)
		kSubs, err := k.ReadSubKeyNames(-1)
		if err != nil {
			return nil
		}

		for _, kSub := range kSubs {
			mDrives[strings.ToUpper(kSub)] = "" // Uppercase drive letters
		}
	}

	for drive := range mDrives { // appending each path
		kD, err := registry.OpenKey(registry.CURRENT_USER, `Network\`+drive, registry.READ)
		if err != nil {
			return nil
		}

		value, _, err := kD.GetStringValue("RemotePath")
		if err != nil {
			return nil
		}
		mDrives[drive] = value

		err = kD.Close()
		if err != nil {
			return nil
		}
	}

	return mDrives
}
