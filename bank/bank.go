package main

import "fmt"

func main() {

	var accountBalance = 1000.0
	fmt.Println("Welcom to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		//wantsCheckBalance := choice == 1
		if choice == 1 {
			fmt.Println("Your balance is:", accountBalance)

		} else if choice == 2 {
			fmt.Print("Your depost: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated. New Amount:", accountBalance)
		} else if choice == 3 {
			fmt.Print("Your withdrawl: ")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)

			if withdrawlAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}

			if withdrawlAmount > accountBalance {
				fmt.Println("Invalid Amount. Cannot withdrawl more than balance")
				continue
			}
			accountBalance -= withdrawlAmount
			fmt.Println("Balance updated. New Amount:", accountBalance)
		} else if choice == 4 {
			fmt.Println("See Ya ")
			//return
			break

		}

		//fmt.Println("Your choice:", choice)

	}
	fmt.Println("Thanks for choosing our bank!")

}
