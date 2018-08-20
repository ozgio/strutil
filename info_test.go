package strutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"    ", 0},
		{"he/she/it", 3}, //FIXME should be 3
		{"...,", 0},      //FIXME should be 3
		{`It is not known exactly when the text obtained its current standard form; it may have been as late as the 1960s. Dr. Richard McClintock, a Latin scholar who was the publications director at Hampden–Sydney College in Virginia, discovered the source of the passage sometime before 1982 while searching for instances of the Latin word "consectetur" ("that [he/she/it] pursue", subjunctive), rarely used in classical literature.[2][a] The physical source of the lorem ipsum text may be the 1914 Loeb Classical Library Edition of the De Finibus, where the Latin text, presented on the left-hand (even) pages, breaks off on page 34 with "Neque porro quisquam est qui do-" and continues on page 36 with "lorem ipsum ...", suggesting that the galley type of that page was mixed up to make the dummy text seen today.`, 138},
		{`Υπάρχουν αρκετές παραλλαγές, μερικές από τις οποίες δεν έχουν καμία σχέση με το "κλασικό" κείμενο. `, 15},
	}

	for i, test := range tests {
		output := CountWords(test.input)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func ExampleCountWords() {
	fmt.Println(CountWords("It is not known exactly!"))
	// Output: 5
}

func TestSubstring(t *testing.T) {
	tests := []struct {
		input     string
		start     int
		end       int
		expected  string
		mustPanic bool
	}{
		{"lorem", 0, 1, "l", false},
		{"", 0, 1, "", true},
		{"lorem", 0, 5, "lorem", false},
		{"lorem", 0, 10, "", true},
		{"lorem", -1, 4, "", true},
		{"lorem", 9, 10, "", true},
		{"lorem", 4, 3, "", true},
		{"Υπάρχουν", 1, 4, "πάρ", false},
		{"Υπάρχουν", 1, 0, "πάρχουν", false},
		{"Υπάρχουν", 1, 9, "", true},
		{"žůžo", 1, 4, "ůžo", false},
	}

	for i, test := range tests {
		if test.mustPanic {
			assert.Panicsf(t, func() {
				_ = Substring(test.input, test.start, test.end)
			}, "Test case %d is not successful\n", i)
		} else {
			output := Substring(test.input, test.start, test.end)
			assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
		}
	}

}

func ExampleSubstring() {
	fmt.Println(Substring("Υπάρχουν", 1, 4))
	// Output: πάρ
}

func ExampleSubstring_tillTheEnd() {
	fmt.Println(Substring("Υπάρχουν", 1, 0))
	// Output: πάρχουν
}

func TestIsASCII(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"abcdef", true},
		{"    ", true},
		{" '.'._  --  ", true},
		{"", true},
		{"Υπάρχουν", false},
	}

	for i, test := range tests {
		output := IsASCII(test.input)
		assert.Equalf(t, test.expected, output, "Test case %d is not successful\n", i)
	}
}

func TestWords(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"    ", 0},
		{" '.'._  --  ", 0},
		{"he/she/it", 3},
		{"'Hello' O'Neil, --- 1900's kebab-case snake_case camelCase ' ...,", 6},
		{`It is not known exactly when the text obtained its current standard form; it may have been as late as the 1960s. Dr. Richard McClintock, a Latin scholar who was the publications director at Hampden–Sydney College in Virginia, discovered the source of the passage sometime before 1982 while searching for instances of the Latin word "consectetur" ("that [he/she/it] pursue", subjunctive), rarely used in classical literature.[2][a] The physical source of the lorem ipsum text may be the 1914 Loeb Classical Library Edition of the De Finibus, where the Latin text, presented on the left-hand (even) pages, breaks off on page 34 with "Neque porro quisquam est qui do-" and continues on page 36 with "lorem ipsum ...", suggesting that the galley type of that page was mixed up to make the dummy text seen today.`, 138}, //FIXME should be 135
		{`Υπάρχουν αρκετές παραλλαγές, μερικές από τις οποίες δεν έχουν καμία σχέση με το "κλασικό" κείμενο. `, 15},
		{`Το Lorem Ipsum είναι απλά ένα κείμενο χωρίς νόημα για τους επαγγελματίες της τυπογραφίας και στοιχειοθεσίας. Το Lorem Ipsum είναι το επαγγελματικό πρότυπο όσον αφορά το κείμενο χωρίς νόημα, από τον 15ο αιώνα, όταν ένας ανώνυμος τυπογράφος πήρε ένα δοκίμιο και ανακάτεψε τις λέξεις για να δημιουργήσει ένα δείγμα βιβλίου. Όχι μόνο επιβίωσε πέντε αιώνες, αλλά κυριάρχησε στην ηλεκτρονική στοιχειοθεσία, παραμένοντας με κάθε τρόπο αναλλοίωτο. Έγινε δημοφιλές τη δεκαετία του '60 με την έκδοση των δειγμάτων της Letraset όπου περιελάμβαναν αποσπάσματα του Lorem Ipsum, και πιο πρόσφατα με το λογισμικό ηλεκτρονικής σελιδοποίησης όπως το Aldus PageMaker που περιείχαν εκδοχές του Lorem Ipsum`, 102},
	}

	for i, test := range tests {
		//FIXME add real output check
		output := Words(test.input)
		assert.Equalf(t, test.expected, len(output), "Test case %d is not successful\n", i)
	}
}
