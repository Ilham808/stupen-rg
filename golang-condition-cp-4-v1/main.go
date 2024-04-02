package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	total := float32(VIP*30 + regular*20 + student*10)

	var diskon float32
	if day%2 == 0 {
		if VIP+regular+student < 5 {
			diskon = 0.10
		} else {
			diskon = 0.20
		}
	} else {
		if VIP+regular+student < 5 {
			diskon = 0.15
		} else {
			diskon = 0.25
		}
	}

	result := total - (total * diskon)

	if result < 100 {
		return total
	}

	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
