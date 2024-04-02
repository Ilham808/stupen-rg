package main

func CountProfit(data [][][2]int) []int {
	profits := make([]int, len(data[0]))

	for i := range data {
		for j := range data[i] {
			profits[j] += data[i][j][0] - data[i][j][1]
		}
	}

	return profits
}
