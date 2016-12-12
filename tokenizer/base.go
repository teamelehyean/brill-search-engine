package tokenizer

import (
	"bytes"
	"strings"
)

var stopWordsSet = make(map[string]bool)

// Tokenize splits a string into tokens
func Tokenize(text string) []string {
	tokens := strings.Split(text, " ")

	for i := 0; i < len(tokens); i++ {
		tokens[i] = strings.Trim(tokens[i], " ")
	}
	return tokens
}

// RemoveUnwantedSymbols removes symbols that symbols that are not of importance when searching
func RemoveUnwantedSymbols(text string, symbols []string) string {
	var buffer bytes.Buffer
	set := NewHashSet(symbols)
	for _, char := range text {
		if !set.Contains(string(char)) {
			buffer.WriteString(string(char))
		}
	}
	return buffer.String()
}

// RemoveNonPrintableCharacters removes control characters from text
func RemoveNonPrintableCharacters(text *string) {
	chars := []string{"\r\n", "\n", "\t"}
	for _, char := range chars {
		*text = strings.Replace(*text, char, "", -1)
	}
}

// ToLower converts a string to upper case
func ToLower(text string) string {
	return strings.ToLower(text)
}

// Set set data structure similar to sets in maths
type Set struct {
	elements map[string]bool
}

// NewHashSet creates a new hash set
func NewHashSet(elements []string) *Set {
	symbolMap := make(map[string]bool)
	for _, symbol := range elements {
		symbolMap[symbol] = true
	}
	return &Set{elements: symbolMap}
}

// Contains checks if an element exists in the set
func (s Set) Contains(element string) bool {
	_, ok := s.elements[element]
	return ok
}
