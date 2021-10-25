package strutil

import (
	"strings"
	"testing"
)

var longBenchmarkText = `It is not known exactly when the text obtained its current standard form; it may have been as late as the 1960s. Dr. Richard McClintock, a Latin scholar who was the publications director at Hampden–Sydney College in Virginia, discovered the source of the passage sometime before 1982 while searching for instances of the Latin word "consectetur" ("that [he/she/it] pursue", subjunctive), rarely used in classical literature.[2][a] The physical source of the lorem ipsum text may be the 1914 Loeb Classical Library Edition of the De Finibus, where the Latin text, presented on the left-hand (even) pages, breaks off on page 34 with "Neque porro quisquam est qui do-" and continues on page 36 with "lorem ipsum ...", suggesting that the galley type of that page was mixed up to make the dummy text seen today.Hey`

var longUTF8BenchmarkText = `Το Lorem Ipsum είναι απλά ένα κείμενο χωρίς νόημα για τους επαγγελματίες της τυπογραφίας και στοιχειοθεσίας. Το Lorem Ipsum είναι το επαγγελματικό πρότυπο όσον αφορά το κείμενο χωρίς νόημα, από τον 15ο αιώνα, όταν ένας ανώνυμος τυπογράφος πήρε ένα δοκίμιο και ανακάτεψε τις λέξεις για να δημιουργήσει ένα δείγμα βιβλίου. Όχι μόνο επιβίωσε πέντε αιώνες, αλλά κυριάρχησε στην ηλεκτρονική στοιχειοθεσία, παραμένοντας με κάθε τρόπο αναλλοίωτο. Έγινε δημοφιλές τη δεκαετία του '60 με την έκδοση των δειγμάτων της Letraset όπου περιελάμβαναν αποσπάσματα του Lorem Ipsum, και πιο πρόσφατα με το λογισμικό ηλεκτρονικής σελιδοποίησης όπως το Aldus PageMaker που περιείχαν εκδοχές του Lorem Ipsum`

//BenchmarkFields benchmarks builtin strings.Fields function to compare with strutil.Words
func BenchmarkFields(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strings.Fields(longBenchmarkText)
	}
}

func BenchmarkWords(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Words(longBenchmarkText)
	}
}

func BenchmarkFieldsUTF8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		strings.Fields(longUTF8BenchmarkText)
	}
}

func BenchmarkWordsUTF8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Words(longUTF8BenchmarkText)
	}
}

func BenchmarkSubstringShort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Substring("lorem ipsum", 1, 3) //nolint: errcheck
	}
}

func BenchmarkSubstringLong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Substring(longUTF8BenchmarkText, 16, 24) //nolint: errcheck
	}
}
