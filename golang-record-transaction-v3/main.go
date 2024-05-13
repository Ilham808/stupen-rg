package main

import (
	"fmt"
	"os"
	"sort"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	dailyTotals := make(map[string]int)

	// Menghitung total income dan expense per tanggal
	for _, t := range transactions {
		dailyTotals[t.Date] += t.Amount
	}

	// Membuka file untuk ditulis
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Mengurutkan tanggal secara ascending
	dates := make([]string, 0, len(dailyTotals))
	for date := range dailyTotals {
		dates = append(dates, date)
	}
	sort.Strings(dates)

	// Menulis data transaksi ke file
	for _, date := range dates {
		// Tentukan jenis transaksi berdasarkan selisih total income dan total expense
		transactionType := "expense"
		if dailyTotals[date] > 0 {
			transactionType = "income"
		}

		// Tulis ke file dengan format yang sesuai
		_, err := fmt.Fprintf(file, "%s;%s;%d\n", date, transactionType, dailyTotals[date])
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 100000},
		{"01/01/2021", "expense", 50000},
		{"01/01/2021", "expense", 30000},
		{"01/01/2021", "income", 20000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
