package File

import (
	"os"
	"path/filepath"
)

// stats holds the metrics for file details
type stats struct {
	size             int64
	count, countDirs int
}

// statsCollector gathers data on requested path
func statsCollector(path string) (pS stats, err error) {
	err = filepath.Walk(path, func(_ string, f os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		pS.count++

		if f.IsDir() {
			pS.countDirs++
		} else {
			pS.size += f.Size()
		}
		return nil
	})
	return pS, err
}

// GetSize returns the size of the file or directory
func GetSize(path string) (int64, error) {
	sts, err := statsCollector(path)
	if err != nil {
		return 0, err
	}
	return sts.size, nil
}

// GetCount returns the number of files on the path
func GetCount(path string) (int, error) {
	sts, err := statsCollector(path)
	if err != nil {
		return 0, err
	}
	return sts.count, nil
}

// GetCountDirs returns the number of directories on the path
func GetCountDirs(path string) (int, error) {
	sts, err := statsCollector(path)
	if err != nil {
		return 0, err
	}
	return sts.countDirs, nil
}

// GetCountFiles returns the number of files on the path
func GetCountFiles(path string) (int, error) {
	sts, err := statsCollector(path)
	if err != nil {
		return 0, err
	}
	return sts.count - sts.countDirs, nil
}
