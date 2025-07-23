package Str_test

import (
	"testing"

	"github.com/Synertry/GoSysUtils/Str"
)

func TestRandom_Len(t *testing.T) {
	sLen := random.Intn(100)
	str := Str.Random(sLen)
	if sLen != len(str) {
		t.Errorf("expected: %d, got: %d", sLen, len(str))
	}
}

func TestRandom_Pattern(t *testing.T) {
	sLen := random.Intn(100)
	str := Str.Random(sLen)
	for _, r := range str {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			t.Errorf("expected: %c to be in range [a-Z]", r)
		}
		if t.Failed() {
			break
		}
	}
}

func BenchmarkRandom(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			var result string
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = Str.Random(bm.len)
			}
			resultString = result
		})
	}
}
