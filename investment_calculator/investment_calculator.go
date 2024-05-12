package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Please enter inital investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Please enter expected Return Rate Amount: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Please enter amount of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value: ", futureValue)
	fmt.Println("Future Real Value: ", futureRealValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {

	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
	//return
}
