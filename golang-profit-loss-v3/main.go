package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	if len(data) == 0 {
		return []string{}, nil
	} else {
		return data, nil
	}
	// TODO: replace this
}

func CalculateProfitLoss(data []string) string {
	var totalIncome int
	var totalExpense int

	lastTransactionDate := data[len(data)-1][:10]

	for _, line := range data {
		transactionDetails := strings.Split(line, ";")

		amount, _ := strconv.Atoi(transactionDetails[2])
		switch transactionDetails[1] {
		case "income":
			totalIncome += amount
		case "expense":
			totalExpense += amount
		default:
			continue
		}
	}

	profitLossAmount := totalIncome - totalExpense

	profitLossStatus := "profit"
	if profitLossAmount < 0 {
		profitLossAmount = -profitLossAmount
		profitLossStatus = "loss"
	}

	return fmt.Sprintf("%s;%s;%d", lastTransactionDate, profitLossStatus, profitLossAmount) // TODO: replace this
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}
	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
