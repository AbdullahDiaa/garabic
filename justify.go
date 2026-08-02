package garabic

import (
	"strings"
)

// forwardConnecting holds the shaped glyph forms that connect to the
// following letter and can therefore host a kashida after them
var forwardConnecting = buildForwardConnecting()

func buildForwardConnecting() map[rune]bool {
	set := make(map[rune]bool)
	for letter, shapes := range arabicAlphabetShapes {
		//Letters that never connect forward, tatweel itself and the
		//lam-alef ligatures cannot be stretched
		if isAlwaysInitial(letter) || letter == Tatweel ||
			(letter >= 'ﻵ' && letter <= 'ﻼ') {
			continue
		}
		set[shapes.Initial] = true
		set[shapes.Medial] = true
	}
	return set
}

// Justify will shape an arabic sentence and stretch its words with
// kashida (tatweel) — or pad with spaces — until the line reaches the
// requested length in runes
func Justify(input string, length int) string {
	var shapedSentence []string
	if len([]rune(input)) >= length {
		return input
	}
	neededLen := length - len([]rune(input))
	sections := strings.Fields(input)
	lengths := make([]int, len(sections))
	wordCount := len(sections)
	adjustCount := wordCount
	if !IsArabic(input) {
		adjustCount--
	}
	if adjustCount <= 0 {
		return input
	}

	for i := 0; i < adjustCount; i++ {
		lengths[i] = neededLen / adjustCount
		if i < neededLen%adjustCount {
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

// justifyWord shapes a word and stretches it with kashida at its last
// letter connection, falling back to trailing spaces for words that
// cannot be stretched. Extra kashidas compensate for runes merged into
// ligatures so the word still grows by exactly `length` runes
func justifyWord(word string, length int) string {
	if !IsArabic(word) {
		return word
	}
	shaped := shapeWord(word)
	needed := length + len([]rune(word)) - len([]rune(shaped))
	if needed <= 0 {
		return shaped
	}
	runes := []rune(shaped)
	//Shaped words are reversed for RTL, so the last connection of the
	//word is the first forward-connecting glyph in the slice
	for i, r := range runes {
		if forwardConnecting[r] {
			stretched := make([]rune, 0, len(runes)+needed)
			stretched = append(stretched, runes[:i]...)
			for j := 0; j < needed; j++ {
				stretched = append(stretched, Tatweel)
			}
			return string(append(stretched, runes[i:]...))
		}
	}
	return shaped + strings.Repeat(" ", needed)
}
