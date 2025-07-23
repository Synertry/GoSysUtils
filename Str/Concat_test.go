package Str_test

import (
	"bytes"
	"math"
	"strconv"
	"testing"

	"github.com/Synertry/GoSysUtils/Str"
	"github.com/google/go-cmp/cmp"
)

func TestConcat(t *testing.T) {
	type test struct {
		input []string
		want  string
	}
	tests := map[string]test{
		"Empty":    {input: []string{}, want: ""},
		"Single":   {input: []string{"a"}, want: "a"},
		"Multiple": {input: []string{"a", "b", "c"}, want: "abc"},
	}

	for i := 0; i < maxTestArrLen; i++ {
		input := make([]string, i)
		var want bytes.Buffer
		for j := 0; j < i; j++ {
			input[j] = Str.Random(int(math.Log(float64(i))))
			want.WriteString(input[j])
		}
		tests["Random"+strconv.Itoa(i)] = test{input: input, want: want.String()}
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Str.Concat(tc.input...)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("expected: %s, got: %s", tc.want, got)
				t.Log(diff)
				t.Logf("input: %#v\n", tc.input)
			}
		})
		if t.Failed() {
			break
		}
	}
}

func BenchmarkConcat(b *testing.B) {
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			input, result := make([]string, bm.len), ""
			for i := 0; i < bm.len; i++ {
				input[i] = Str.Random(10)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				result = Str.Concat(input...)
			}
			resultString = result
		})
	}
}
