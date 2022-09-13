package Archive

import (
	"archive/zip"
	"compress/flate"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// Zip compresses a file or directory into a zip file
func Zip(source, dest string) (err error) {

	var zipFile *os.File
	if zipFile, err = os.Create(dest); err != nil { // create zip file
		return fmt.Errorf("failed to create zip file: %w", err)
	}

	defer func() { // error handling closure
		errC := zipFile.Close()
		if err == nil {
			err = errC
		}
	}()

	zipWriter := zip.NewWriter(zipFile) // create zip writer interface
	defer func() {
		errC := zipWriter.Close()
		if err == nil {
			err = errC
		}
	}()

	// Register a custom Deflate compressor.
	zipWriter.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	err = filepath.Walk(source, func(p string, f fs.FileInfo, e error) (err error) {
		if e != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", p, e)
			return e
		}
		var zipPath = filepath.ToSlash(p) // converts windows path to unix path

		var fHeader *zip.FileHeader
		if fHeader, err = zip.FileInfoHeader(f); err != nil {
			return fmt.Errorf("failed to get file info header: %w", err)
		}

		fHeader.Name = zipPath

		if f.IsDir() {
			fHeader.Name += "/"
		} else {
			fHeader.Method = zip.Deflate
		}

		var zipFileWriter io.Writer
		if zipFileWriter, err = zipWriter.CreateHeader(fHeader); err != nil {
			return fmt.Errorf("failed to create file header for %s in zip: %w", fHeader.Name, err)
		}

		if f.IsDir() { // no compression for directories possible, so skip file handling
			return
		}

		// create file handle with deferred error closure
		var file *os.File
		if file, err = os.Open(p); err != nil {
			return fmt.Errorf("failed to open path %s: %w", p, err)
		}
		defer func() {
			errC := file.Close()
			if err == nil {
				err = errC
			}
		}()

		// copy file to zip
		if _, err = io.Copy(zipFileWriter, file); err != nil {
			return fmt.Errorf("failed to copy file content into zip: %w", err)
		}

		return
	})

	if err = zipWriter.SetComment("Packed by Synertry"); err != nil {
		return fmt.Errorf("failed to set comment: %w", err)
	}

	return
}
