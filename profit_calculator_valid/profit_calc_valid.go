package main

import (
	"fmt"
	"os"
)

// 1) Validate user input
// => show error message & exit if invalid input is provided
//   - no negative numbers
//   - Not 0
// 2) store calcualted results into a file

const balance = "balance.txt"

func main() {
	var revenue, tax_rate, expenses float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	if revenue <= 0 {
		fmt.Println("Invalid value")
		return
	}
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	if expenses <= 0 {
		fmt.Println("Invalid value")
		return
	}
	outputText("Tax Rate: ")
	fmt.Scan(&tax_rate)

	if tax_rate <= 0 {
		fmt.Println("Invalid value")
		return
	}

	ebt := revenue - expenses
	fmt.Println("EBT", ebt)

	profit := (ebt) - ((ebt) * (tax_rate / 100))
	ratio := (ebt / profit)

	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", ratio)

	balText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n", ebt, profit, ratio)

	os.WriteFile(balance, []byte(balText), 0644)
}

func outputText(text string) {
	fmt.Print(text)
}
