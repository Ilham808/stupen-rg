package main

func ExchangeCoin(amount int) []int {
	coins := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	result := make([]int, 0)

	for _, coin := range coins {
		for amount >= coin {
			result = append(result, coin)
			amount -= coin
		}
	}

	return result
}
