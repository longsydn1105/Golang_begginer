package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFileName = "balance.txt"

func main() {
	var accountBalance, err = readBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------")
		panic("Can't continue, sorry")
	}
	var choice int

	fmt.Println("Welcome to Go Bank!")
	for {
		showMenu()
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Printf("Your balance: %v\n", accountBalance)
		case 2:
			fmt.Print("Desposit: ")
			var despositAmount float64
			fmt.Scan(&despositAmount)

			if despositAmount <= 0 {
				fmt.Println("haha you don't have money to desposit!! LOL ")
				continue
			}

			accountBalance += despositAmount
			writeBalanceToFile(accountBalance)
			fmt.Printf("Balance update! Your Balance: %v \n ", accountBalance)
		case 3:
			fmt.Print("WithDraw Money: ")
			var withDrawAmount float64
			fmt.Scan(&withDrawAmount)

			if withDrawAmount <= 0 {
				fmt.Println("I'm increasing your balance!! 1/4 LOL ")
				continue
			}
			if withDrawAmount > accountBalance {
				fmt.Println("Not enought money in your account")
			} else {
				accountBalance -= withDrawAmount
				writeBalanceToFile(accountBalance)
				fmt.Printf("Successly withdraw %v$ !! Your Balance: %v$ \n", withDrawAmount, accountBalance)
			}
		default:
			fmt.Println("Goodbye")
			return
		}
	}
}

func showMenu() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFileName, []byte(balanceText), 0644)
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFileName)

	if err != nil {
		return 1000, errors.New("Fail to find balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Fail to convert balance")
	}
	return balance, nil
}
