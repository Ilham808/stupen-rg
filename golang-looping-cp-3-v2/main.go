package main

import "fmt"

func CountingLetter(text string) int {
	invalidLetters := map[rune]bool{'r': true, 's': true, 't': true, 'z': true, 'R': true, 'S': true, 'T': true, 'Z': true}
	count := 0
	for _, char := range text {
		if invalidLetters[char] {
			count++
		}
	}
	return count
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
