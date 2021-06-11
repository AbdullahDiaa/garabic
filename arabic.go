//Package arabic ..
package arabic

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
)

// Normalizable Arabic letters
var normalizable = &unicode.RangeTable{
	R16: []unicode.Range16{
		/*
			Arabic Harakat (Harakat تَشْكِيل)
		*/
		//Tatweel => ـ
		{0x0640, 0x0640, 1},
		//TanwinFatḥah
		{0x064B, 0x64B, 1},
		//TanwinDammah
		{0x064C, 0x64C, 1},
		//TanwinKasrah
		{0x064D, 0x64D, 1},
		//Fatḥah
		{0x064E, 0x64E, 1},
		//Dammah
		{0x064F, 0x64F, 1},
		//Kasrah
		{0x0650, 0x650, 1},
		//Shaddah
		{0x0651, 0x651, 1},
		//Sukun
		{0x0652, 0x652, 1},
		//DaggerAlif =>
		{0x0670, 0x0670, 1},
	},
}

// Normalizable letters [alef/Yae/Hae]
const (
	//Alef  => ا
	Alef = '\u0627'
	//AlefMad =>  آ
	AlefMad = '\u0622'
	//AlefHamzaAbove => أ
	AlefHamzaAbove = '\u0623'
	//AlefHamzaBelow إ
	AlefHamzaBelow = '\u0625'
	//Yae =>  ي
	Yae = '\u064A'
	//DotlessYae =>  ى
	DotlessYae = '\u0649'
	//TehMarbuta => ة
	TehMarbuta = '\u0629'
	//Hae => ه
	Hae = '\u0647'
	//AlefWaslah ٱ / Waslah is considered part of harakat/تَشْكِيل ?
	AlefWaslah = '\u0671'
)

//Number groups in Arabic
var _zeroToNine = []string{
	"صفر", "واحد", "اثنان", "ثلاثة", "أربعة",
	"خمسة", "ستة", "سبعة", "ثمانية", "تسعة",
}

var _elevenToNineteen = []string{
	"عشرة", "أحد عشر", "اثنا عشر", "ثلاثة عشر", "أربعة عشر",
	"خمسة عشر", "ستة عشر", "سبعة عشر", "ثمانية عشر", "تسعة عشر",
}

var _tens = []string{
	"", "", "عشرون", "ثلاثون", "أربعون", "خمسون",
	"ستون", "سبعون", "ثمانون", "تسعون",
}
var _hundreds = []string{
	"", "مئة", "مئتان", "ثلاثمئة", "أربعمئة", "خمسمئة", "ستمئة", "سبعمئة", "ثمانمئة", "تسعمئة",
}
var _scaleNumbers = []string{
	"", "ألف", "مليون", "مليار",
}

//RemoveHarakat ...
func RemoveHarakat(input string) string {
	input = normalizeTransform(input)
	runes := bytes.Runes([]byte(input))
	for i := 0; i < len(runes); i++ {
		//fmt.Println(string(runes[i]))
		switch runes[i] {
		//Remove Waslah from AlefWaslah / Waslah is considered part of harakat/تَشْكِيل ?
		case AlefWaslah:
			runes[i] = Alef
		}
	}
	return string(runes)
}

//Normalize ..
func Normalize(input string) string {
	input = normalizeTransform(input)
	runes := bytes.Runes([]byte(input))
	for i := 0; i < len(runes); i++ {
		//fmt.Println(string(runes[i]))
		switch runes[i] {
		//Normalizable letters
		case AlefMad, AlefHamzaAbove, AlefHamzaBelow, AlefWaslah:
			runes[i] = Alef
		case DotlessYae:
			runes[i] = Yae
		case TehMarbuta:
			runes[i] = Hae
		}
	}
	//@TODO: optimize runes by converting it to bytes, arabic letters use only 2 bytes
	return string(runes)
}

// Use text/transform algorithm for faster normalization
func normalizeTransform(input string) string {
	//Use text/transform algorithm for faster normalization
	tm := transform.Chain(runes.Remove(runes.In(normalizable)))
	input, _, _ = transform.String(tm, input)
	return input
}

//deleteRune will delete a rune from the slice while keeping the order of runes
func deleteRune(runes []rune, i int) []rune {
	if i >= len(runes) {
		return runes
	}
	runes = append(runes[:i], runes[i+1:]...)
	return runes
}

// SpellNumber will transform a number into a readable arabic version
func SpellNumber(input int) string {

	var stringOfNum []string

	if input < 0 {
		stringOfNum = append(stringOfNum, "سالب")
		input *= -1
	}

	if input < 10 {
		stringOfNum = append(stringOfNum, _zeroToNine[input])
		return strings.TrimSpace(strings.Join(stringOfNum, " "))
	}

	groups := []int{}

	for input > 0 {
		groups = append(groups, input%1000)
		input = input / 1000
	}

	for i := len(groups) - 1; i >= 0; i-- {
		//Get each group with its decimal position
		group := groups[i]
		if group == 0 {
			continue
		}

		// [0 0 x]
		hundreds := group / 100 % 10
		// [0 x 0]
		tens := group / 10 % 10
		// [x 0 0]
		zeros := group % 10

		if hundreds > 0 {
			if i == len(groups)-1 {
				stringOfNum = append(stringOfNum, _hundreds[hundreds])
			} else {
				stringOfNum = append(stringOfNum, "و", _hundreds[hundreds])
			}
		}

		//Move to scale number
		if tens == 0 && zeros == 0 {
			goto scale
		}

		switch tens {
		case 0:
			if zeros > 1 {
				stringOfNum = append(stringOfNum, _zeroToNine[zeros])
			}
		case 1:
			stringOfNum = append(stringOfNum, _elevenToNineteen[zeros])
			break
		default:
			if zeros > 0 {
				word := fmt.Sprintf("و %s و %s", _zeroToNine[zeros], _tens[tens])
				stringOfNum = append(stringOfNum, word)
			} else {
				stringOfNum = append(stringOfNum, "و", _tens[tens])
			}
			break
		}

		// Scale position
	scale:
		if mega := _scaleNumbers[i]; mega != "" {
			stringOfNum = append(stringOfNum, mega)
		}
	}

	return strings.TrimSpace(strings.Join(stringOfNum, " "))
}
