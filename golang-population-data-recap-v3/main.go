package main

import (
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} {
	population := make([]map[string]interface{}, len(data))

	for i, datum := range data {
		parts := strings.Split(datum, ";")
		person := make(map[string]interface{})
		person["name"] = parts[0]
		age, _ := strconv.Atoi(parts[1])
		person["age"] = age
		person["address"] = parts[2]

		if parts[3] != "" {
			height, _ := strconv.ParseFloat(parts[3], 64)
			person["height"] = height
		}

		if len(parts) > 4 && parts[4] != "" {
			isMarried, _ := strconv.ParseBool(parts[4])
			person["isMarried"] = isMarried
		}

		population[i] = person
	}

	return population
}
