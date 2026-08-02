package garabic

import (
	"fmt"
	"testing"
)

// persianNumbersTestCases
var persianNumbersTestCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		"Checking converting english digits to persian digits",
		"سال 1942",
		"سال ۱۹۴۲",
	},
	{
		"Checking converting all english digits to persian digits",
		"0123456789",
		"۰۱۲۳۴۵۶۷۸۹",
	},
	{
		"Checking text without digits stays unchanged",
		"متن فارسی",
		"متن فارسی",
	},
}

// persianShapingTestCases
var persianShapingTestCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		"Shaping word with peh and zhe",
		"پژوهش",
		"ﺶﻫوﮋﭘ",
	},
	{
		"Shaping word with gaf and cheh",
		"گچ",
		"ﭻﮔ",
	},
	{
		"Shaping word with cheh and gaf in medial position",
		"چگونه",
		"ﻪﻧﻮﮕﭼ",
	},
	{
		"Shaping word starting with zhe",
		"ژاله",
		"ﻪﻟاﮊ",
	},
	{
		"Shaping persian sentence",
		"زبان فارسی",
		"ﻰﺳرﺎﻓ ﻥﺎﺑز",
	},
}

// persianLetterTestCases
var persianLetterTestCases = []struct {
	description string
	input       rune
	expected    bool
}{
	{"Checking persian letter peh", 'پ', true},
	{"Checking persian letter cheh", 'چ', true},
	{"Checking persian letter zhe", 'ژ', true},
	{"Checking persian letter gaf", 'گ', true},
	{"Checking persian letter yeh", 'ی', true},
	{"Checking latin letter", 'p', false},
}

// TestToPersianDigits ...
func TestToPersianDigits(t *testing.T) {
	t.Log("Given an english number it should be converted to persian digits")
	{
		for i, tt := range persianNumbersTestCases {
			t.Logf("\tTest: %d\t Converting %s", i, tt.input)
			if ToPersianDigits(tt.input) != tt.expected {
				t.Errorf("\t%s\t(%s)\tShould return %s, got %s instead", failed, tt.description, tt.expected, ToPersianDigits(tt.input))
			} else {
				t.Logf("\t%s\t(%s)\tShould return %s", succeed, tt.description, tt.expected)
			}
		}
	}
}

// TestPersianDigitsRoundTrip ...
func TestPersianDigitsRoundTrip(t *testing.T) {
	t.Log("Given persian digits, ToEnglishDigits should convert them back")
	{
		input := "سال 1942 و 365 روز"
		roundTrip := ToEnglishDigits(ToPersianDigits(input))
		if roundTrip != input {
			t.Errorf("\t%s\tShould round-trip to %s, got %s instead", failed, input, roundTrip)
		} else {
			t.Logf("\t%s\tShould round-trip to %s", succeed, input)
		}
	}
}

// TestShapePersian ...
func TestShapePersian(t *testing.T) {
	t.Log("Given a persian string it should be shaped correctly")
	{
		for i, tt := range persianShapingTestCases {
			shaped := Shape(tt.input)
			t.Logf("\tTest: %d\t Shaping %s", i, tt.input)
			if shaped != tt.expected {
				t.Errorf("\t%s\t(%s)\tShould be shaped to %+q, got %+q instead", failed, tt.description, tt.expected, shaped)
			} else {
				t.Logf("\t%s\t(%s)\tShould be shaped to %s", succeed, tt.description, tt.expected)
			}
		}
	}
}

// TestIsArabicLetterPersian ...
func TestIsArabicLetterPersian(t *testing.T) {
	t.Log("Given a persian letter it should be detected as an arabic-script letter")
	{
		for i, tt := range persianLetterTestCases {
			t.Logf("\tTest: %d\t Checking %c", i, tt.input)
			if IsArabicLetter(tt.input) != tt.expected {
				t.Errorf("\t%s\t(%s)\tShould return %t", failed, tt.description, tt.expected)
			} else {
				t.Logf("\t%s\t(%s)\tShould return %t", succeed, tt.description, tt.expected)
			}
		}
	}
}

// TestIsArabicPersianText ...
func TestIsArabicPersianText(t *testing.T) {
	t.Log("Given a persian text it should be detected as arabic-script text")
	{
		if !IsArabic("متن فارسی") {
			t.Errorf("\t%s\tPersian text should be detected as arabic-script text", failed)
		} else {
			t.Logf("\t%s\tPersian text should be detected as arabic-script text", succeed)
		}
	}
}

func ExampleToPersianDigits() {
	fmt.Println(ToPersianDigits("سال 2021"))
	// Output:
	// سال ۲۰۲۱
}
