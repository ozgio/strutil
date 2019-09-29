package strutil

import (
	"testing"
)

func BenchmarkDrawSmallBox(b *testing.B) {
	content := "Is this the real life?"
	var r string
	for n := 0; n < b.N; n++ {
		r, _ = DrawBox(content, 40, Center)
	}
	_ = r
}

func BenchmarkDrawBigBox(b *testing.B) {
	content := `Is this the real life? Is this just fantasy?
Caught in a landslide, no escape from reality
Open your eyes, look up to the skies and see
I'm just a poor boy, I need no sympathy
Because I'm easy come, easy go, little high, little low
Any way the wind blows doesn't really matter to me, to me`
	var r string
	for n := 0; n < b.N; n++ {
		r, _ = DrawBox(content, 40, Center)
	}
	_ = r
}
