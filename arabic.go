//Package arabic ..
package arabic

import (
	"bytes"
)

const (
	/*
		Normalizable Arabic letters
	*/

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
	//Tatweel => ـ
	Tatweel = '\u0640'

	/*
		Arabic Harakat (Harakat تَشْكِيل)
	*/
	Fatḥah     = '\u064E'
	Kasrah     = '\u0650'
	Dammah     = '\u064F'
	DaggerAlif = '\u0670'
	Sukun      = '\u0652'

	TanwinFatḥah = '\u064B'
	TanwinDammah = '\u064C'
	TanwinKasrah = '\u064D'
	Shaddah      = '\u0651'
	//AlefWaslah ٱ / Waslah is considered part of harakat/تَشْكِيل ?
	AlefWaslah = '\u0671'
)

//RemoveHarakat ...
func RemoveHarakat(input string) string {
	runes := bytes.Runes([]byte(input))
	for i := 0; i < len(runes); i++ {
		//fmt.Println(string(runes[i]))
		switch runes[i] {
		// diacritics
		case DaggerAlif, TanwinKasrah, TanwinDammah, TanwinFatḥah, Fatḥah, Dammah, Kasrah, Shaddah, Sukun:
			//Delete the matching diacritic while preserving order
			runes = deleteRune(runes, i)
			i--
		//Remove Waslah from AlefWaslah / Waslah is considered part of harakat/تَشْكِيل ?
		case AlefWaslah:
			runes[i] = Alef
		}
	}
	return string(runes)
}

//Normalize ..
func Normalize(input string) string {
	runes := bytes.Runes([]byte(input))
	for i := 0; i < len(runes); i++ {
		//fmt.Println(string(runes[i]))
		switch runes[i] {
		// diacritics
		case Tatweel, DaggerAlif, TanwinKasrah, TanwinDammah, TanwinFatḥah, Fatḥah, Dammah, Kasrah, Shaddah, Sukun:
			//Delete the matching diacritic while preserving order
			runes = deleteRune(runes, i)
			i--
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

//deleteRune will delete a rune from the slice while keeping the order of runes
func deleteRune(runes []rune, i int) []rune {
	if i >= len(runes) {
		return runes
	}
	runes = append(runes[:i], runes[i+1:]...)
	return runes
}
