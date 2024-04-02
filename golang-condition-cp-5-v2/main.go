package main

import "fmt"

func TicketPlayground(height, age int) int {
	switch {
	case age < 5:
		return -1
	case age >= 5 && age <= 7:
		if height > 120 {
			return 15000
		}
	case age >= 8 && age <= 9:
		if height > 135 {
			return 25000
		}
	case age >= 10 && age <= 11:
		if height > 150 {
			return 40000
		}
	case age == 12:
		if height > 160 {
			return 60000
		} else {
			return 60000
		}
	default:
		return 100000
	}

	if age >= 12 {
		return 100000
	}

	return -1
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
}
