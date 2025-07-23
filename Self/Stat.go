/*
 *           GoSysUtils
 *     Copyright (c) Synertry 2022 - 2025.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *           https://www.boost.org/LICENSE_1_0.txt)
 */

// Package Self provides the path to the executable and its directory
package Self

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

var (
	// PathExe is the path to the executable file
	PathExe string
	// PathExeDir is the directory containing the executable file
	// it depends on PathExe
	PathExeDir string
)

func init() {
	var err error
	PathExe, err = os.Executable()
	if err != nil {
		slog.Error(fmt.Sprintf("failed to get executable path: %v", err.Error()))
	}
	PathExeDir = filepath.Dir(PathExe)
}
