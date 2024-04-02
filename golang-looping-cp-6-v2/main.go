package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	strNum := strconv.Itoa(numbers)
	max := 0

	for i := 0; i < len(strNum)-1; i++ {
		pair, _ := strconv.Atoi(strNum[i : i+2])
		if pair > max {
			max = pair
		}
	}

	return max
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}
