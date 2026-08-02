package garabic

import (
	"bytes"
	"strings"
)

func Justify(input string, length int) string {
	var shapedSentence []string
	if len(input) > length {
		return input
	}
	neededLen := length - len(input)
	sections := strings.Fields(input)
	lengths := make([]int, len(sections))
	wordCount := len(sections)
	adjustCount := wordCount
	if !IsArabic(input) {
		adjustCount--
	}

	for i := 0; i < adjustCount; i++ {
		lengths[i] = neededLen / adjustCount
		if i <= neededLen%adjustCount {
			lengths[i]++
		}
	}
	for id, word := range sections {
		if IsArabic(word) {
			shapedSentence = append(shapedSentence, justifyWord(word, lengths[id]))
		} else {
			shapedSentence = append(shapedSentence, word+strings.Repeat(" ", lengths[id]))
		}
	}

	return strings.Join(shapedSentence, " ")
}

func justifyWord(word string, length int) string {
	if !IsArabic(word) {
		return word
	}
	if len(word) == 1 {
		return word + strings.Repeat(" ", length)
	}
	countMedials := 0
	for i, r := range word {
		if i > 0 && arabicAlphabetShapes[r].Final != arabicAlphabetShapes[r].Independent &&
			arabicAlphabetShapes[r].Medial != arabicAlphabetShapes[r].Independent {
			countMedials++
		}
	}

	var shapedInput bytes.Buffer

	//Convert input into runes
	inputRunes := []rune(RemoveHarakat(word))
	countIgnored := 0
	for i := range inputRunes {
		//Get Bounding back and front letters
		var backLetter, frontLetter rune
		if i-1 >= 0 {
			backLetter = inputRunes[i-1]
		}
		if i != len(inputRunes)-1 {
			frontLetter = inputRunes[i+1]
		}
		//Fix the letter based on bounding letters
		if _, ok := arabicAlphabetShapes[inputRunes[i]]; ok {
			adjustedLetter := adjustLetter(letterGroup{backLetter, inputRunes[i], frontLetter})
			if adjustedLetter != 0 {
				shapedInput.WriteRune(adjustedLetter)
			} else {
				countIgnored++
			}
		} else {
			shapedInput.WriteRune(inputRunes[i])
		}
	}

	//In case no Tashkeel deteted, same size of runes
	if len([]rune(shapedInput.String())) == len([]rune(word))-countIgnored {
		return reverse(shapedInput.String())
	}

	var shapedInputTashkeel bytes.Buffer
	inputTashkeelRunes := []rune(word)

	letterIndex := 0
	//Restore Tashkeel
	for i := range inputTashkeelRunes {
		if _, ok := arabicAlphabetShapes[inputTashkeelRunes[i]]; ok {
			shapedInputTashkeel.WriteRune([]rune(shapedInput.String())[letterIndex])
			letterIndex++
		} else {
			shapedInputTashkeel.WriteRune(inputTashkeelRunes[i])
		}
	}

	return reverse(shapedInputTashkeel.String())
}
