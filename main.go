package main

import (
	"atm/logics"
	"fmt"
	"log"
	"os"
)

func main() {
	logics.Login()
	welcome()
	menu()

}

func welcome() {
	fmt.Println("*********** Welcome to my ATM Application **************")
}

func menu() {
	fmt.Println("Select Operation")
	fmt.Println("1. Change Pin \n2. Check Account Balance \n3. Withdrawal \n4. Deposit \n5. Exit program")

	var SelectedOp int

	_, err := fmt.Scan(&SelectedOp)

	if err != nil {
		log.Println("Please the correct menu item")

	}
	switch SelectedOp {
	case 1:
		logics.ChangeUserPin()
		PerformAnotherAction()
	case 2:
		logics.CheckBalance()
		PerformAnotherAction()

	case 3:
		logics.Withdraw()
		PerformAnotherAction()
	case 4:
		logics.Deposit()
		PerformAnotherAction()
	case 5:
		exitProgram()
	default:
		fmt.Println("Invalid Input")
		PerformAnotherAction()
	}

}
func PerformAnotherAction() {
	var YesorNo string

	fmt.Println("\nWould you like to perform another action: (Y/N)")
	fmt.Scan(&YesorNo)

	if YesorNo == "y" || YesorNo == "Y" {
		menu()
	}
	exitProgram()
}

func exitProgram() {
	fmt.Println("Thank you for banking with us.")
	os.Exit(0)
}
