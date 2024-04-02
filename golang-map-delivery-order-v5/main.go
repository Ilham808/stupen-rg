package main

import (
	"fmt"
	"strings"
)

const (
	JKT = "JKT"
	BDG = "BDG"
	BKS = "BKS"
	DPK = "DPK"
)

func DeliveryOrder(data []string, day string) map[string]float32 {
	result := make(map[string]float32)

	for _, order := range data {
		parts := strings.Split(order, ":")
		price := parseFloat(parts[2])
		location := parts[3]
		if canDeliver(location, day) {
			discount := getDiscount(day)
			totalPrice := calculateTotalPrice(price, discount)
			result[order] = totalPrice
		}
	}

	return result
}

func parseFloat(s string) float32 {
	var f float32
	fmt.Sscanf(s, "%f", &f)
	return f
}

func canDeliver(location, day string) bool {
	switch location {
	case JKT:
		return day != "minggu"
	case BDG:
		return day == "rabu" || day == "kamis" || day == "sabtu"
	case BKS:
		return day == "selasa" || day == "kamis" || day == "jumat"
	case DPK:
		return day == "senin" || day == "selasa"
	default:
		return false
	}
}

func getDiscount(day string) float32 {
	if day == "senin" || day == "rabu" || day == "jumat" {
		return 0.1
	}
	return 0.05
}

func calculateTotalPrice(price, discount float32) float32 {
	return price * (1 - discount)
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
