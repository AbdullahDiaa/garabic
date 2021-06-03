//Package arabic ..
package arabic

import (
	"bytes"
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
