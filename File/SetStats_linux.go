package File

import (
	"fmt"
	"os"
)

// SetStats copies the file infos from src to dst
func SetStats(src, dst string) (err error) {
	fi, err := os.Stat(src)
	if err != nil {
		return
	}

	stat := fi.Sys().(*syscall.Stat_t)

	err = os.Chtimes(dst, time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec)), fi.ModTime())
	if err != nil {
		return fmt.Errorf("setting preserved times failed: %w", err)
	}

	err = os.Chmod(dst, fi.Mode())
	if err != nil {
		return fmt.Errorf("setting preserved modes failed: %w", err)
	}

	err = os.Chown(dst, stat.Uid, stat.Gid)
	if err != nil {
		return fmt.Errorf("setting preserved owner failed: %w", err)
	}

	return
}
