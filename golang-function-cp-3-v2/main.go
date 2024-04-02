package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindShortestName(names string) string {
	namesData := strings.FieldsFunc(names, func(r rune) bool {
		return r == ' ' || r == ',' || r == ';'
	})

	sort.Slice(namesData, func(i, j int) bool {
		return namesData[i] < namesData[j]
	})

	shortestName := namesData[0]
	for _, name := range namesData[1:] {
		if len(name) < len(shortestName) {
			shortestName = name
		}
	}

	return shortestName
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
