package main

func ReverseData(arr [5]int) [5]int {
	reversed := [5]int{}

	for i := 0; i < 5; i++ {
		reversed[i] = reverse(arr[4-i])
	}

	return reversed
}

func reverse(num int) int {
	reversed := 0
	for num > 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}
	return reversed
}
