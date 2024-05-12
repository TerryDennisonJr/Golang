package main

import (
	"fmt"
)

// Ask for revenue, expenses, and tax rate
//Cacl earnings before tax (EBT) and earnings after tax (profit)
//Calc ratio (EBT/profit)
//Output EBT, profit, and the ratio

func main() {
	var expenses, tax_rate, revenue float64

	//Replace with custom function
	//Call custom function
	//fmt.Print("Revenue: ")
	//fmt.Scan(&revenue)
	revenue = rev()

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	outputText("Tax Rate: ")
	fmt.Scan(&tax_rate)

	//make a single function to calculate all 3 and return values
	calcs(expenses, tax_rate, revenue)
}

func outputText(text string) {
	fmt.Print(text)
}

func rev() (revenue float64) {
	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	return
}

func calcs(expenses, tax_rate, revenue float64) {
	ebt := revenue - expenses
	fmt.Println("EBT", ebt)
	profit := (ebt) - ((ebt) * (tax_rate / 100))
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", (ebt / profit))
	return
}
