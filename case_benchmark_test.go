package strutil

import "testing"

var slugifyTestString = "Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of \"de Finibus Bonorum et Malorum\" (The Extremes of Good and Evil) by Cicero"

func BenchmarkSlugifySpecial(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SlugifySpecial(slugifyTestString, "-")
	}
}
