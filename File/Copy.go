package File

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/Synertry/GoSysUtils/File/internal"
)

// Copy copies a file from src to dst. If src and dst files exist, and are the same, then return success.
// Otherwise, copy the file contents from src to dst.
func Copy(src, dst string) (err error) {
	sfi, err := os.Stat(src)
	if err != nil {
		return
	}

	if !sfi.Mode().IsRegular() { // cannot copy non-regular files (e.g., directories, symlinks, devices, etc.)
		return fmt.Errorf("CopyFile: non-regular source file %s (%q)", sfi.Name(), sfi.Mode().String())
	}

	dfi, err := os.Stat(dst)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return
		}
	} else {
		if !(dfi.Mode().IsRegular()) {
			return fmt.Errorf("CopyFile: non-regular destination file %s (%q)", dfi.Name(), dfi.Mode().String())
		}

		if os.SameFile(sfi, dfi) {
			return
		}
	}

	/*if err = os.Link(src, dst); err == nil { // I see problems here with network paths
		return
	}*/

	err = copyContents(src, dst)
	if err != nil {
		return
	}

	return internal.SetStats(&sfi, &dst)
}

// copyContents is the core function of Copy to copy file contents
func copyContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer func() {
		errCi := in.Close()
		if err == nil {
			err = errCi
		}
	}()

	out, err := os.Create(dst)
	if err != nil {
		return
	}

	defer func() {
		errCo := out.Close()
		if err == nil {
			err = errCo
		}
	}()

	if _, err = io.Copy(out, in); err != nil { // core copy
		return fmt.Errorf("core function copy failed: %w", err)
	}

	return out.Sync() // flush
}
