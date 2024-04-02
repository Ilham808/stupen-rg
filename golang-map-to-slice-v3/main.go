package main

func MapToSlice(mapData map[string]string) [][]string {
	result := make([][]string, 0)

	for key, value := range mapData {
		result = append(result, []string{key, value})
	}

	return result
}
