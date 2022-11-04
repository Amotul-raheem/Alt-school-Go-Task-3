package logics

import "fmt"

var accountBalance float32

func CheckBalance() {
	fmt.Printf("Your current balnce is: %v Naira", accountBalance)
}

func Withdraw() {
	fmt.Println("Withdraw cash")
	amount := getUserInput()
	if amount > accountBalance {
		fmt.Println("You do not have sufficient balance. Enter a lesser amount.")
		fmt.Printf("Your current balance is: %v \n", accountBalance)
		return
	} else if amount < 0 {
		fmt.Println("Invalid amount of money")
		fmt.Println("Enter a valid amount")
		return
	}
	accountBalance -= amount
	fmt.Printf("Your new account balance is: %v Naira", accountBalance)

}

func Deposit() {
	fmt.Println("Deposit an amount")
	amount := getUserInput()
	if amount < 0 {
		fmt.Println("Invalid amount of money")
		fmt.Println("Enter a valid amount")
	}

	fmt.Printf("You have deposited %v Naira. \n", amount)
	accountBalance += amount
	fmt.Printf("Your new account balance is: %v Naira.\n", accountBalance)
}

func getUserInput() float32 {
	var amount float32
	fmt.Println("Enter an amount:")
	fmt.Scan(&amount)

	return amount
}
