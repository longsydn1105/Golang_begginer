package main

import (
	"fmt"

	"examle.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const balanceFileName = "balance.txt"

func main() {
	var accountBalance, err = fileops.ReadBalanceFromFile(balanceFileName)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------")
		panic("Can't continue, sorry")
	}
	var choice int

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 ",randomdata.PhoneNumber())

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
			fileops.WriteBalanceToFile(balanceFileName, accountBalance)
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
				fileops.WriteBalanceToFile(balanceFileName, accountBalance)
				fmt.Printf("Successly withdraw %v$ !! Your Balance: %v$ \n", withDrawAmount, accountBalance)
			}
		default:
			fmt.Println("Goodbye")
			return
		}
	}
}
