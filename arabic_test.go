package arabic

import (
	"fmt"
	"reflect"
	"testing"
)

const succeed = "\u2705"
const failed = "\u274C"

//TestRemoveHarakat ...
func TestRemoveHarakat(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    string
	}{
		{
			description: "Removing Alif Khanjariyah",
			input:       "رَحْمَٰن",
			expected:    "رحمن",
		},
		{
			description: "Removing Waslah",
			input:       "ٱمْشُوا",
			expected:    "امشوا",
		},

		{
			description: "Removing all harakat 1",
			input:       "يَا أَيُّهَا الَّذِينَ آمَنُوا أَوْفُوا بِالْعُقُودِ",
			expected:    "يا أيها الذين آمنوا أوفوا بالعقود",
		},
		{
			description: "Removing all harakat 2",
			input:       "سَنواتٌ قَليلةٌ وأَدْخُلُ الجامِعَةَ، يا لها مِنْ رِحْلةٍ شاقَّةٍ طَويلَةٍ، ما أصعبَ أيامَ الدِّراسَةِ",
			expected:    "سنوات قليلة وأدخل الجامعة، يا لها من رحلة شاقة طويلة، ما أصعب أيام الدراسة",
		},
		{
			description: "Removing all harakat 3",
			input:       "إِنِّني أَشْكُرُ رَبِّي دائماً، لكنني مُنْزعجةٌ.. أَلاَ يَحِقُّ لِيَ التعبيرُ عن ضِيقِ صَدْري؟!",
			expected:    "إنني أشكر ربي دائما، لكنني منزعجة.. ألا يحق لي التعبير عن ضيق صدري؟!",
		},
	}

	t.Log("Given an arabic string it should be normalized")
	{
		for i, tt := range testCases {
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

//TestNormalize ...
func TestNormalize(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    string
	}{
		{
			description: "AlefHamzaAbove removal from beginning of word",
			input:       "أحمد",
			expected:    "احمد",
		},
		{
			description: "AlefHamzaAbove removal from middle of word",
			input:       "مأمون",
			expected:    "مامون",
		},
		{
			description: "AlefHamzaAbove removal from end of word",
			input:       "نمأ",
			expected:    "نما",
		},
		{
			description: "Replacing DotlessYae with Yae",
			input:       "منى",
			expected:    "مني",
		},
		{
			description: "Replacing TehMarbuta with Hae",
			input:       "مكتبة",
			expected:    "مكتبه",
		},
		{
			description: "Trimming Tatweel",
			input:       "بريـــــــــد",
			expected:    "بريد",
		},
		{
			description: "Removing Maddah",
			input:       "قُرْآن",
			expected:    "قران",
		},
		{
			description: "Normalizing Alif Waslah",
			input:       "ٱمْشُوا",
			expected:    "امشوا",
		},
		{
			description: "Removing all harakat combined with text normalization 1",
			input:       "يَا أَيُّهَا الَّذِينَ آمَنُوا أَوْفُوا بِالْعُقُودِ",
			expected:    "يا ايها الذين امنوا اوفوا بالعقود",
		},
		{
			description: "Removing all harakat combined with text normalization 2",
			input:       "سَنواتٌ قَليلةٌ وأَدْخُلُ الجامِعَةَ، يا لها مِنْ رِحْلةٍ شاقَّةٍ طَويلَةٍ، ما أصعبَ أيامَ الدِّراسَةِ",
			expected:    "سنوات قليله وادخل الجامعه، يا لها من رحله شاقه طويله، ما اصعب ايام الدراسه",
		},
		{
			description: "Removing all harakat combined with text normalization 3",
			input:       "إِنِّني أَشْكُرُ رَبِّي دائماً، لكنني مُنْزعجةٌ.. أَلاَ يَحِقُّ لِيَ التعبيرُ عن ضِيقِ صَدْري؟!",
			expected:    "انني اشكر ربي دائما، لكنني منزعجه.. الا يحق لي التعبير عن ضيق صدري؟!",
		},
	}

	t.Log("Given an arabic string it should be normalized")
	{
		for i, tt := range testCases {
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
			input:       []rune{'a', 'b'},
			index:       0,
			expected:    []rune{'b'},
		},
		{
			description: "Deleting rune with index 3 that exists in the array",
			input:       []rune{'a', 'b', 'c', 'd', 'e'},
			index:       3,
			expected:    []rune{'a', 'b', 'c', 'e'},
		},
		{
			description: "Deleting rune with index 0 that doesn't exist in the array",
			input:       []rune{},
			index:       0,
			expected:    []rune{},
		},
		{
			description: "Deleting rune with index 0 that's more than array length",
			input:       []rune{'a', 'v'},
			index:       10,
			expected:    []rune{'a', 'v'},
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
