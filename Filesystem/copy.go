package Filesystem

import (
	"io"
	"os"
	"runtime"
	"syscall"
	"time"

	"github.com/pkg/errors" // convenient stacktrace creation for errors
)

/*
	copyFile copies the contents of the file named src to the file named
	by dst. The file will be created if it does not already exist. If the
	destination file exists, all its contents will be replaced by the contents
	of the source file. The file mode will be copied from the source.
	Source: https://github.com/golang/dep/blob/v0.5.4/internal/fs/fs.go#L411
*/
func CopyFile(src, dst string) (err error) {
	if sym, err := IsSymlink(src); err != nil {
		return errors.Wrap(err, "symlink check failed")
	} else if sym {
		if err := cloneSymlink(src, dst); err != nil {
			if runtime.GOOS == "windows" {
				// If cloning the symlink fails on Windows because the user
				// does not have the required privileges, ignore the error and
				// fall back to copying the file contents.
				//
				// ERROR_PRIVILEGE_NOT_HELD is 1314 (0x522):
				// https://msdn.microsoft.com/en-us/library/windows/desktop/ms681385(v=vs.85).aspx
				if lerr, ok := err.(*os.LinkError); ok && lerr.Err != syscall.Errno(1314) {
					return err
				}
			} else {
				return err
			}
		} else {
			return nil
		}
	}

	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}

	if _, err = io.Copy(out, in); err != nil {
		out.Close()
		return
	}

	// Check for write errors on Close
	if err = out.Close(); err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}

	// Temporary fix for Go < 1.9
	//
	// See: https://github.com/golang/dep/issues/774
	// and https://github.com/golang/go/issues/20829
	if runtime.GOOS == "windows" {
		dst = fixLongPath(dst)
	}

	// preserve permissions
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}

	// preserve last modified and access time
	timeStat := si.Sys().(*syscall.Win32FileAttributeData)
	timeA := time.Unix(0, timeStat.LastAccessTime.Nanoseconds()) // AccessTime only, calls to LastModified and Creation do not return dates properly

	err = os.Chtimes(dst, timeA, timeA)

	return
}

/*
	Source: https://github.com/golang/dep/blob/v0.5.4/internal/fs/fs.go#L474
	cloneSymlink will create a new symlink that points to the resolved path of sl.
	If sl is a relative symlink, dst will also be a relative symlink.
*/
func cloneSymlink(sl, dst string) error {
	resolved, err := os.Readlink(sl)
	if err != nil {
		return err
	}

	return os.Symlink(resolved, dst)
}

/*
	Source: https://github.com/golang/dep/blob/v0.5.4/internal/fs/fs.go#L557
	IsSymlink determines if the given path is a symbolic link.
*/
func IsSymlink(path string) (bool, error) {
	l, err := os.Lstat(path)
	if err != nil {
		return false, err
	}

	return l.Mode()&os.ModeSymlink == os.ModeSymlink, nil
}

/*
	Source: https://github.com/golang/dep/blob/v0.5.4/internal/fs/fs.go#L574
	fixLongPath returns the extended-length (\\?\-prefixed) form of
	path when needed, in order to avoid the default 260 character file
	path limit imposed by Windows. If path is not easily converted to
	the extended-length form (for example, if path is a relative path
	or contains .. elements), or is short enough, fixLongPath returns
	path unmodified.

	See https://msdn.microsoft.com/en-us/library/windows/desktop/aa365247(v=vs.85).aspx#maxpath
*/
func fixLongPath(path string) string {
	// Do nothing (and don't allocate) if the path is "short".
	// Empirically (at least on the Windows Server 2013 builder),
	// the kernel is arbitrarily okay with < 248 bytes. That
	// matches what the docs above say:
	// "When using an API to create a directory, the specified
	// path cannot be so long that you cannot append an 8.3 file
	// name (that is, the directory name cannot exceed MAX_PATH
	// minus 12)." Since MAX_PATH is 260, 260 - 12 = 248.
	//
	// The MSDN docs appear to say that a normal path that is 248 bytes long
	// will work; empirically the path must be less then 248 bytes long.
	if len(path) < 248 {
		// Don't fix. (This is how Go 1.7 and earlier worked,
		// not automatically generating the \\?\ form)
		return path
	}

	// The extended form begins with \\?\, as in
	// \\?\c:\windows\foo.txt or \\?\UNC\server\share\foo.txt.
	// The extended form disables evaluation of . and .. path
	// elements and disables the interpretation of / as equivalent
	// to \. The conversion here rewrites / to \ and elides
	// . elements as well as trailing or duplicate separators. For
	// simplicity it avoids the conversion entirely for relative
	// paths or paths containing .. elements. For now,
	// \\server\share paths are not converted to
	// \\?\UNC\server\share paths because the rules for doing so
	// are less well-specified.
	if len(path) >= 2 && path[:2] == `\\` {
		// Don't canonicalize UNC paths.
		return path
	}
	if !isAbs(path) {
		// Relative path
		return path
	}

	const prefix = `\\?`

	pathbuf := make([]byte, len(prefix)+len(path)+len(`\`))
	copy(pathbuf, prefix)
	n := len(path)
	r, w := 0, len(prefix)
	for r < n {
		switch {
		case os.IsPathSeparator(path[r]):
			// empty block
			r++
		case path[r] == '.' && (r+1 == n || os.IsPathSeparator(path[r+1])):
			// /./
			r++
		case r+1 < n && path[r] == '.' && path[r+1] == '.' && (r+2 == n || os.IsPathSeparator(path[r+2])):
			// /../ is currently unhandled
			return path
		default:
			pathbuf[w] = '\\'
			w++
			for ; r < n && !os.IsPathSeparator(path[r]); r++ {
				pathbuf[w] = path[r]
				w++
			}
		}
	}
	// A drive's root directory needs a trailing \
	if w == len(`\\?\c:`) {
		pathbuf[w] = '\\'
		w++
	}
	return string(pathbuf[:w])
}

// Source: https://github.com/golang/dep/blob/v0.5.4/internal/fs/fs.go#646
func isAbs(path string) (b bool) {
	v := volumeName(path)
	if v == "" {
		return false
	}
	path = path[len(v):]
	if path == "" {
		return false
	}
	return os.IsPathSeparator(path[0])
}

// Source: https://github.com/golang/dep/blob/v0.5.4/internal/fs/fs.go#L658
func volumeName(path string) (v string) {
	if len(path) < 2 {
		return ""
	}
	// with drive letter
	c := path[0]
	if path[1] == ':' &&
		('0' <= c && c <= '9' || 'a' <= c && c <= 'z' ||
			'A' <= c && c <= 'Z') {
		return path[:2]
	}
	// is it UNC
	if l := len(path); l >= 5 && os.IsPathSeparator(path[0]) && os.IsPathSeparator(path[1]) &&
		!os.IsPathSeparator(path[2]) && path[2] != '.' {
		// first, leading `\\` and next shouldn't be `\`. its server name.
		for n := 3; n < l-1; n++ {
			// second, next '\' shouldn't be repeated.
			if os.IsPathSeparator(path[n]) {
				n++
				// third, following something characters. its share name.
				if !os.IsPathSeparator(path[n]) {
					if path[n] == '.' {
						break
					}
					for ; n < l; n++ {
						if os.IsPathSeparator(path[n]) {
							break
						}
					}
					return path[:n]
				}
				break
			}
		}
	}
	return ""
}
