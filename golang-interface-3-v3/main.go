package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time interface{}) string {
	var hour, minute int

	switch val := time.(type) {
	case string:
		parts := parseStringTime(val)
		hour = parts[0]
		minute = parts[1]
	case []int:
		if len(val) != 2 {
			return "Invalid input"
		}
		hour = val[0]
		minute = val[1]
	case map[string]int:
		h, ok1 := val["hour"]
		m, ok2 := val["minute"]
		if !ok1 || !ok2 {
			return "Invalid input"
		}
		hour = h
		minute = m
	case Time:
		hour = val.Hour
		minute = val.Minute
	default:
		return "Invalid input"
	}

	if hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return "Invalid input"
	}

	ampm := "AM"
	if hour >= 12 {
		ampm = "PM"
	}
	if hour > 12 {
		hour -= 12
	}
	if hour == 0 {
		hour = 12
	}

	return fmt.Sprintf("%02d:%02d %s", hour, minute, ampm)
}

func parseStringTime(timeStr string) [2]int {
	parts := [2]int{-1, -1}
	hourMinute := strings.Split(timeStr, ":")
	if len(hourMinute) != 2 {
		return parts
	}
	hour, err1 := strconv.Atoi(hourMinute[0])
	minute, err2 := strconv.Atoi(hourMinute[1])
	if err1 != nil || err2 != nil {
		return parts
	}
	return [2]int{hour, minute}
}

func main() {
	fmt.Println(ChangeToStandartTime("16:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}
