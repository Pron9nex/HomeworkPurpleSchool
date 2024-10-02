package main

//HomeWork 1
import "fmt"

func main() {
	const USD = 94
	const EUR = 122
	var RUB float64
	RUB = 150 // For example
	convertRubToEur := RUB / EUR
	convertRubToUsd := RUB / USD
	fmt.Printf("%.2f Rub convert to Euro: %.2f\n", RUB, convertRubToEur)
	fmt.Printf("%.2f Rub convert to Usd: %.2f\n", RUB, convertRubToUsd)

}
