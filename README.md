[![Go Reference](https://pkg.go.dev/badge/github.com/Synertry/GoSysUtils.svg)](https://pkg.go.dev/github.com/Synertry/GoSysUtils)
[![License](https://img.shields.io/badge/License-Boost_1.0-lightblue.svg)](https://www.boost.org/LICENSE_1_0.txt)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Synertry/GoSysUtils?logo=Go)

This repo is a collection of functions and tools which I find useful in my daily tool building for Systemadministration with Go.
<br>
I've exposed it to access it easier with go get, from anywhere.

## Going Forward 2025-08-03

I advise you to fork this repository if you want to use it,
as I will migrate to a slightly different name [GoSynUtils](https://github.com/Synertry/GoSynUtils) as well as trying to go for 0 dependencies.
Also I made the wrong assumption that I needed to capitalize package to export them. You only need to export it's members.

## Package Overview

- Archive: Functions related to compressible artifacts, like zip, tar, gzip, bzip2
- Byte: byte manipulation
- Cmd: Functions to control and interact with the command line
- Data: Structs and functions to handle complex data structures, like heaps and tries
- File: Functions to handle file operations, like copying and moving
- IO: Functions to handle input and output streams
- JSON: Functions to handle JSON data (should be merged with other serialization packages)
- Math: Functions to handle mathematical operations, which are not covered by the standard library
- OS: Functions to handle operating system interactions like process management and system information
- Path: Functions to handle file paths, like joining, splitting, etc. (should probably be merged under pkg File)
- Self: Special package to handle the current executable, like getting its path and name (looking to be merged under another pkg)
- Slice: Functions to handle slice operations
- Str: Functions to handle string operations, like string building, reversing, and validation.
- UI: Progress bar and other UI related functions, like printing to the console. (only ProgressBar is implemented so far and should be merged under pkg IO or Cmd)

## License

This repository is licensed under the Boost Software License 1.0. See [LICENSE](https://github.com/Synertry/GoSysUtils/blob/main/LICENSE)

Further attribution belongs to:

- djherbis
  - Filetime modifier [atime](https://github.com/djherbis/atime)
- Google (The Go Authors)
  - coding language [Go](https://go.dev/)
  - low-level os interaction [golang.org/x/sys](https://cs.opensource.google/go/x/sys)
