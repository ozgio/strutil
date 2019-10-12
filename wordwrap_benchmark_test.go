package strutil

import (
	"testing"
)

func BenchmarkWordWrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WordWrap(longBenchmarkText, 20, false)
	}
}

func BenchmarkWordWrap_breakLongWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WordWrap(longBenchmarkText, 20, true)
	}
}

func BenchmarkWordWrap_UTF8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WordWrap(longUTF8BenchmarkText, 20, false)
	}
}

func BenchmarkWordWrap_BreakLongWordsUTF8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WordWrap(longUTF8BenchmarkText, 20, true)
	}
}
