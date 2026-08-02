package garabic

import (
	"testing"
)

// TestJustifyNoPanic covers inputs that previously caused runtime panics.
func TestJustifyNoPanic(t *testing.T) {
	t.Log("Given inputs that used to panic, Justify should return normally")
	{
		cases := []struct {
			description string
			input       string
			length      int
		}{
			{"Empty input", "", 10},
			{"Single non-arabic word", "hello", 10},
			{"Harakat word with lam-alef", "وَالْآخِرَة لَا رَيْب", 30},
			{"Single harakat lam-alef word", "لَا", 10},
		}
		for i, tt := range cases {
			t.Logf("\tTest: %d\t Justifying %q to %d", i, tt.input, tt.length)
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("\t%s\t(%s)\tShould not panic, got panic: %v", failed, tt.description, r)
					}
				}()
				Justify(tt.input, tt.length)
				t.Logf("\t%s\t(%s)\tShould not panic", succeed, tt.description)
			}()
		}
	}
}

// TestJustifyPadding verifies space distribution for non-arabic sentences.
func TestJustifyPadding(t *testing.T) {
	t.Log("Given a non-arabic sentence it should be padded to the requested length")
	{
		cases := []struct {
			description string
			input       string
			length      int
			expected    string
		}{
			{"Evenly distributed", "a b c", 7, "a  b  c"},
			{"Remainder to leftmost word", "a b c", 8, "a   b  c"},
			{"No padding needed", "a b c", 5, "a b c"},
			{"Input longer than target", "a b c", 3, "a b c"},
		}
		for i, tt := range cases {
			justified := Justify(tt.input, tt.length)
			t.Logf("\tTest: %d\t Justifying %q to %d", i, tt.input, tt.length)
			if justified != tt.expected {
				t.Errorf("\t%s\t(%s)\tShould be justified to %q, got %q instead", failed, tt.description, tt.expected, justified)
			} else {
				t.Logf("\t%s\t(%s)\tShould be justified to %q", succeed, tt.description, justified)
			}
		}
	}
}

// TestJustifyKashida verifies arabic words are stretched with tatweel.
func TestJustifyKashida(t *testing.T) {
	t.Log("Given an arabic text it should be stretched with kashida to the requested length")
	{
		cases := []struct {
			description string
			input       string
			length      int
			expected    string
		}{
			{
				"Stretching a word at its last connection",
				"كتاب",
				6,
				"ﺏﺎــﺘﻛ",
			},
			{
				"Stretching compensates for the lam-alef ligature",
				"سلام",
				6,
				"ﻡﻼـــﺳ",
			},
			{
				"Unstretchable word falls back to spaces",
				"دار",
				6,
				"ﺭاد   ",
			},
			{
				"Stretching is distributed over the words of a sentence",
				"بسم الله",
				12,
				"ﻢــﺴﺑ ﻪــﻠﻟا",
			},
			{
				"Stretching a word with tashkeel and lam-alef",
				"لَاعِب",
				8,
				"ﺐِـــﻋَﻻ",
			},
		}
		for i, tt := range cases {
			justified := Justify(tt.input, tt.length)
			t.Logf("\tTest: %d\t Justifying %q to %d", i, tt.input, tt.length)
			if justified != tt.expected {
				t.Errorf("\t%s\t(%s)\tShould be justified to %+q, got %+q instead", failed, tt.description, tt.expected, justified)
			} else {
				t.Logf("\t%s\t(%s)\tShould be justified to %s", succeed, tt.description, justified)
			}
			if len([]rune(justified)) != tt.length {
				t.Errorf("\t%s\t(%s)\tShould have %d runes, got %d instead", failed, tt.description, tt.length, len([]rune(justified)))
			}
		}
	}
}
