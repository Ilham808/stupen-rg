package main

import (
	"fmt"
	"strings"
)

func SlurredTalk(words *string) {
	// TODO: answer here
	*words = strings.ReplaceAll(*words, "S", "L")
	*words = strings.ReplaceAll(*words, "s", "l")
	*words = strings.ReplaceAll(*words, "R", "L")
	*words = strings.ReplaceAll(*words, "r", "l")
	*words = strings.ReplaceAll(*words, "Z", "L")
	*words = strings.ReplaceAll(*words, "z", "l")
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Saya Steven, saya suka menggoreng telur dan suka hewan zebra"
	SlurredTalk(&words)
	fmt.Println(words)
}
