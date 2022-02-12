// Package GoSysUtils contains commonly used functions
package GoSysUtils

import (
	"github.com/Synertry/GoSysUtils/File"
)

// checkErr is just a plain panic call for errors
func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	err := File.Move(`C:\Users\synertry\test\test_times.txt`, `C:\Users\synertry\test\nested\test_times.txt`)
	checkErr(err)
}
