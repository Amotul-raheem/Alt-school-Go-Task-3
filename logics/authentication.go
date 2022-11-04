package logics

import (
	"fmt"
	"log"
)

var userPin string = "0000"

func Login() {
	userPin := getUserPin()
	for {
		if !verifyPin(userPin) {
			fmt.Println("Invalid Pin. Enter a valid or default pin.")
			getUserPin()
			continue
		}
		break
	}
}

func validatePin(pin string) (string, error) {
	if len(pin) != 4 {
		return " ", fmt.Errorf("invalid Pin. Ensure pin is 4 digits")
	}
	return pin, nil
}

func verifyPin(pin string) bool {
	pin, err := validatePin(pin)
	if err != nil {
		log.Println(err)
		return false
	}

	if pin != userPin {
		return false
	}
	return true
}

func ChangeUserPin() {
	var oldPin string
	var newPin string

	fmt.Println("Enter Old or default Pin: ")
	fmt.Scan(&oldPin)
	oldPin, err := validatePin(oldPin)

	if err != nil {
		log.Println("Ensure pin is 4 digits")
	}

	if !verifyPin(oldPin) {
		log.Println("Ensure to enter old pin or default pin")
	}
	fmt.Println("Enter your new pin: ")
	fmt.Scan(&newPin)

	userPin = newPin

	fmt.Println("Your Pin has been changed successfully!")

}
func getUserPin() string {
	var pin string
	fmt.Printf("Your default pin is: %v \n", userPin)
	fmt.Println("Please enter your pin: ")
	fmt.Scan(&pin)

	return pin
}
