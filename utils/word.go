package utils

import "strings"

type Word string

func (word Word) Censor(discriminator func(letter string) bool) string {
	temporaryWord := []string{}

	for _, r := range word {
		char := string(r)

		if discriminator(char) {
			temporaryWord = append(temporaryWord, char)
		} else {
			temporaryWord = append(temporaryWord, "_")
		}
	}

	return strings.Join(temporaryWord, " ")
}

func (word Word) ContainsLetter(letter string) bool {
	reader := strings.NewReader(letter)
	r, _, _ := reader.ReadRune()

	for _, char := range word {
		if r == char {
			return true
		}
	}
	return false
}
