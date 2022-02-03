// Package GoSysUtils contains commonly used functions
package GoSysUtils

// errCheck is just a plain panic call for errors
func errCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}