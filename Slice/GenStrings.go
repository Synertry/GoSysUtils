package Slice

import (
	"github.com/Synertry/GoSysUtils/Str"
)

func GenRandomStrings(length int) []string {
	slice := make([]string, length)
	for i := 0; i < length; i++ {
		slice[i] = Str.GenRandom((i + 1) / 2)
	}
	return slice
}
