package main

import "sort"

func SchedulableDays(date1 []int, date2 []int) []int {
	sort.Ints(date1)
	sort.Ints(date2)

	i, j := 0, 0
	result := []int{}

	for i < len(date1) && j < len(date2) {
		if date1[i] == date2[j] {
			result = append(result, date1[i])
			i++
			j++
		} else if date1[i] < date2[j] {
			i++
		} else {
			j++
		}
	}

	return result
}
