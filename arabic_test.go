package arabic

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

const succeed = "\u2705"
const failed = "\u274C"

//TestRemoveHarakat ...
func TestRemoveHarakat(t *testing.T) {

	t.Log("Given an arabic string it should be normalized")
	{
		for i, tt := range removeHarakatTestCases {
			normalized := RemoveHarakat(tt.input)
			t.Logf("\tTest: %d\t Normalizing %s", i, tt.input)
			if normalized != tt.expected {
				t.Errorf("\t%s\t(%s)\tShould be normalized to %s, got %s instead", failed, tt.description, tt.expected, normalized)
			} else {
				t.Logf("\t%s\t(%s)\tShould be normalized to %s", succeed, tt.description, tt.expected)
			}
		}
	}
}

//TestNormalizeBigText ..
func TestNormalizeBigText(t *testing.T) {
	originalArabicText, err := ioutil.ReadFile("test_data/bigText.txt")
	if err != nil {
		t.Errorf("\t%s\t Reading file failed with error:(%s)\t", failed, err)
	}
	preNormalizedArabicText, err := ioutil.ReadFile("test_data/normalizedBigText.txt")
	if err != nil {
		t.Errorf("\t%s\t Reading prenormalized file failed with error:(%s)\t", failed, err)
	}
	//Try to normalize test
	normalized := Normalize(string(originalArabicText))

	if normalized != string(preNormalizedArabicText) {
		t.Errorf("\t%s\t Normalized text doesn't match [length of the normalized version: %d\t length of the original prenormalized version: %d\t", failed, len(normalized), len(preNormalizedArabicText))
	}
	t.Logf("\t%s\t Should normalize all text file\t", succeed)
}

//TestNormalize ...
func TestNormalize(t *testing.T) {
	t.Log("Given an arabic string it should be normalized")
	{
		for i, tt := range normalizeTestCases {
			normalized := Normalize(tt.input)
			t.Logf("\tTest: %d\t Normalizing %s", i, tt.input)
			if normalized != tt.expected {
				t.Errorf("\t%s\t(%s)\tShould be normalized to %s, got %s instead", failed, tt.description, tt.expected, normalized)
			} else {
				t.Logf("\t%s\t(%s)\tShould be normalized to %s", succeed, tt.description, tt.expected)
			}
		}
	}
}

//TestDeleteRune ...
func TestDeleteRune(t *testing.T) {

	testCases := []struct {
		description string
		input       []rune
		index       int
		expected    []rune
	}{
		{
			description: "Deleting rune with index 0 that exists in the array",
			input:       []rune{'م', 'ح'},
			index:       0,
			expected:    []rune{'ح'},
		},
		{
			description: "Deleting rune with index 3 that exists in the array",
			input:       []rune{'م', 'ا', 'د', 'ة', 'ن'},
			index:       3,
			expected:    []rune{'م', 'ا', 'د', 'ن'},
		},
		{
			description: "Deleting rune with index 0 that doesn't exist in the array",
			input:       []rune{},
			index:       0,
			expected:    []rune{},
		},
		{
			description: "Deleting rune with index 0 that's more than array length",
			input:       []rune{'أ', 'ب'},
			index:       10,
			expected:    []rune{'أ', 'ب'},
		},
	}

	t.Log("Given a slice of runes and a position of a rune, it should be deleted from the slice while keeping order")
	{
		for i, tt := range testCases {
			input := tt.input
			input = deleteRune(input, tt.index)
			t.Logf("\tTest: %d\t Deleting rune %d from %v", i, tt.index, tt.input)
			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf(
					"\t%s\t(%s)\tShould be %v (len %d, cap %d), got %v (len %d, cap %d) instead",
					failed, tt.description, tt.expected, len(tt.expected), cap(tt.expected), input, len(input), cap(input),
				)
			} else {
				t.Logf("\t%s\t(%s)\tShould be %v", succeed, tt.description, tt.expected)
			}
		}
	}
}

func TestSpellNumber(t *testing.T) {
	t.Log("Given a number it should be return readable string of it in arabic")
	{
		for i, tt := range spellNumberTestCases {
			textOfNum := SpellNumber(tt.input)
			t.Logf("\tTest: %d\t Spelling Number %d", i, tt.input)
			if textOfNum != tt.expected {
				t.Errorf("\t%s\t\tShould be converted to %s, got %s instead", failed, tt.expected, textOfNum)
			} else {
				t.Logf("\t%s\t\tShould be converted to %s", succeed, tt.expected)
			}
		}
	}
}

func BenchmarkNormalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range normalizeTestCases {
			Normalize(c.input)
		}
	}
}

func BenchmarkRemoveHarakat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range removeHarakatTestCases {
			RemoveHarakat(c.input)
		}
	}
}

func BenchmarkNormalizeBigText(b *testing.B) {
	originalArabicText, err := ioutil.ReadFile("test_data/bigText.txt")
	if err != nil {
		b.Errorf("\t%s\t Reading file failed with error:(%s)\t", failed, err)
	}
	for i := 0; i < b.N; i++ {
		Normalize(string(originalArabicText))
	}
}

func ExampleNormalize() {
	normalized := Normalize("أحمد")
	fmt.Println(normalized)
	// Output:
	// احمد
}

func ExampleRemoveHarakat() {
	normalized := RemoveHarakat("سَنواتٌ")
	fmt.Println(normalized)
	// Output:
	// سنوات
}
