package main

import "fmt"

func GraduateStudent(score int, absent int) string {
	if score >= 70 && absent < 5 {
		return "lulus"
	} else if score >= 90 && absent == 0 {
		return "lulus"
	} else if score < 70 {
		return "tidak lulus"
	} else if absent >= 5 {
		return "tidak lulus"
	} else {
		return "lulus"
	}
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GraduateStudent(70, 4))
}
