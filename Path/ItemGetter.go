package Path

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/Synertry/GoSysUtils/Str"
)

// ItemGetter is the helper struct for the GetItems method
// Fields can be set one by one, mis-modifications to Path gets fixed by fixPath
type ItemGetter struct {
	Path, Filter                string
	Include, Exclude, Extension []string
	Directory, Recurse          bool
}

// NewItemGetter inits ItemGetter with path and returns a ItemGetter struct.
func NewItemGetter(path string) *ItemGetter {
	var IG ItemGetter
	IG.Path = path
	IG.fixPath()
	return &IG
}

// fixPath applies string fixes to field Path; should not be exposed
func (IG *ItemGetter) fixPath() {
	IG.Path = strings.TrimSpace(IG.Path)
	if IG.Path == "" {
		IG.Path = "."
	}
}

// Set sets the additional params from ItemGetter
func (IG *ItemGetter) Set(filter string, include, exclude, extension []string, directory, recurse bool) {
	IG.Filter = filter
	IG.Include = include
	IG.Exclude = exclude
	IG.Extension = extension
	IG.Directory = directory
	IG.Recurse = recurse
}

// GetItemsString collects filenames and returns a slice of string.
// Calls GetItems and appends each os.DirEntry.Name() to returned []string.
func (IG *ItemGetter) GetItemsString() ([]string, error) {
	dirEntries, err := IG.GetItems()
	if err != nil {
		return nil, err
	}

	items := make([]string, 0, len(dirEntries))
	for _, v := range dirEntries {
		items = append(items, v.Name())
	}
	return items, nil
}

// GetItems is a Go native Get-ChildItem from PowerShell, although not all flags are provided.
// See https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.management/get-childitem for params info.
// path will default to "." if empty, everything else can be the initial value for the var type if not required.
//
// Filter is cross-platform case-sensitive
// Include and Exclude are case-insensitive on Windows
// Extension is case-insensitive
func (IG *ItemGetter) GetItems() (dirEntries []os.DirEntry, err error) {
	IG.fixPath()
	err = filepath.WalkDir(IG.Path, func(p string, d os.DirEntry,
		e error) (err error) {
		if e != nil {
			return e
		}

		if d.IsDir() { // directory; like "find path -maxdepth 1"
			if !IG.Recurse { // recurse
				return filepath.SkipDir
			}
			if !IG.Directory { // directory param false + dir
				return
			}
		} else {
			if len(IG.Extension) > 0 {
				hasSuffix := false
				for _, ext := range IG.Extension {
					if strings.HasSuffix(strings.ToLower(d.Name()), strings.ToLower("."+ext)) { // extension
						hasSuffix = true
						break
					}
				}
				if !hasSuffix {
					return
				}
			} else if IG.Directory {
				return
			}
		}

		if IG.Filter != "" && !strings.Contains(d.Name(), IG.Filter) { // filter
			return
		}

		var match bool
		if len(IG.Include) > 0 {
			for _, include := range IG.Include {
				if match, err = IG.patternCludes(&include, &d); match {
					break
				}
				if err != nil {
					return fmt.Errorf("invalid Include param for ItemGetter: %w", err)
				}
			}
			if !match {
				return
			}
		}

		if len(IG.Exclude) > 0 {
			for _, exclude := range IG.Exclude {
				if match, err = IG.patternCludes(&exclude, &d); match {
					return
				}
				if err != nil {
					return fmt.Errorf("invalid Exclude param for ItemGetter: %w", err)
				}
			}
		}

		dirEntries = append(dirEntries, d) // core get func
		return
	})
	return
}

// patternCludes is a helper function.
// Converts Include or Exclude path param to regexp params and determines if file name matches.
// Params are pointers to save allocs.
func (IG *ItemGetter) patternCludes(clude *string, info *os.DirEntry) (bool, error) {
	// find special path chars
	patternPathToRegexp, err := regexp.Compile(`([\^\.\(\)\[\]\\\$])`)
	if err != nil {
		return false, err
	}

	// escape path special path chars, so it does not mess with regexp chars
	patternInclude := patternPathToRegexp.ReplaceAllString(*clude, `\$1`)

	// convert glob wildcard to regexp wildcard
	patternInclude = strings.ReplaceAll(patternInclude, "*", ".*")

	// casing does not matter on Windows
	var fileCasing string
	if runtime.GOOS == "windows" {
		fileCasing = `(?i)`
	}

	return regexp.MatchString(Str.Concat(fileCasing, `^`, patternInclude, `$`), (*info).Name())
}
