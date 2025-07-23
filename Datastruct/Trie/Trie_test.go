package Trie_test

import (
	"strconv"
	"testing"

	"github.com/Synertry/GoSysUtils/Datastruct/Trie"
	"github.com/Synertry/GoSysUtils/Math/Int"
	"github.com/Synertry/GoSysUtils/Slice"
	"github.com/Synertry/GoSysUtils/Str"
)

var resultBool bool

func TestTrie(t *testing.T) {
	// Table-driven test cases
	tests := map[string][]string{
		"empty":            {},
		"single_word":      {"hello"},
		"multiple_words":   {"hello", "world", "trie"},
		"mixed_case":       {"Hello", "World", "Trie"},
		"with_spaces":      {" hello ", " world ", " trie "},
		"duplicates":       {"hello", "hello", "world"},
		"long_words":       {"supercalifragilisticexpialidocious", "antidisestablishmentarianism"},
		"short_words":      {"a", "b", "c"},
		"RandomStrings":    Slice.RemoveElement(Slice.RandomStrings(100), ""),
		"RandomStringsLen": Slice.RandomStringsLen(100, 10),
	}

	/*
		// remove empty strings from RandomStrings
		for idx, word := range tests["RandomStrings"] {
			if len(word) == 0 {
				// slices.Delete uses slow but ordered "zeroing" delete: https://github.com/golang/go/blob/71c2bf551303930faa32886446910fa5bd0a701a/src/slices/slices.go#L230
				// tests["RandomStrings"] = slices.Delete(tests["RandomStrings"], idx, idx+1)

				// our swap and reslice is faster
				tests["RandomStrings"] = Slice.RemoveIndex(tests["RandomStrings"], idx)
			}
		}
	*/

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			trie := Trie.InitTrie()
			for _, w := range tc {
				trie.Insert(w)
			}
			for _, w := range tc {
				if !trie.Search(w) {
					t.Errorf("expected: '%s', but it was not found", w)
				}
			}
		})
	}
}

func FuzzTrie(f *testing.F) {
	trie := Trie.InitTrie()

	testcases := []string{
		"hello",
		"world",
		"trie",
		"test",
		"fuzz",
		"example",
	}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, w string) {
		if len(w) == 0 {
			return // Skip empty strings
		}
		trie.Insert(w)

		if !trie.Search(w) {
			t.Errorf("expected: '%s', but it was not found", w)
		}
	})
}

func BenchmarkTrieInsert(b *testing.B) {
	maxExpArrLen := 4
	type benchmark struct {
		name string
		len  int
	}
	benchmarks := make([]benchmark, maxExpArrLen+1) // do not use maps! Order will be randomized; + 1 for 2^0

	for i := 0; i <= maxExpArrLen; i++ {
		arrLen := Int.Pow(10, i)
		benchmarks[i] = benchmark{name: "ArrLen10^" + strconv.Itoa(i), len: arrLen}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			trie, words := Trie.InitTrie(), Slice.RandomStringsLen(bm.len, 10)
			for b.Loop() {
				for _, w := range words {
					trie.Insert(w)
				}
			}
		})
	}
}

func BenchmarkTrieSearch(b *testing.B) {
	maxExpStrLen := 4
	type benchmark struct {
		name string
		len  int
	}
	benchmarks := make([]benchmark, maxExpStrLen+1) // + 1 for single 10^0 -> 1

	for i := 0; i <= maxExpStrLen; i++ { // -1 as start, because substraction is more costly than addition
		strLen := Int.Pow(10, i)
		benchmarks[i] = benchmark{name: "StrLen10^" + strconv.Itoa(i), len: strLen}
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			trie, words := Trie.InitTrie(), Slice.RandomStringsLen(1000, 10)
			for _, w := range words {
				trie.Insert(w)
			}
			for b.Loop() {
				trie.Search(Str.Random(bm.len))
			}
		})
	}
}
