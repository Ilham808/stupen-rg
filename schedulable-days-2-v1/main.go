package main

func intersect(a, b []int) []int {
	result := make([]int, 0)
	m := make(map[int]bool)

	for _, num := range a {
		m[num] = true
	}

	for _, num := range b {
		if m[num] {
			result = append(result, num)
		}
	}

	return result
}

func SchedulableDays(villager [][]int) []int {
	if len(villager) == 0 {
		return []int{}
	}

	commonDays := villager[0]

	for i := 1; i < len(villager); i++ {
		commonDays = intersect(commonDays, villager[i])
	}

	return commonDays
}
