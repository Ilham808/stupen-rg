package main

import (
	"fmt"
	"strings"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {
	result := make(map[string][]string)

	for _, item := range data {
		parts := strings.Split(item, "-")
		header := parts[0]
		value := strings.Join(parts[3:], " ")
		index := parts[1]
		resultKey := fmt.Sprintf("%s-%s", header, index)
		result[resultKey] = append(result[resultKey], value)
	}

	return result
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
