package main

import (
	"fmt"
	"unicode"
)

func CountVowelConsonant(str string) (int, int, bool) {
	vowelCount := 0
	consonantCount := 0
	hasVowel := false
	hasConsonant := false

	for _, char := range str {
		if unicode.IsLetter(char) {
			switch unicode.ToLower(char) {
			case 'a', 'i', 'u', 'e', 'o':
				vowelCount++
				hasVowel = true
			default:
				consonantCount++
				hasConsonant = true
			}
		}
	}

	return vowelCount, consonantCount, !hasVowel || !hasConsonant
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
}
