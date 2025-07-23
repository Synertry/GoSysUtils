package Str_test

// holds package level variables for testing
import (
	"strconv"

	"github.com/Synertry/GoSysUtils/Math"
	"github.com/Synertry/GoSysUtils/Math/Int"
)

type benchmark struct {
	name string
	len  int
}

const (
	maxTestArrLen     = 100
	maxBenchExpArrLen = 4
)

var (
	resultString string
	random       = Math.GetRand()
	benchmarks   = make([]benchmark, maxBenchExpArrLen+1) // + 1 for empty
)

func init() {
	for i := 0; i <= maxBenchExpArrLen; i++ {
		arrLen := Int.Pow(10, i)
		benchmarks[i] = benchmark{name: "ArrLen10^" + strconv.Itoa(i), len: arrLen}
	}
}
