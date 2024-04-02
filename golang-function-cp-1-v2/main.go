package main

import (
	"fmt"
	"time"
)

func DateFormat(day, month, year int) string {
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return date.Format("02-January-2006")
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
