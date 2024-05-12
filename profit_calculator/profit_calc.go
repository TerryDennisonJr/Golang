package main

import (
	"fmt"
)

func main() {
	var revenue, tax_rate, expenses float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	outputText("Tax Rate: ")
	fmt.Scan(&tax_rate)

	ebt := revenue - expenses
	fmt.Println("EBT", ebt)

	profit := (ebt) - ((ebt) * (tax_rate / 100))

	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", (ebt / profit))

}

func outputText(text string) {
	fmt.Print(text)
}
