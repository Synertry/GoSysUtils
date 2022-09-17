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
func statsCollector(path string) (stats, error) {
	var pS stats // pathStats
	err := filepath.Walk(path, func(_ string, f os.FileInfo, e error) (err error) {
		if e != nil {
			return e
		}

		pS.count++

		if f.IsDir() {
			pS.countDirs++
		} else {
			pS.size += f.Size()
		}
		return
	})
	return pS, err
}

// GetSize returns the size of the file or directory
func GetSize(path string) (int64, error) {
	stats, err := statsCollector(path)
	if err != nil {
		return 0, err
	}
	return stats.size, nil
}

// GetCount returns the number of files on the path
func GetCount(path string) (int, error) {
	stats, err := statsCollector(path)
	if err != nil {
		return 0, err
	}
	return stats.count, nil
}

// GetCountDirs returns the number of directories on the path
func GetCountDirs(path string) (int, error) {
	stats, err := statsCollector(path)
	if err != nil {
		return 0, err
	}
	return stats.countDirs, nil
}

// GetCountFiles returns the number of files on the path
func GetCountFiles(path string) (int, error) {
	stats, err := statsCollector(path)
	if err != nil {
		return 0, err
	}
	return stats.count - stats.countDirs, nil
}
