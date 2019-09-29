package strutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"abc", "cba"},
		{"a", "a"},
		{"çınar", "ranıç"},
		{"    yağmur", "rumğay    "},
		{"επαγγελματίες", "ςείταμλεγγαπε"},
	}

	for i, test := range tests {
		output := Reverse(test.input)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleReverse() {
	fmt.Println(Reverse("επαγγελματίες"))
	// Output: ςείταμλεγγαπε
}

var reverseBenchmarkText = `It is not known exactly - επαγγελματίες`

func BenchmarkReverse(b *testing.B) {
	var s string
	for n := 0; n < b.N; n++ {
		s = Reverse(reverseBenchmarkText)
	}
	_ = s
}
