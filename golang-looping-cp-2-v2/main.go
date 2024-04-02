package main

import (
	"fmt"
)

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	var reverseResult string
	for i := len(str) - 1; i >= 0; i-- {
		reverseResult += string(str[i])
	}

	var result string
	for i := 0; i < len(reverseResult); i++ {
		if (i+1 < len(reverseResult) && reverseResult[i+1] == ' ') || reverseResult[i] == ' ' {
			result += string(reverseResult[i])
		} else {
			result += string(reverseResult[i]) + "_"
		}
	}
	return result[0 : len(result)-1]
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
