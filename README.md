[![Go Reference](https://pkg.go.dev/badge/github.com/Synertry/GoSysUtils.svg)](https://pkg.go.dev/github.com/Synertry/GoSysUtils)
[![License](https://img.shields.io/badge/License-Boost_1.0-lightblue.svg)](https://www.boost.org/LICENSE_1_0.txt)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Synertry/GoSysUtils?logo=Go)

This repo is a collection of functions and tools which I find useful in my daily tool building for Systemadministration.
I've exposed it to access it easier with go get, from anywhere.

### Notable Functions
###### File.Copy
is a file copy function which preserves permissions and filetimes.
File.Move either just renames or calls File.Copy and remove source afterwards.

###### Cmd.Timeout
is like the batch pause cmd, which waits the time in seconds or until a key is pressed.

###### IO.PrintToErr
I'm still searching for a conform error handling.
I feel like in Go most functions pass up the error to the caller.
Calling panic on an error is not my preferred way, except for really simple tools.

###### Str.Concat
Concat with underlying string builder, so it can be used in other functions.
See [this comment](https://dev.to/fasmat/comment/1k5n8) for benchmarks.

###### Path.Check
Checks if path exists.
I had enough googling multiple times for the [same answer](https://stackoverflow.com/a/12518877/5516320)
Path.CheckDir builds on top of Path.Check to check if the path is also a directory.

###### Path.ItemGetter
My personal native take on approach to recreate Get-ChildItem from PowerShell.