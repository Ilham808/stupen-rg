package main

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	totalPrice := 0

	for _, product := range products {
		totalPrice += product.Price + product.Tax
	}

	change := amount - totalPrice
	coins := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	result := make([]int, 0)

	for _, coin := range coins {
		for change >= coin {
			result = append(result, coin)
			change -= coin
		}
	}

	return result
}
