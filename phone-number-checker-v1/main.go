package main

import (
	"fmt"
	"strings"
)

func PhoneNumberChecker(number string, result *string) {
	// TODO: answer here
	if strings.HasPrefix(number, "0811") || strings.HasPrefix(number, "0812") || strings.HasPrefix(number, "0813") || strings.HasPrefix(number, "0814") || strings.HasPrefix(number, "0815") ||
		strings.HasPrefix(number, "62811") || strings.HasPrefix(number, "62812") || strings.HasPrefix(number, "62813") || strings.HasPrefix(number, "62814") || strings.HasPrefix(number, "62815") {
		*result = "Telkomsel"
	} else if strings.HasPrefix(number, "0816") || strings.HasPrefix(number, "0817") || strings.HasPrefix(number, "0818") || strings.HasPrefix(number, "0819") ||
		strings.HasPrefix(number, "62816") || strings.HasPrefix(number, "62817") || strings.HasPrefix(number, "62818") || strings.HasPrefix(number, "62819") {
		*result = "Indosat"
	} else if strings.HasPrefix(number, "0821") || strings.HasPrefix(number, "0822") || strings.HasPrefix(number, "0823") ||
		strings.HasPrefix(number, "62821") || strings.HasPrefix(number, "62822") || strings.HasPrefix(number, "62823") {
		*result = "XL"
	} else if strings.HasPrefix(number, "0827") || strings.HasPrefix(number, "0828") || strings.HasPrefix(number, "0829") ||
		strings.HasPrefix(number, "62827") || strings.HasPrefix(number, "62828") || strings.HasPrefix(number, "62829") {
		*result = "Tri"
	} else if strings.HasPrefix(number, "0852") || strings.HasPrefix(number, "0853") || strings.HasPrefix(number, "62852") || strings.HasPrefix(number, "62853") {
		*result = "AS"
	} else if strings.HasPrefix(number, "088") || strings.HasPrefix(number, "6288") {
		*result = "Smartfren"
	} else {
		*result = "invalid"
	}
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "08813456123"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
